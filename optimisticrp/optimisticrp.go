package optimisticrp

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/trie"
)

type Oprollups struct {
	AccountsTrie *trie.Trie
}

var opr Oprollups

//https://github.com/ethereum/go-ethereum/blob/bbfb1e4008a359a8b57ec654330c0e674623e52f/core/types/transaction.go#L68
func (from *Account) NewOptimisticTx(to Account, value, gas uint64) error {
	from.Nonce += 1
	_ = Transaction{
		From:  opr.getAccountID(from.Addr),
		To:    opr.getAccountID(to.Addr),
		Value: value,
		Gas:   gas,
		Nonce: from.Nonce,
	}
	from.Balance -= value
	to.Balance += value

	return nil
}

func (opr *Oprollups) getAccountID(addr common.Address) uint32 {
	/*
		it := NewIterator(trie.NodeIterator(nil))
		for it.Next() {
			found[string(it.Key)] = string(it.Value)
		}
	*/
	return 8
}
