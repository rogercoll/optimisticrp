package aggregator

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rogercoll/optimisticrp"
)

type AggregatorNode struct {
	transactions []optimisticrp.Transaction
	accountsTrie optimisticrp.Optimistic
	ethContract  optimisticrp.OptimisticSContract
}

func New(newAccountsTrie optimisticrp.Optimistic, newEthContract optimisticrp.OptimisticSContract) *AggregatorNode {
	return &AggregatorNode{
		accountsTrie: newAccountsTrie,
		ethContract:  newEthContract,
	}
}

//if sendBatch succeeds we should notify all user transactions
func (ag *AggregatorNode) sendBatch() error {
	return nil
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

//Reads all transactions to the smart contracts and computes the whole accounts trie from scratch
func (ag *AggregatorNode) computeAccountsTrie() (optimisticrp.Optimistic, common.Hash, error) {
	transactions := make(chan optimisticrp.Transaction)
	go ag.ethContract.GetAllTransactions(transactions)
	stateRoot := common.Hash{}
	for transaction := range transactions {
		fromAcc, err := ag.accountsTrie.GetAccount(transaction.From)
		if err != nil {
			return ag.accountsTrie, common.Hash{}, err
		}
		toAcc, err := ag.accountsTrie.GetAccount(transaction.To)
		switch err {
		case optimisticrp.AccountNotFound{}:
			toAcc = optimisticrp.Account{Balance: new(big.Int).SetUint64(0), Nonce: 0}
			ag.accountsTrie.UpdateAccount(transaction.To, toAcc)
		default:
			return ag.accountsTrie, common.Hash{}, err
		}
		fromAcc.Balance.Sub(fromAcc.Balance, transaction.Value)
		toAcc.Balance.Add(toAcc.Balance, transaction.Value)
		fromAcc.Nonce++
		stateRoot = ag.accountsTrie.UpdateAccount(transaction.From, fromAcc)
		stateRoot = ag.accountsTrie.UpdateAccount(transaction.To, toAcc)
	}
	return ag.accountsTrie, stateRoot, nil
}

//Should be private
func (ag *AggregatorNode) StateRoot() (common.Hash, error) {
	return ag.ethContract.GetStateRoot()
}
