package verifier

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rogercoll/optimisticrp"
)

type VerifierNode struct {
	accountsTrie optimisticrp.Optimistic
	ethContract  optimisticrp.OptimisticSContract
	privKey      *ecdsa.PrivateKey
	onChainRoot  common.Hash
}

func New(newAccountsTrie optimisticrp.Optimistic, newEthContract optimisticrp.OptimisticSContract, privateKey *ecdsa.PrivateKey) *VerifierNode {
	return &VerifierNode{
		accountsTrie: newAccountsTrie,
		ethContract:  newEthContract,
		privKey:      privateKey,
	}
}

//Sync with on-chain smart contract
func (v *VerifierNode) Synced() (bool, error) {
	onChainStateRoot, err := v.ethContract.GetStateRoot()
	if err != nil {
		return false, err
	}
	if onChainStateRoot == v.accountsTrie.StateRoot() {
		return true, nil
	}
	stateRoot, err := v.computeAccountsTrie()
	if err != nil {
		return false, err
	}
	log.Printf("Computed state root: %v", stateRoot)
	log.Printf("OnChain state root: %v", onChainStateRoot)
	if stateRoot != onChainStateRoot {
		return false, fmt.Errorf("Aggregator was not able to compute a valid StateRoot")
	}
	return true, nil
}

//Generate proof data to be send onchain, a proof proves that key with a certain value exits on the trie
func (v *VerifierNode) generateProof(acc common.Address) {
	_, err := v.accountsTrie.NewProve(acc)
	if err != nil {
		return
	}
}

func (v *VerifierNode) VerifyOnChainData(logs chan<- interface{}) {
	defer close(logs)
	//Every 20 seconds scan the chain looking for new batches with errors
	ticker := time.NewTicker(20 * time.Second)
	quit := make(chan struct{})
	defer close(quit)
	for {
		select {
		case <-ticker.C:
			isSync, err := v.Synced()
			if err != nil {
				logs <- err
				//we shall continue as maybe there was a network error
				continue
			} else if isSync == false {
				logs <- fmt.Errorf("Not synced with onChain data")
				continue
			}
			logs <- "All onChain data verified"
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

//Reads all transactions to the smart contracts and computes the whole accounts trie from scratch
//IMPORTANT: This implementation uses the already defined OptimisticTrie object to prevent implementing the AddFunds and ProcessTx functions
func (v *VerifierNode) computeAccountsTrie() (common.Hash, error) {
	optimisticTrie, ok := v.accountsTrie.(*optimisticrp.OptimisticTrie)
	if ok != true {
		return common.Hash{}, fmt.Errorf("This verifier implementation uses the OptimisticTrie object, if you are not, please develop your own verifier functions")
	}
	onChainData := make(chan interface{})
	go v.ethContract.GetOnChainData(onChainData)
	stateRoot := common.Hash{}
	pendingDeposits := []optimisticrp.Deposit{}
	var err error
	for methodData := range onChainData {
		switch input := methodData.(type) {
		case optimisticrp.Batch:
			log.Println("New batch recevied")
			//if there is a new batch we MUST update the stateRoot with the previous deposits (rule 1.)
			for _, deposit := range pendingDeposits {
				err := optimisticTrie.AddFunds(deposit.From, deposit.Value)
				if err != nil {
					return stateRoot, err
				}
			}
			pendingDeposits = nil
			for _, txInBatch := range input.Transactions {
				stateRoot, err = optimisticTrie.ProcessTx(txInBatch)
				if err != nil {
					return stateRoot, err
				}
			}
		case optimisticrp.Deposit:
			log.Printf("New deposit recevied from %v\n", input.From)
			pendingDeposits = append(pendingDeposits, input)
		case error:
			return stateRoot, input

		default:
			return common.Hash{}, errors.New("There was an error while fetching onChain data")
		}
	}
	return stateRoot, nil
}
