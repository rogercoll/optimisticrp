package optimisticrp

import (
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
)

var (
	address1 = common.HexToAddress("0xF93e6F80229B8Bbe14f6311b5EC1Fa36DfdEC5eB")
	acc1     = Account{Balance: new(big.Int).SetUint64(10e+18), Nonce: 0}
)

func TestGetAccount(t *testing.T) {
	var (
		diskdb = memorydb.New()
		triedb = trie.NewDatabase(diskdb)
	)
	tr, err := NewTrie(triedb)
	if err != nil {
		log.Fatal(err)
	}
	_, err = tr.GetAccount(address1)
	switch err.(type) {
	case *AccountNotFound:
		tr.UpdateAccount(address1, acc1)
	default:
		t.Error(err)
	}
	result, err := tr.GetAccount(address1)
	if err != nil || (result.Balance.Cmp(acc1.Balance) != 0) {
		t.Errorf("Account not inserted correctly")
	}
}

func TestUpdateAccount(t *testing.T) {
	var (
		diskdb = memorydb.New()
		triedb = trie.NewDatabase(diskdb)
	)
	tr, err := NewTrie(triedb)
	if err != nil {
		log.Fatal(err)
	}
	oldStateRoot := tr.UpdateAccount(address1, acc1)
	tr.UpdateAccount(address1, Account{Balance: new(big.Int).SetUint64(2e+18), Nonce: 1})
	newValues, err := tr.GetAccount(address1)
	if err != nil {
		t.Error(err)
	}
	if oldStateRoot == tr.StateRoot() || (newValues.Balance.Cmp(acc1.Balance) == 0 && newValues.Nonce == acc1.Nonce) {
		t.Errorf("Update did not change account values")
	}

}
