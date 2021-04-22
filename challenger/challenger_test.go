package challenger

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/rogercoll/optimisticrp"
	"github.com/sirupsen/logrus"
)

var ver *ChallengerNode
var addrAccount1 = common.HexToAddress("0x048C82fe2C85956Cf2872FBe32bE4AD06de3Db1E")
var addrAccount2 = common.HexToAddress("0x9185eAE1c5AD845137AaDf34a955e1D676fE421B")
var addrAccount3 = common.HexToAddress("0x522fE0423db9de4e8Bb88aF3bF24aBE9B7dBF787")
var account1 = optimisticrp.Account{Balance: new(big.Int).SetUint64(0), Nonce: 0}

type mockBridge struct {
}

func (m *mockBridge) Client() *ethclient.Client          { return nil }
func (m *mockBridge) GetStateRoot() (common.Hash, error) { return common.Hash{}, nil }
func (m *mockBridge) NewBatch(optimisticrp.SolidityBatch, *bind.TransactOpts) (*types.Transaction, error) {
	return nil, nil
}
func (m *mockBridge) FraudProof(*bind.TransactOpts, []byte, []byte, []byte, []byte, optimisticrp.SolidityBatch) (*types.Transaction, error) {
	return nil, nil
}
func (m *mockBridge) Deposit(*bind.TransactOpts) (*types.Transaction, error) { return nil, nil }
func (m *mockBridge) Bond(*bind.TransactOpts) (*types.Transaction, error)    { return nil, nil }
func (m *mockBridge) Withdraw()                                              {}
func (m *mockBridge) OriAddr() common.Address                                { return common.Address{} }
func (m *mockBridge) GetPendingDeposits(depChannel chan<- interface{}) {
	defer close(depChannel)
	depChannel <- optimisticrp.Deposit{addrAccount2, big.NewInt(1e+18)}
}

func (m *mockBridge) IsStateRootValid(common.Hash) (bool, error) {
	return true, nil
}

func (m *mockBridge) PrepareTxOptions(*big.Int, *big.Int, *big.Int, *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	return nil, nil
}
func (m *mockBridge) GetOnChainData(txChannel chan<- interface{}) {
	defer close(txChannel)
	txs := []optimisticrp.Transaction{
		{
			From:  addrAccount1,
			To:    addrAccount2,
			Value: big.NewInt(1e+18),
		},
		{
			From:  addrAccount1,
			To:    addrAccount3,
			Value: big.NewInt(1e+18),
		},
	}
	txs2 := []optimisticrp.Transaction{
		{
			From:  addrAccount2,
			To:    addrAccount1,
			Value: big.NewInt(3e+18),
		},
	}
	txChannel <- optimisticrp.Deposit{addrAccount1, big.NewInt(0).SetUint64(10e+18)}
	txChannel <- optimisticrp.Batch{Transactions: txs}
	txChannel <- optimisticrp.Deposit{addrAccount3, big.NewInt(0).SetUint64(8e+18)}
	txChannel <- optimisticrp.Batch{Transactions: txs2}
}
func TestMain(m *testing.M) {
	var (
		diskdb = memorydb.New()
		triedb = trie.NewDatabase(diskdb)
	)
	tr, err := optimisticrp.NewTrie(triedb)
	if err != nil {
		log.Fatal(err)
	}
	mockBridgeContract := mockBridge{}
	ver = New(tr, &mockBridgeContract, nil, logrus.New())
	ver.accountsTrie.UpdateAccount(addrAccount1, account1)
	m.Run()
}

func TestComputeAccountsTrie(t *testing.T) {
	logs := make(chan interface{})
	go ver.VerifyOnChainData(logs)
	for {
		select {
		case input := <-logs:
			switch vlog := input.(type) {
			case error:
				got, ok := vlog.(*optimisticrp.InvalidBalance)
				if !ok {
					t.Fatal("Error not properly handled")
				}
				if got.Addr != addrAccount2 {
					t.Fatal("Wrong attacker detected")
				}
				t.SkipNow()
			default:
				t.Fatal("Invalid transaction included in the last batch and not detected!")
			}
		}
	}
}
