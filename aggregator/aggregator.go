package aggregator

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rogercoll/optimisticrp"
	"github.com/rogercoll/optimisticrp/utils"
)

type AggregatorNode struct {
	transactions []optimisticrp.Transaction
	accountsTrie optimisticrp.Optimistic
	ethContract  optimisticrp.OptimisticSContract
	privKey      *ecdsa.PrivateKey
}

func New(newAccountsTrie optimisticrp.Optimistic, newEthContract optimisticrp.OptimisticSContract, privateKey *ecdsa.PrivateKey) *AggregatorNode {
	return &AggregatorNode{
		accountsTrie: newAccountsTrie,
		ethContract:  newEthContract,
		privKey:      privateKey,
	}
}

//Sync with on-chain smart contract
func (ag *AggregatorNode) Synced() (bool, error) {
	onChainStateRoot, err := ag.onChainStateRoot()
	if err != nil {
		return false, err
	}
	if onChainStateRoot == ag.accountsTrie.StateRoot() {
		return true, nil
	}
	stateRoot, err := ag.computeAccountsTrie()
	if err != nil {
		return false, err
	}
	if stateRoot != onChainStateRoot {
		return false, fmt.Errorf("Aggregator was not able to compute a valid StateRoot")
	}
	return true, nil
}

//if sendBatch succeeds we should notify all user transactions
func (ag *AggregatorNode) sendBatch() error {
	b := optimisticrp.Batch{
		PrevStateRoot: ag.accountsTrie.StateRoot(),
	}
	for _, tx := range ag.transactions {
		_, err := ag.processTx(tx)
		if err != nil {
			return err
		}
	}
	b.StateRoot = ag.accountsTrie.StateRoot()
	b.Transactions = ag.transactions
	txOpts, err := utils.PrepareTxOptions(big.NewInt(0), big.NewInt(-1), big.NewInt(-1), ag.privKey, ag.ethContract.Client(), ag.ethContract.OriAddr())
	if err != nil {
		return err
	}
	_, err = ag.ethContract.NewBatch(b, txOpts)
	if err != nil {
		return err
	}
	ag.transactions = nil
	return err
}

func (ag *AggregatorNode) ActualNonce(acc common.Address) (uint64, error) {
	val, err := ag.accountsTrie.GetAccount(acc)
	if err != nil {
		return 0, nil
	}
	return val.Nonce, nil
}

func (ag *AggregatorNode) ReceiveTransaction(tx optimisticrp.Transaction) error {
	ag.transactions = append(ag.transactions, tx)
	if len(ag.transactions) == optimisticrp.MAX_TRANSACTIONS_BATCH {
		ag.sendBatch()
	}
	return nil
}

//Should be private
func (ag *AggregatorNode) onChainStateRoot() (common.Hash, error) {
	return ag.ethContract.GetStateRoot()
}

func (ag *AggregatorNode) processTx(transaction optimisticrp.Transaction) (common.Hash, error) {
	fromAcc, err := ag.accountsTrie.GetAccount(transaction.From)
	//fromAcc should not be added to the trie if destination addr != 0x0
	switch err.(type) {
	case *optimisticrp.AccountNotFound:
		fromAcc = optimisticrp.Account{Balance: new(big.Int).SetUint64(0), Nonce: 0}
		ag.accountsTrie.UpdateAccount(transaction.From, fromAcc)
	case nil:
	default:
		return common.Hash{}, err
	}
	toAcc, err := ag.accountsTrie.GetAccount(transaction.To)
	switch err.(type) {
	case nil:
	case *optimisticrp.AccountNotFound:
		toAcc = optimisticrp.Account{Balance: new(big.Int).SetUint64(0), Nonce: 0}
		ag.accountsTrie.UpdateAccount(transaction.To, toAcc)
	default:
		return common.Hash{}, err
	}
	log.Println(fromAcc)
	fromAcc.Balance.Sub(fromAcc.Balance, transaction.Value)
	toAcc.Balance.Add(toAcc.Balance, transaction.Value)
	fromAcc.Nonce++
	ag.accountsTrie.UpdateAccount(transaction.From, fromAcc)
	log.Println(fromAcc)
	return ag.accountsTrie.UpdateAccount(transaction.To, toAcc), nil
}

func (ag *AggregatorNode) addFunds(account common.Address, value *big.Int) error {
	acc, err := ag.accountsTrie.GetAccount(account)
	switch err.(type) {
	case nil:
	case *optimisticrp.AccountNotFound:
		newAcc := optimisticrp.Account{Balance: value, Nonce: 0}
		ag.accountsTrie.UpdateAccount(account, newAcc)
		return nil
	default:
		return err
	}
	acc.Balance.Add(acc.Balance, value)
	ag.accountsTrie.UpdateAccount(account, acc)
	return nil
}

//Reads all transactions to the smart contracts and computes the whole accounts trie from scratch
func (ag *AggregatorNode) computeAccountsTrie() (common.Hash, error) {
	onChainData := make(chan interface{})
	go ag.ethContract.GetOnChainData(onChainData)
	stateRoot := common.Hash{}
	pendingDeposits := []optimisticrp.Deposit{}
	for methodData := range onChainData {
		switch input := methodData.(type) {
		case optimisticrp.Batch:
			log.Println("New batch recevied")
			//if there is a new batch we MUST update the stateRoot with the previous deposits (rule 1.)
			for _, deposit := range pendingDeposits {
				err := ag.addFunds(deposit.From, deposit.Value)
				if err != nil {
					return stateRoot, err
				}
			}
			pendingDeposits = nil
			for _, txInBatch := range input.Transactions {
				stateRoot, err := ag.processTx(txInBatch)
				if err != nil {
					return stateRoot, err
				}
			}
		case optimisticrp.Deposit:
			log.Printf("New deposit recevied from %v\n", input.From)
			pendingDeposits = append(pendingDeposits, input)
		default:
			log.Println("On chain data could not be mapped to any data type")
		}
	}
	return stateRoot, nil
}
