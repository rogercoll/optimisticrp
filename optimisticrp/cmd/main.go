package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/rogercoll/optimisticrp"
	"github.com/rogercoll/optimisticrp/aggregator"
	"github.com/rogercoll/optimisticrp/bridge"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the ETH client")
	mybridge, err := bridge.New(common.HexToAddress("0xaf9b3894c68c73c0D5e7a2172B76E513b0008858"), client)
	if err != nil {
		log.Fatal(err)
	}
	//TESTING A CLIENT/VERIFIER
	log.Println(mybridge.GetStateRoot())
	//TESTING A NEW AGGREGATOR
	var (
		diskdb = memorydb.New()
		triedb = trie.NewDatabase(diskdb)
	)
	tr, err := optimisticrp.NewTrie(triedb)
	if err != nil {
		log.Fatal(err)
	}
	myaggregator := aggregator.New(tr, mybridge, nil)
	log.Println(myaggregator.StateRoot())
	/*
		opr, err := optimisticrp.New("hello")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("State Root: %v", opr.StateRoot)
		address1 := common.HexToAddress("0x8B503cA1beF55A904276138f2EA60906d2c58781")
		err = opr.AddAccount(address1)
		if err != nil {
			log.Fatal(err)
		}
		address2 := common.HexToAddress("0x048C82fe2C85956Cf2872FBe32bE4AD06de3Db1E")
		err = opr.AddAccount(address2)
		if err != nil {
			log.Fatal(err)
		}

			opTx := optimisticrp.Transaction{
				From: address1,
				To:   address2,
			}
			log.Println(opTx.MarshalBinary())

		//Sending 1 eth
		err = opr.NewOptimisticTx(address1, address2, big.NewInt(1e+18), big.NewInt(5e+10))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("State Root: %v", opr.StateRoot)
		acc1, _ := opr.GetAccount(address1)
		log.Printf("Account Balance: %v", acc1.Balance)
		acc2, _ := opr.GetAccount(address2)
		log.Printf("Account Balance: %v", acc2.Balance)
		opr.SendBatch()

		r, _ := optimisticrp.Encode(opr.StateRoot)
		log.Println(r)
		tr, _ := optimisticrp.Encode(opr.NewBatch.Transactions[0])
		log.Println(tr)
		tob, _ := optimisticrp.Encode(opr.NewBatch.Transactions[0].To)
		log.Println(tob)
	*/
}
