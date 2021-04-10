package main

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/rogercoll/optimisticrp"
	"github.com/rogercoll/optimisticrp/aggregator"
	"github.com/rogercoll/optimisticrp/bridge"
	"github.com/sirupsen/logrus"
)

var addrAccount1 = common.HexToAddress("0x048C82fe2C85956Cf2872FBe32bE4AD06de3Db1E")
var addrAccount2 = common.HexToAddress("0x9185eAE1c5AD845137AaDf34a955e1D676fE421B")

func main() {
	var logger = logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Connected to the ETH client")
	mybridge, err := bridge.New(common.HexToAddress("0x84Cb561d6cDd8b3697004303e5cda2f7a84b057B"), client, logger)
	if err != nil {
		logger.Fatal(err)
	}
	var (
		diskdb = memorydb.New()
		triedb = trie.NewDatabase(diskdb)
	)
	tr, err := optimisticrp.NewTrie(triedb)
	if err != nil {
		logger.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("ff10aa6af851c1b49b7d3a94611d7823adbcfae76e153fc2757b4108a1dc402d")
	if err != nil {
		logger.Fatal(err)
	}
	myaggregator := aggregator.New(tr, mybridge, privateKey, logger)
	syn, err := myaggregator.Synced()
	if err != nil {
		logger.Fatal(err)
	} else if syn == false {
		logger.Fatal("Was not able to syncronize")
	}
	logger.Info("Successfully syncronized with on-chain data")
	for i := 0; i < 10; i++ {
		tx := optimisticrp.Transaction{Value: big.NewInt(1e+17), Gas: big.NewInt(1e+18), To: addrAccount2, From: addrAccount1}
		err := myaggregator.ReceiveTransaction(tx)
		if err != nil {
			logger.Fatal(err)
		}
	}
}
