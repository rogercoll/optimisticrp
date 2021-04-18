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
	address2 = common.HexToAddress("0xA13e6F80229B8Bbe14f6311b5EC1Fa36DfdEC5eB")
	acc2     = Account{Balance: new(big.Int).SetUint64(3e+18), Nonce: 2}
	address3 = common.HexToAddress("0xC13e6F31259B8Bbe94f6311b5EC1Fa36DfdEC5eB")
	acc3     = Account{Balance: new(big.Int).SetUint64(1e+18), Nonce: 3}
	address4 = common.HexToAddress("0xB23e6F31259B8Bbe94f6311b5EC1Fa36DfdEC5eB")
	acc4     = Account{Balance: new(big.Int).SetUint64(1e+18), Nonce: 1}
	address5 = common.HexToAddress("0x193e6F31259B8Bbe94f6311b5EC1Fa36DfdEC5eB")
	acc5     = Account{Balance: new(big.Int).SetUint64(1e+18), Nonce: 5}
)

func TestNewProve(t *testing.T) {
	var (
		diskdb = memorydb.New()
		triedb = trie.NewDatabase(diskdb)
	)
	tr, err := NewTrie(triedb)
	if err != nil {
		log.Fatal(err)
	}
	_ = tr.UpdateAccount(address1, acc1)
	_ = tr.UpdateAccount(address2, acc2)
	_ = tr.UpdateAccount(address3, acc3)
	_ = tr.UpdateAccount(address4, acc4)
	_ = tr.UpdateAccount(address5, acc5)
	toSend, err := tr.NewProve(address3)
	if err != nil {
		t.Log(err)
	}
	if len(toSend[2]) == 0 {
		t.Error("Proof data cannot be empty")
	}
}

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
