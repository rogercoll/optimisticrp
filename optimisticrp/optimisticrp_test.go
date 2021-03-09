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
	opr, err := New(common.HexToAddress("0x98458AcD61521d93E6f708Aeb868d5DE8F4A5337"), client)
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
	log.Println(testOpr.StateRoot)
	got := testOpr.TmpStateRoot
	if got != common.HexToHash("0x781fc3bf01fa081e808611a5832f8a8ef6fef374a2be29cdbdaf9d26aa53310e") {
		t.Errorf("StateRoot = %d; want 0x781fc3bf01fa081e808611a5832f8a8ef6fef374a2be29cdbdaf9d26aa53310e", got)
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
