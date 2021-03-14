package bridge

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rogercoll/optimisticrp"
	store "github.com/rogercoll/optimisticrp/contracts"
)

type Bridge struct {
	oriContract *store.Contracts
	oriAddr     common.Address
	client      *ethclient.Client
}

func New(oriAddr common.Address, ethClient *ethclient.Client) (*Bridge, error) {
	instance, err := store.NewContracts(oriAddr, ethClient)
	if err != nil {
		return nil, err
	}
	return &Bridge{instance, oriAddr, ethClient}, nil
}

func (b *Bridge) Client() *ethclient.Client {
	return b.client
}

func (b *Bridge) GetStateRoot() (common.Hash, error) {
	onChainStateRoot, err := b.oriContract.StateRoot(nil)
	if err != nil {
		return common.Hash{}, err
	}
	return onChainStateRoot, nil
}

func (b *Bridge) NewBatch(batch optimisticrp.Batch, txOpts *bind.TransactOpts) (*types.Transaction, error) {
	result, err := batch.MarshalBinary()
	if err != nil {
		return nil, err
	}
	txresult, err := b.oriContract.NewBatch(txOpts, result)
	if err != nil {
		return nil, err
	}
	return txresult, nil
}

func (b *Bridge) OriAddr() common.Address {
	return b.oriAddr
}

func (b *Bridge) FraudProof() {

}
func (b *Bridge) Bond() {

}
func (b *Bridge) Withdraw() {

}

//Reads all transactions to the smart contracts and computes the whole accounts trie from scratch

func (b *Bridge) GetAllTransactions(chan<- optimisticrp.Transaction) error {
	/*
		block, err := b.client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			log.Fatal(err)
		}
	*/
	return nil
}
