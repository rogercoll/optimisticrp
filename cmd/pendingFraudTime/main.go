package main

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
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
	time, err := mybridge.RemainingFraudPeriod()
	if err != nil {
		logger.Fatal(err)
	}
	logger.WithFields(logrus.Fields{"seconds": time}).Warn("Remaining fraud proof time")

}
