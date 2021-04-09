package bridge

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
	log.Println(batch)
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
func (b *Bridge) GetOnChainData(dataChannel chan<- interface{}) {
	defer close(dataChannel)
	header, err := b.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		dataChannel <- err
	}
	myAbi, err := abi.JSON(strings.NewReader(abiJsonString))
	if err != nil {
		dataChannel <- err
	}
	log.Printf("Analyzing %v blocks\n", header.Number)
	for i := int64(0); i <= header.Number.Int64(); i++ {
		block, err := b.client.BlockByNumber(context.Background(), big.NewInt(i))
		if err != nil {
			dataChannel <- err
		}
		for _, tx := range block.Transactions() {
			//if tx.To() == nil => Contract creation
			if tx.To() != nil && (*(tx.To()) == b.oriAddr) {
				txReceipt, err := b.client.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					dataChannel <- err
				}
				if txReceipt.Status == 1 {
					inputData := tx.Data()
					sigdata, argdata := inputData[:4], inputData[4:]
					method, err := myAbi.MethodById(sigdata)
					if err != nil {
						dataChannel <- err
					}
					if method.Name == "newBatch" {
						log.Println("New batch transaction data detected, reading transactions")
						data, err := method.Inputs.UnpackValues(argdata)
						if err != nil {
							dataChannel <- err
						}
						batch, err := optimisticrp.UnMarshalBatch(data[0].([]byte))
						if err != nil {
							log.Println("Transaction does not contain a batch, skipping...")
							continue
						}
						dataChannel <- optimisticrp.Batch{batch.PrevStateRoot, batch.StateRoot, batch.Transactions}
					} else if method.Name == "deposit" {
						msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()))
						if err != nil {
							dataChannel <- err
						}
						dataChannel <- optimisticrp.Deposit{msg.From(), tx.Value()}
					}
				}
			}
		}
	}
	log.Printf("All blocks analized")
}

func (b *Bridge) GetPendingDeposits(depChannel chan<- interface{}) {
	defer close(depChannel)
	header, err := b.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		depChannel <- err
	}
	myAbi, err := abi.JSON(strings.NewReader(abiJsonString))
	if err != nil {
		depChannel <- err
	}
	for i := header.Number.Int64(); i >= 0; i-- {
		block, err := b.client.BlockByNumber(context.Background(), big.NewInt(i))
		if err != nil {
			depChannel <- err
		}
		for _, tx := range block.Transactions() {
			//if tx.To() == nil => Contract creation
			if tx.To() != nil && (*(tx.To()) == b.oriAddr) {
				txReceipt, err := b.client.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					depChannel <- err
				}
				//only proceed if the transaction was not reverted => valid == 1
				if txReceipt.Status == 1 {
					inputData := tx.Data()
					sigdata, _ := inputData[:4], inputData[4:]
					method, err := myAbi.MethodById(sigdata)
					if err != nil {
						depChannel <- err
					}
					if method.Name == "deposit" {
						msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()))
						if err != nil {
							log.Fatal(err)
						}
						depChannel <- optimisticrp.Deposit{msg.From(), tx.Value()}
					} else if method.Name == "newBatch" {
						depChannel <- err
					}
				}
			}
		}
	}
}

//If gasPrice == -1 => ask to the client suggested gas price
func (b *Bridge) PrepareTxOptions(value, gasLimit, gasPrice *big.Int, privKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	var err error
	if gasPrice.Cmp(big.NewInt(-1)) == 0 {
		gasPrice, err = b.client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, err
		}
	}
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := b.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privKey)
	auth.Nonce = new(big.Int).SetUint64(nonce)
	auth.Value = value              // in wei
	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}
