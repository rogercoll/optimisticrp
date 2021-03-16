package bridge

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rogercoll/optimisticrp"
	store "github.com/rogercoll/optimisticrp/contracts"
)

var abiJsonString = `[{"inputs":[{"internalType":"uint256","name":"_lock_time","type":"uint256"},{"internalType":"uint256","name":"_required_bond","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"getMessage","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getStateRoot","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getToAddress","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"lock_time","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes","name":"_batch","type":"bytes"}],"name":"newBatch","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes","name":"_hash","type":"bytes"}],"name":"readHash","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes","name":"_hash","type":"bytes"}],"name":"readHashRLP","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"required_bond","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"stateRoot","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"}]`

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
//This implementation is used for local chains, few blocks. In production (main chain) you shall use an ingestion service to get all the transactions of a given address.
func (b *Bridge) GetOnChainData(dataChannel chan<- interface{}) error {
	defer close(dataChannel)
	header, err := b.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}
	myAbi, err := abi.JSON(strings.NewReader(abiJsonString))
	if err != nil {
		return err
	}
	for i := int64(0); i < header.Number.Int64(); i++ {
		block, err := b.client.BlockByNumber(context.Background(), big.NewInt(i))
		if err != nil {
			return err
		}
		for _, tx := range block.Transactions() {
			//if tx.To() == nil => Contract creation
			if tx.To() != nil && (*(tx.To()) == b.oriAddr) {
				inputData := tx.Data()
				sigdata, argdata := inputData[:4], inputData[4:]
				method, err := myAbi.MethodById(sigdata)
				if err != nil {
					return err
				}
				if method.Name == "newBatch" {
					log.Println("New batch transaction data detected, reading transactions")
					data, err := method.Inputs.UnpackValues(argdata)
					if err != nil {
						return err
					}
					batch, err := optimisticrp.UnMarshalBatch(data[0].([]byte))
					if err != nil {
						log.Println("Transaction does not contain a batch, skipping...")
						continue
					}
					dataChannel <- batch
				} else if method.Name == "deposit" {
					msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()))
					if err != nil {
						log.Fatal(err)
					}
					dataChannel <- optimisticrp.Deposit{msg.From(), tx.Value()}

				}
			}
		}
	}
	return nil
}
