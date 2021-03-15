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
	syn, err := myaggregator.Synced()
	if err != nil {
		log.Fatal(err)
	} else if syn == false {
		log.Fatal("Was not able to syncronize")
	}
	log.Println("Successfully syncronized with on-chain data")
}
