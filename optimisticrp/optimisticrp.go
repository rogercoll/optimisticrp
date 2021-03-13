package optimisticrp

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
	store "github.com/rogercoll/optimisticrp/contracts"
)

type OptimisticTrie struct {
	*trie.Trie
}

type Oprollups struct {
	aggregatorPrivKey *ecdsa.PrivateKey
	aggregatorAddr    common.Address
	ethClient         *ethclient.Client
	AccountsTrie      *trie.Trie
	StateRoot         common.Hash
	TmpStateRoot      common.Hash
	//ORI Addr => Optimistic Rollups Implementation Smart Contract Address
	OriAddr      common.Address
	ori_contract *store.Contracts
	NewBatch     Batch
	RequiredBond *big.Int
}

func (ot *OptimisticTrie) GetAccount(address common.Address) (Account, error) {
	fBytes := ot.Get(address.Bytes())
	var acc Account
	_, err := acc.UnMarshalBinary(fBytes)
	if err != nil {
		return Account{}, err
	}
	return acc, nil
}

func (ot *OptimisticTrie) UpdateAccount(address common.Address, acc Account) common.Hash {
	val := acc.MarshalBinary()
	ot.Update(address.Bytes(), val)
	return ot.Hash()
}

func New(oriAddr common.Address, ethClient *ethclient.Client, hexPrivKey string) (*Oprollups, error) {
	instance, err := store.NewContracts(oriAddr, ethClient)
	if err != nil {
		return nil, err
	}
	requiredBond, err := instance.RequiredBond(nil)
	if err != nil {
		return nil, err
	}
	var (
		diskdb = memorydb.New()
		triedb = trie.NewDatabase(diskdb)
	)
	tr, err := trie.New(common.Hash{}, triedb)
	if err != nil {
		return nil, err
	}
	onChainStateRoot, err := instance.StateRoot(nil)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.HexToECDSA(hexPrivKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	aggregatorAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Oprollups{privateKey, aggregatorAddress, ethClient, tr, onChainStateRoot, tr.Hash(), oriAddr, instance, Batch{PrevStateRoot: onChainStateRoot}, requiredBond}, nil
}

//If gasPrice == -1 => ask to the client suggested gas price
func (opr *Oprollups) prepareTxOptions(value, gasLimit, gasPrice *big.Int) (*bind.TransactOpts, error) {
	nonce, err := opr.ethClient.PendingNonceAt(context.Background(), opr.aggregatorAddr)
	if err != nil {
		return nil, err
	}
	if gasPrice.Cmp(big.NewInt(-1)) == 0 {
		gasPrice, err = opr.ethClient.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, err
		}
	}
	auth := bind.NewKeyedTransactor(opr.aggregatorPrivKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value             // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}

func (opr *Oprollups) GetAccount(address common.Address) (Account, error) {
	fBytes := opr.AccountsTrie.Get(address.Bytes())
	var acc Account
	_, err := acc.UnMarshalBinary(fBytes)
	if err != nil {
		return Account{}, err
	}
	return acc, nil
}

func (opr *Oprollups) UpdateAccount(address common.Address, acc Account) error {
	val := acc.MarshalBinary()
	opr.AccountsTrie.Update(address.Bytes(), val)
	opr.TmpStateRoot = opr.AccountsTrie.Hash()
	return nil
}

//https://github.com/ethereum/go-ethereum/blob/bbfb1e4008a359a8b57ec654330c0e674623e52f/core/types/transaction.go#L68
func (opr *Oprollups) NewOptimisticTx(to, from common.Address, value, gas *big.Int) error {
	fromAcc, err := opr.GetAccount(from)
	if err != nil {
		return err
	}
	toAcc, err := opr.GetAccount(to)
	if err != nil {
		return err
	}
	fromAcc.Nonce += 1
	tx := Transaction{
		From:  from,
		To:    to,
		Value: value,
		Gas:   gas,
		Nonce: fromAcc.Nonce,
	}
	fromAcc.Balance.Sub(fromAcc.Balance, value)
	toAcc.Balance.Add(toAcc.Balance, value)
	err = opr.UpdateAccount(from, fromAcc)
	if err != nil {
		return err
	}
	err = opr.UpdateAccount(to, toAcc)
	opr.StateRoot = opr.AccountsTrie.Hash()
	opr.NewBatch.StateRoot = opr.AccountsTrie.Hash()
	opr.NewBatch.Transactions = append(opr.NewBatch.Transactions, tx)
	return nil
}

func (opr *Oprollups) AddAccount(addr common.Address) error {
	acc := Account{Balance: new(big.Int).SetUint64(10e+18), Nonce: 0}
	err := opr.UpdateAccount(addr, acc)
	return err
}

func (opr *Oprollups) SendBatch() error {
	result, err := opr.NewBatch.MarshalBinary()
	if err != nil {
		return err
	}
	txOpts, err := opr.prepareTxOptions(big.NewInt(0), big.NewInt(-1), big.NewInt(-1))
	if err != nil {
		return err
	}
	_, err = opr.ori_contract.NewBatch(txOpts, result)
	if err != nil {
		return err
	}
	opr.NewBatch.PrevStateRoot = opr.NewBatch.StateRoot
	opr.StateRoot = opr.NewBatch.StateRoot
	return nil
}
