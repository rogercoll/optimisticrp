package optimisticrp

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

func randomAddress() common.Address {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	return common.HexToAddress(crypto.PubkeyToAddress(*publicKeyECDSA).Hex())
}

func TestNewProve(t *testing.T) {
	var (
		diskdb = memorydb.New()
		triedb = trie.NewDatabase(diskdb)
	)
	tr, err := NewTrie(triedb)
	if err != nil {
		log.Fatal(err)
	}
	_ = tr.UpdateAccount(address3, acc3)
	toSend, err := tr.NewProve(address3)
	if err != nil {
		t.Log(err)
	}
	fmt.Printf("%v,%v\n", 1, len(toSend[2]))
	_ = tr.UpdateAccount(randomAddress(), acc1)
	toSend, err = tr.NewProve(address3)
	if err != nil {
		t.Log(err)
	}
	fmt.Printf("%v,%v\n", 2, len(toSend[2]))
	_ = tr.UpdateAccount(randomAddress(), acc2)
	toSend, err = tr.NewProve(address3)
	if err != nil {
		t.Log(err)
	}
	fmt.Printf("%v,%v\n", 3, len(toSend[2]))
	_ = tr.UpdateAccount(randomAddress(), acc3)
	toSend, err = tr.NewProve(address3)
	if err != nil {
		t.Log(err)
	}
	fmt.Printf("%v,%v\n", 4, len(toSend[2]))
	_ = tr.UpdateAccount(randomAddress(), acc4)
	toSend, err = tr.NewProve(address3)
	if err != nil {
		t.Log(err)
	}
	fmt.Printf("%v,%v\n", 5, len(toSend[2]))
	for i := 5; i <= 100; i += 5 {
		_ = tr.UpdateAccount(randomAddress(), acc1)
		_ = tr.UpdateAccount(randomAddress(), acc2)
		_ = tr.UpdateAccount(randomAddress(), acc3)
		_ = tr.UpdateAccount(randomAddress(), acc4)
		_ = tr.UpdateAccount(randomAddress(), acc5)
		toSend, err := tr.NewProve(address3)
		if err != nil {
			t.Log(err)
		}
		fmt.Printf("%v,%v\n", i, len(toSend[2]))
	}
	for i := 200; i <= 10000; i += 100 {
		for j := 0; j < 100; j++ {
			_ = tr.UpdateAccount(randomAddress(), acc1)
		}
		toSend, err := tr.NewProve(address3)
		if err != nil {
			t.Log(err)
		}
		fmt.Printf("%v,%v\n", i, len(toSend[2]))
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
