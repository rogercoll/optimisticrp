package challenger

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rogercoll/optimisticrp"
	"github.com/sirupsen/logrus"
)

type ChallengerNode struct {
	accountsTrie optimisticrp.Optimistic
	ethContract  optimisticrp.OptimisticSContract
	privKey      *ecdsa.PrivateKey
	onChainRoot  common.Hash
	log          *logrus.Entry
}

func New(newAccountsTrie optimisticrp.Optimistic, newEthContract optimisticrp.OptimisticSContract, privateKey *ecdsa.PrivateKey, logger *logrus.Logger) *ChallengerNode {
	challengerLogger := logger.WithFields(logrus.Fields{
		"service": "Challenger",
	})
	return &ChallengerNode{
		accountsTrie: newAccountsTrie,
		ethContract:  newEthContract,
		privKey:      privateKey,
		log:          challengerLogger,
	}
}

//Sync with on-chain smart contract
func (v *ChallengerNode) Synced() (bool, error) {
	v.log.Info("Starting sync process with onchain data")
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
	v.log.WithFields(logrus.Fields{"StateRoot": stateRoot}).Info("Computed accounts state")
	v.log.WithFields(logrus.Fields{"StateRoot": onChainStateRoot}).Info("OnChain accounts state")
	if stateRoot != onChainStateRoot {
		return false, fmt.Errorf("Aggregator was not able to compute a valid StateRoot")
	}
	return true, nil
}

//Generate proof data to be send onchain, a proof proves that key with a certain value exits on the trie
func (v *ChallengerNode) generateProof(acc common.Address) {
	_, err := v.accountsTrie.NewProve(acc)
	if err != nil {
		return
	}
}

func (v *ChallengerNode) VerifyOnChainData(errs chan<- interface{}) {
	defer close(errs)
	//Every 20 seconds scan the chain looking for new batches with errors
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	defer close(quit)
	for {
		select {
		case <-ticker.C:
			isSync, err := v.Synced()
			if err != nil {
				errs <- err
				//we shall continue as maybe there was a network error
				continue
			} else if isSync == false {
				errs <- fmt.Errorf("Not synced with onChain data")
				continue
			}
			v.log.Info("All onChain data verified")
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

//Reads all transactions to the smart contracts and computes the whole accounts trie from scratch
//IMPORTANT: This implementation uses the already defined OptimisticTrie object to prevent implementing the AddFunds and ProcessTx functions
func (v *ChallengerNode) computeAccountsTrie() (common.Hash, error) {
	optimisticTrie, ok := v.accountsTrie.(*optimisticrp.OptimisticTrie)
	if ok != true {
		return common.Hash{}, fmt.Errorf("This challenger implementation uses the OptimisticTrie object, if you are not, please develop your own challenger functions")
	}
	onChainData := make(chan interface{})
	go v.ethContract.GetOnChainData(onChainData)
	stateRoot := common.Hash{}
	pendingDeposits := []optimisticrp.Deposit{}
	for methodData := range onChainData {
		switch input := methodData.(type) {
		case optimisticrp.Batch:
			v.log.Info("New onChain Batch received")
			//if there is a new batch we MUST update the stateRoot with the previous deposits (rule 1.)
			isValid, err := v.ethContract.IsStateRootValid(input.StateRoot)
			if err != nil {
				return stateRoot, err
			}
			onChainStateRoot, err := v.ethContract.GetStateRoot()
			if err != nil {
				return stateRoot, err
			}
			if isValid {
				v.log.Info("Updating accounts state as the provided batch is valid")
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
			} else if !isValid && input.StateRoot == onChainStateRoot {
				v.log.Warn("Found last submitted batch, looking for a fraud...")
			} else {
				v.log.Debug("Skipping invalid onChain batch")
			}
		case optimisticrp.Deposit:
			v.log.WithFields(logrus.Fields{"Account": input.From, "Value": input.Value}).Info("New onChain deposit")
			pendingDeposits = append(pendingDeposits, input)
		case error:
			return stateRoot, input

		default:
			return common.Hash{}, errors.New("There was an error while fetching onChain data")
		}
	}
	v.log.Info("Finished analyzing onChian data")
	return stateRoot, nil
}
