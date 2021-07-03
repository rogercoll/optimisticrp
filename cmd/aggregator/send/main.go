package main

import (
	"crypto/ecdsa"
	"log"
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
	"github.com/rogercoll/optimisticrp/cmd"
	"github.com/sirupsen/logrus"
)

var addrAccount1 = common.HexToAddress(cmd.AggregatorPub)
var addrAccount2 = common.HexToAddress(cmd.WithdrawerPub)

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

func main() {
	var logger = logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Connected to the ETH client")
	mybridge, err := bridge.New(common.HexToAddress(cmd.ContractAddr), client, logger)
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
	privateKey, err := crypto.HexToECDSA(cmd.AggregatorPriv)
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
	for i := 0; i < 1; i++ {
		tx := optimisticrp.Transaction{Value: big.NewInt(1e+18), Gas: big.NewInt(1e+18), To: addrAccount2, From: addrAccount1}
		err := myaggregator.ReceiveTransaction(tx)
		if err != nil {
			logger.Fatal(err)
		}
	}
	for i := 0; i < aggregator.MAX_TRANSACTIONS_BATCH-1; i++ {
		logger.Info("Generating random receivers address to increase the trie size")
		tx := optimisticrp.Transaction{Value: big.NewInt(1e+14), Gas: big.NewInt(1e+18), To: randomAddress(), From: addrAccount1}
		err := myaggregator.ReceiveTransaction(tx)
		if err != nil {
			logger.Fatal(err)
		}
	}
	/*
		proof, err := tr.NewProve(addrAccount2)
		if err != nil {
			logger.Fatal(err)
		}
		for m, p := range proof {
			if m == 0 {
				fmt.Printf("[")
			} else {
				fmt.Printf(",[")
			}
			for n, i := range p {
				if n == 0 {
					fmt.Printf("%v", i)
				} else {
					fmt.Printf(",%v", i)
				}
			}
			fmt.Printf("]")
		}
		fmt.Println()
	*/
}
