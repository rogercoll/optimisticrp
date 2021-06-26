package main

import (
	"crypto/ecdsa"
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
	privateKey, err := crypto.HexToECDSA(cmd.WithdrawerPriv)
	if err != nil {
		logger.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Fatal("Could not get public key from private key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	myaggregator := aggregator.New(tr, mybridge, privateKey, logger)
	syn, err := myaggregator.Synced()
	if err != nil {
		logger.Fatal(err)
	} else if syn == false {
		logger.Fatal("Was not able to syncronize")
	}
	logger.Info("Successfully syncronized with on-chain data")
	logger.Warn("Withdrawing all funds..")
	proof, err := tr.NewProve(fromAddress)
	if err != nil {
		logger.Fatal(err)
	}
	txOpts, err := mybridge.PrepareTxOptions(big.NewInt(0), big.NewInt(2), big.NewInt(2), privateKey)
	if err != nil {
		logger.Fatal(err)
	}
	logger.WithFields(logrus.Fields{"bytes": len(proof[2])}).Warn("Withdraw proof size")
	_, err = mybridge.Withdraw(txOpts, proof[0], proof[1], proof[2], proof[3])
	if err != nil {
		logger.Fatal(err)
	}

}
