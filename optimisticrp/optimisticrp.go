package optimisticrp

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	"math/big"
)

type Oprollups struct {
	AccountsTrie *trie.Trie
	StateRoot    common.Hash
	//ORI Addr => Optimistic Rollups Implementation Smart Contract Address
	OriAddr string
}

func New(oriAddr string) (*Oprollups, error) {
	var (
		diskdb = memorydb.New()
		triedb = trie.NewDatabase(diskdb)
	)
	tr, err := trie.New(common.Hash{}, triedb)
	if err != nil {
		return nil, err
	}
	return &Oprollups{tr, tr.Hash(), oriAddr}, nil
}

func (opr *Oprollups) GetAccount(address common.Address) (Account, error) {
	fBytes := opr.AccountsTrie.Get(address.Bytes())
	var acc Account
	if err := rlp.DecodeBytes(fBytes, &acc); err != nil {
		return Account{}, err
	}
	return acc, nil
}

func (opr *Oprollups) UpdateAccount(address common.Address, acc Account) error {
	val, err := rlp.EncodeToBytes(acc)
	if err != nil {
		return err
	}
	opr.AccountsTrie.Update(address.Bytes(), val)
	return nil
}

//https://github.com/ethereum/go-ethereum/blob/bbfb1e4008a359a8b57ec654330c0e674623e52f/core/types/transaction.go#L68
func (opr *Oprollups) NewOptimisticTx(to, from common.Address, value, gas *big.Int) error {
	fromAcc, err := opr.GetAccount(from)
	if err != nil {
		return err
	}
	toAcc, err := opr.GetAccount(to)
	if err != nil {
		return err
	}
	fromAcc.Nonce += 1
	_ = Transaction{
		From:  from,
		To:    to,
		Value: value,
		Gas:   gas,
		Nonce: fromAcc.Nonce,
	}
	fromAcc.Balance.Sub(fromAcc.Balance, value)
	toAcc.Balance.Add(toAcc.Balance, value)
	err = opr.UpdateAccount(from, fromAcc)
	if err != nil {
		return err
	}
	err = opr.UpdateAccount(to, toAcc)
	opr.StateRoot = opr.AccountsTrie.Hash()
	return nil
}

func (opr *Oprollups) AddAccount(addr common.Address) error {
	acc := Account{Balance: big.NewInt(1e+18), Nonce: 0}
	err := opr.UpdateAccount(addr, acc)
	return err
}

//TODO: USE merkle tree indexes
func (opr *Oprollups) getAccountID(addr common.Address) uint32 {
	/*
		it := NewIterator(trie.NodeIterator(nil))
		for it.Next() {
			found[string(it.Key)] = string(it.Value)
		}
	*/
	return 8
}
