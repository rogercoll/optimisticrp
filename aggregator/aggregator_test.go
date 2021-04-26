package aggregator

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/rogercoll/optimisticrp"
	"github.com/sirupsen/logrus"
)

var agg *AggregatorNode
var addrAccount1 = common.HexToAddress("0x048C82fe2C85956Cf2872FBe32bE4AD06de3Db1E")
var addrAccount2 = common.HexToAddress("0x9185eAE1c5AD845137AaDf34a955e1D676fE421B")
var addrAccount3 = common.HexToAddress("0x522fE0423db9de4e8Bb88aF3bF24aBE9B7dBF787")
var account1 = optimisticrp.Account{Balance: new(big.Int).SetUint64(0), Nonce: 0}

type mockBridge struct {
}

func (m *mockBridge) Client() *ethclient.Client { return nil }
func (m *mockBridge) GetStateRoot() (common.Hash, error) {
	return common.HexToHash("0x9968e894a03093c6902640366e457efb26d32ea6363cdad8c05090156bcd8587"), nil
}
func (m *mockBridge) NewBatch(optimisticrp.SolidityBatch, *bind.TransactOpts) (*types.Transaction, error) {
	return nil, nil
}
func (m *mockBridge) FraudProof(*bind.TransactOpts, []byte, []byte, []byte, []byte, optimisticrp.SolidityBatch) (*types.Transaction, error) {
	return nil, nil
}
func (m *mockBridge) Withdraw(*bind.TransactOpts, []byte, []byte, []byte, []byte) (*types.Transaction, error) {
	return nil, nil
}
func (m *mockBridge) Deposit(*bind.TransactOpts) (*types.Transaction, error) { return nil, nil }
func (m *mockBridge) Bond(*bind.TransactOpts) (*types.Transaction, error)    { return nil, nil }
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
	oneEth := big.NewInt(1e+18)
	txs := []optimisticrp.SolidityTransaction{
		{
			From:  addrAccount1,
			To:    addrAccount2,
			Value: math.U256Bytes(oneEth),
			//Value: big.NewInt(1e+18),
		},
		{
			From:  addrAccount1,
			To:    addrAccount3,
			Value: math.U256Bytes(oneEth),
		},
	}
	txs2 := []optimisticrp.SolidityTransaction{
		{
			From:  addrAccount3,
			To:    addrAccount1,
			Value: math.U256Bytes(big.NewInt(3e+18)),
		},
	}
	txChannel <- optimisticrp.Deposit{addrAccount1, big.NewInt(0).SetUint64(10e+18)}
	txChannel <- optimisticrp.SolidityBatch{Transactions: txs}
	txChannel <- optimisticrp.Deposit{addrAccount3, big.NewInt(0).SetUint64(8e+18)}
	txChannel <- optimisticrp.SolidityBatch{Transactions: txs2}
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
	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel)
	agg = New(tr, &mockBridgeContract, nil, logger)
	agg.accountsTrie.UpdateAccount(addrAccount1, account1)
	m.Run()
}

func TestActualNonce(t *testing.T) {
	got, err := agg.ActualNonce(addrAccount1)
	if err != nil {
		t.Error(err)
	}
	if got != account1.Nonce {
		t.Errorf("Nonce = %d; want %d", got, account1.Nonce)
	}
}

func TestComputeAccountsTrie(t *testing.T) {
	oldStateRoot := agg.accountsTrie.StateRoot()
	newStateRoot, _, err := agg.computeAccountsTrie()
	if err != nil {
		t.Error(err)
	}
	if newStateRoot == oldStateRoot {
		t.Errorf("NewStateRoot = %v; must be different than %v", newStateRoot, oldStateRoot)
	}
	nonceSender, err := agg.ActualNonce(addrAccount1)
	if err != nil {
		t.Error(err)
	}
	if nonceSender != 2 {
		t.Errorf("Nonce = %d; want %d", nonceSender, 2)
	}
}

func TestSendBatch(t *testing.T) {
	for i := 0; i < MAX_TRANSACTIONS_BATCH; i++ {
		tx := optimisticrp.Transaction{Value: big.NewInt(1e+18), Gas: big.NewInt(1e+18), To: addrAccount2, From: addrAccount1}
		err := agg.ReceiveTransaction(tx)
		if err != nil {
			t.Error(err)
		}
	}
}
