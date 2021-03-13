package client

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rogercoll/optimisticrp"
)

type OpClient struct {
	privKey        *ecdsa.PrivateKey
	ethAddr        common.Address
	aggregatorNode *optimisticrp.Aggregator
}

func New(hexPrivKey string, aggregator *optimisticrp.Aggregator) (*OpClient, error) {
	privateKey, err := crypto.HexToECDSA(hexPrivKey)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}
	return &OpClient{privateKey, crypto.PubkeyToAddress(*publicKeyECDSA), aggregator}, nil
}

func (client *OpClient) NewTx(from, to common.Address, value, gas *big.Int) (*optimisticrp.Transaction, error) {
	agg := client.aggregatorNode
	fnonce := (*agg).ActualNonce(from)
	tx := optimisticrp.Transaction{
		From:  from,
		To:    to,
		Value: value,
		Gas:   gas,
		Nonce: fnonce,
	}
	return &tx, nil
}

func (client *OpClient) SignTx(tx *optimisticrp.Transaction) (*optimisticrp.Transaction, error) {
	h := client.Hash(tx)
	sig, err := crypto.Sign(h[:], client.privKey)
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(client, sig)
}

func (client *OpClient) Hash(tx *optimisticrp.Transaction) common.Hash {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", tx)))
	return common.BytesToHash(h.Sum(nil))
}

func (client *OpClient) SignatureValues(sig []byte) (r, s, v *big.Int, err error) {
	if len(sig) != crypto.SignatureLength {
		return nil, nil, nil, fmt.Errorf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength)
	}
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})
	return r, s, v, nil
}

func (client *OpClient) SendTx(tx *optimisticrp.Transaction) error {
	agg := client.aggregatorNode
	return (*agg).ReceiveTransaction(*tx)
}
