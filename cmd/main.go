package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/rogercoll/optimisticrp"
	"github.com/rogercoll/optimisticrp/aggregator"
	"github.com/rogercoll/optimisticrp/bridge"
)

var addrAccount1 = common.HexToAddress("0x048C82fe2C85956Cf2872FBe32bE4AD06de3Db1E")
var addrAccount2 = common.HexToAddress("0x9185eAE1c5AD845137AaDf34a955e1D676fE421B")

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the ETH client")
	mybridge, err := bridge.New(common.HexToAddress("0xeA7bf969d0559f4EF1aD6645C19ED25c742a9F71"), client)
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
	privateKey, err := crypto.HexToECDSA("f99f08ed8bf9e51e1f08700cf782a86b43eab54be5f6d0911097e4f300b09e2a")
	if err != nil {
		log.Fatal(err)
	}
	myaggregator := aggregator.New(tr, mybridge, privateKey)
	syn, err := myaggregator.Synced()
	if err != nil {
		log.Fatal(err)
	} else if syn == false {
		log.Fatal("Was not able to syncronize")
	}
	log.Println("Successfully syncronized with on-chain data")
	/*
		for i := 0; i < 10; i++ {
			tx := optimisticrp.Transaction{Value: big.NewInt(1e+17), Gas: big.NewInt(1e+18), To: addrAccount2, From: addrAccount1}
			err := myaggregator.ReceiveTransaction(tx)
			if err != nil {
				log.Fatal(err)
			}
		}
	*/
}
