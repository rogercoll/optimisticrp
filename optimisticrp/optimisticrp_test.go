package optimisticrp

import (
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var testOpr *Oprollups

var (
	address1 = common.HexToAddress("0xF93e6F80229B8Bbe14f6311b5EC1Fa36DfdEC5eB")
	address2 = common.HexToAddress("0x048C82fe2C85956Cf2872FBe32bE4AD06de3Db1E")
	address3 = common.HexToAddress("0x9185eAE1c5AD845137AaDf34a955e1D676fE421B")
	address4 = common.HexToAddress("0x3168444b98B4Bd55976137DdeEeC7A1d7BF322d3")
	address5 = common.HexToAddress("0x522fE0423db9de4e8Bb88aF3bF24aBE9B7dBF787")
)

func TestMain(m *testing.M) {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the ETH client")
	opr, err := New(common.HexToAddress("0xaf9b3894c68c73c0D5e7a2172B76E513b0008858"), client, "ff10aa6af851c1b49b7d3a94611d7823adbcfae76e153fc2757b4108a1dc402d")
	if err != nil {
		log.Fatal(err)
	}
	testOpr = opr
	//Adding testing accounts
	err = opr.AddAccount(address1)
	if err != nil {
		log.Fatal(err)
	}
	err = opr.AddAccount(address2)
	if err != nil {
		log.Fatal(err)
	}
	err = opr.AddAccount(address3)
	if err != nil {
		log.Fatal(err)
	}
	err = opr.AddAccount(address4)
	if err != nil {
		log.Fatal(err)
	}
	m.Run()
}

func TestRequiredBond(t *testing.T) {
	got := testOpr.RequiredBond
	if got.Cmp(big.NewInt(2)) != 0 {
		t.Errorf("RequiredBond = %d; want 2", got)
	}
}

func TestStateRoot(t *testing.T) {
	got := testOpr.TmpStateRoot
	if got != common.HexToHash("0x916f42eb26bd1999b4b68b3918332582bed474c337deeb8bf6b2699272c3b1d0") {
		t.Errorf("StateRoot = %d; want 0x916f42eb26bd1999b4b68b3918332582bed474c337deeb8bf6b2699272c3b1d0", got)
	}
	err := testOpr.AddAccount(address5)
	if err != nil {
		log.Fatal(err)
	}
	got2 := testOpr.TmpStateRoot
	if got2 == got {
		t.Errorf("TrieHash after adding one account must be different!")
	}
}

func TestSendBatch(t *testing.T) {
	err := testOpr.NewOptimisticTx(address1, address2, big.NewInt(1e+18), big.NewInt(5e+10))
	if err != nil {
		log.Fatal(err)
	}
	err = testOpr.NewOptimisticTx(address2, address2, big.NewInt(3e+18), big.NewInt(5e+10))
	if err != nil {
		log.Fatal(err)
	}
	err = testOpr.NewOptimisticTx(address4, address2, big.NewInt(1e+18), big.NewInt(5e+10))
	if err != nil {
		log.Fatal(err)
	}
	err = testOpr.NewOptimisticTx(address5, address2, big.NewInt(1e+18), big.NewInt(5e+10))
	if err != nil {
		log.Fatal(err)
	}
	err = testOpr.NewOptimisticTx(address2, address1, big.NewInt(1e+18), big.NewInt(5e+10))
	if err != nil {
		log.Fatal(err)
	}
	err = testOpr.NewOptimisticTx(address1, address5, big.NewInt(1e+18), big.NewInt(5e+10))
	if err != nil {
		log.Fatal(err)
	}
	err = testOpr.NewOptimisticTx(address4, address1, big.NewInt(1e+18), big.NewInt(5e+10))
	if err != nil {
		log.Fatal(err)
	}
	err = testOpr.SendBatch()
	if err != nil {
		log.Fatal(err)
	}

}
