package utils

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//If gasPrice == -1 => ask to the client suggested gas price
func PrepareTxOptions(value, gasLimit, gasPrice *big.Int, privKey *ecdsa.PrivateKey, client *ethclient.Client, oriAddr common.Address) (*bind.TransactOpts, error) {
	nonce, err := client.PendingNonceAt(context.Background(), oriAddr)
	if err != nil {
		return nil, err
	}
	if gasPrice.Cmp(big.NewInt(-1)) == 0 {
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, err
		}
	}
	auth := bind.NewKeyedTransactor(privKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value             // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}
