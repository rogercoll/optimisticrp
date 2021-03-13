package aggregator

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rogercoll/optimisticrp"
)

type AggregatorNode struct {
	transactions []optimisticrp.Transaction
	accountsTrie optimisticrp.Optimistic
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
