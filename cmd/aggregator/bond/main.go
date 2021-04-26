package main

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
	privateKey, err := crypto.HexToECDSA(cmd.AggregatorPriv)
	if err != nil {
		logger.Fatal(err)
	}
	txOpts, err := mybridge.PrepareTxOptions(new(big.Int).SetUint64(15e+17), big.NewInt(2), big.NewInt(2), privateKey)
	if err != nil {
		logger.Fatal(err)
	}
	_, err = mybridge.Bond(txOpts)
	if err != nil {
		logger.Fatal(err)
	}
}
