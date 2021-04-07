package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/rogercoll/optimisticrp"
	"github.com/rogercoll/optimisticrp/bridge"
	"github.com/rogercoll/optimisticrp/verifier"
)

var addrAccount1 = common.HexToAddress("0x048C82fe2C85956Cf2872FBe32bE4AD06de3Db1E")
var addrAccount2 = common.HexToAddress("0x9185eAE1c5AD845137AaDf34a955e1D676fE421B")

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the ETH client")
	mybridge, err := bridge.New(common.HexToAddress("0x6E5145ed29Fa700f9d7c5de5F3A0Ba183926d3b9"), client)
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
	myverifier := verifier.New(tr, mybridge, privateKey)
	logs := make(chan interface{})
	go myverifier.VerifyOnChainData(logs)
	for {
		select {
		case input := <-logs:
			log.Println("new data")
			switch vlog := input.(type) {
			case error:
				log.Fatal(vlog)
			default:
				log.Println(vlog)
			}
		}
	}
}
