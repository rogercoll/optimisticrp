package optimisticrp

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/trie"
)

type OptimisticTrie struct {
	*trie.Trie
}

func NewTrie(triedb *trie.Database) (*OptimisticTrie, error) {
	tr, err := trie.New(common.Hash{}, triedb)
	if err != nil {
		return nil, err
	}
	return &OptimisticTrie{tr}, nil
}

func (ot *OptimisticTrie) GetAccount(address common.Address) (Account, error) {
	fBytes := ot.Get(address.Bytes())
	var acc Account
	if len(fBytes) == 0 {
		return acc, &AccountNotFound{}
	}
	_, err := acc.UnMarshalBinary(fBytes)
	if err != nil {
		return Account{}, err
	}
	return acc, nil
}

func (ot *OptimisticTrie) UpdateAccount(address common.Address, acc Account) common.Hash {
	val := acc.MarshalBinary()
	ot.Update(address.Bytes(), val)
	return ot.Hash()
}

func (ot *OptimisticTrie) StateRoot() common.Hash {
	return ot.Hash()
}

//https://github.com/ethereum/go-ethereum/blob/bbfb1e4008a359a8b57ec654330c0e674623e52f/core/types/transaction.go#L68
