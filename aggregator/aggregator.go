package aggregator

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rogercoll/optimisticrp"
	"github.com/rogercoll/optimisticrp/utils"
)

type AggregatorNode struct {
	transactions []optimisticrp.Transaction
	accountsTrie optimisticrp.Optimistic
	ethContract  optimisticrp.OptimisticSContract
	privKey      *ecdsa.PrivateKey
}

func New(newAccountsTrie optimisticrp.Optimistic, newEthContract optimisticrp.OptimisticSContract, privateKey *ecdsa.PrivateKey) *AggregatorNode {
	return &AggregatorNode{
		accountsTrie: newAccountsTrie,
		ethContract:  newEthContract,
		privKey:      privateKey,
	}
}

//Sync with on-chain smart contract
func (ag *AggregatorNode) Synced() (bool, error) {
	onChainStateRoot, err := ag.onChainStateRoot()
	if err != nil {
		return false, err
	}
	if onChainStateRoot == ag.accountsTrie.StateRoot() {
		return true, nil
	}
	stateRoot, err := ag.computeAccountsTrie()
	if err != nil {
		return false, err
	}
	if stateRoot != onChainStateRoot {
		return false, fmt.Errorf("Aggregator was not able to compute a valid StateRoot")
	}
	return true, nil
}

//if sendBatch succeeds we should notify all user transactions
func (ag *AggregatorNode) sendBatch() error {
	b := optimisticrp.Batch{
		PrevStateRoot: ag.accountsTrie.StateRoot(),
	}
	for _, tx := range ag.transactions {
		_, err := ag.processTx(tx)
		if err != nil {
			return err
		}
	}
	b.StateRoot = ag.accountsTrie.StateRoot()
	b.Transactions = ag.transactions
	txOpts, err := utils.PrepareTxOptions(big.NewInt(0), big.NewInt(-1), big.NewInt(-1), ag.privKey, ag.ethContract.Client(), ag.ethContract.OriAddr())
	if err != nil {
		return err
	}
	_, err = ag.ethContract.NewBatch(b, txOpts)
	if err != nil {
		return err
	}
	ag.transactions = nil
	return err
}

func (ag *AggregatorNode) ActualNonce(acc common.Address) (uint64, error) {
	val, err := ag.accountsTrie.GetAccount(acc)
	if err != nil {
		return 0, nil
	}
	return val.Nonce, nil
}

func (ag *AggregatorNode) ReceiveTransaction(tx optimisticrp.Transaction) error {
	ag.transactions = append(ag.transactions, tx)
	if len(ag.transactions) == optimisticrp.MAX_TRANSACTIONS_BATCH {
		ag.sendBatch()
	}
	return nil
}

//Should be private
func (ag *AggregatorNode) onChainStateRoot() (common.Hash, error) {
	return ag.ethContract.GetStateRoot()
}

func (ag *AggregatorNode) processTx(transaction optimisticrp.Transaction) (common.Hash, error) {
	fromAcc, err := ag.accountsTrie.GetAccount(transaction.From)
	//fromAcc should not be added to the trie if destination addr != 0x0
	switch err.(type) {
	case *optimisticrp.AccountNotFound:
		fromAcc = optimisticrp.Account{Balance: new(big.Int).SetUint64(0), Nonce: 0}
		ag.accountsTrie.UpdateAccount(transaction.To, fromAcc)
	default:
		return common.Hash{}, err
	}
	toAcc, err := ag.accountsTrie.GetAccount(transaction.To)
	switch err.(type) {
	case *optimisticrp.AccountNotFound:
		toAcc = optimisticrp.Account{Balance: new(big.Int).SetUint64(0), Nonce: 0}
		ag.accountsTrie.UpdateAccount(transaction.To, toAcc)
	default:
		return common.Hash{}, err
	}
	fromAcc.Balance.Sub(fromAcc.Balance, transaction.Value)
	toAcc.Balance.Add(toAcc.Balance, transaction.Value)
	fromAcc.Nonce++
	ag.accountsTrie.UpdateAccount(transaction.From, fromAcc)
	return ag.accountsTrie.UpdateAccount(transaction.To, toAcc), nil
}

//Reads all transactions to the smart contracts and computes the whole accounts trie from scratch
func (ag *AggregatorNode) computeAccountsTrie() (common.Hash, error) {
	transactions := make(chan optimisticrp.Transaction)
	go ag.ethContract.GetAllTransactions(transactions)
	stateRoot := common.Hash{}
	for transaction := range transactions {
		stateRoot, err := ag.processTx(transaction)
		if err != nil {
			return stateRoot, err
		}
	}
	return stateRoot, nil
}
