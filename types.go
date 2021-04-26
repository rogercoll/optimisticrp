package optimisticrp

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"encoding/binary"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

const OPR_BANNER = "[OPTIMISTICRP]: "

type AccountNotFound struct {
	Addr common.Address
}

type InvalidBalance struct {
	Addr  common.Address
	Total *big.Int
}

func (i *InvalidBalance) Error() string {
	return fmt.Sprintf("%s Account %v has not enough funds (%v)", OPR_BANNER, i.Addr, i.Total)
}

func (e *AccountNotFound) Error() string {
	return fmt.Sprintf("%s Account %v was not found in the trie", OPR_BANNER, e.Addr)
}

type Optimistic interface {
	StateRoot() common.Hash
	GetAccount(common.Address) (Account, error)
	UpdateAccount(common.Address, Account) common.Hash
	NewProve(common.Address) ([][]byte, error)
}
type Signer interface {
	// SignatureValues returns the raw R, S, V values corresponding to the
	// given signature.
	SignatureValues(sig []byte) (r, s, v *big.Int, err error)
	// Hash returns 'signature hash', i.e. the transaction hash that is signed by the
	// private key. This hash does not uniquely identify the transaction.
	Hash(tx *Transaction) common.Hash
}

type Aggregator interface {
	//Synced returns if the Aggregator is syncronized with the on-chain data or not
	Synced() (bool, error)
	ReceiveTransaction(tx Transaction) error
	ActualNonce(acc common.Address) uint64
}

type OptimisticSContract interface {
	OriAddr() common.Address
	GetStateRoot() (common.Hash, error)
	GetOnChainData(chan<- interface{})
	GetPendingDeposits(chan<- interface{})
	IsStateRootValid(common.Hash) (bool, error)
	PrepareTxOptions(*big.Int, *big.Int, *big.Int, *ecdsa.PrivateKey) (*bind.TransactOpts, error)
	NewBatch(SolidityBatch, *bind.TransactOpts) (*types.Transaction, error)
	FraudProof(*bind.TransactOpts, []byte, []byte, []byte, []byte, SolidityBatch) (*types.Transaction, error)
	Bond(*bind.TransactOpts) (*types.Transaction, error)
	Deposit(*bind.TransactOpts) (*types.Transaction, error)
	Withdraw(*bind.TransactOpts, []byte, []byte, []byte, []byte) (*types.Transaction, error)
	Client() *ethclient.Client
}

//Common types
type Deposit struct {
	From  common.Address
	Value *big.Int
}

type Withdraw struct {
	From  common.Address
	Value *big.Int
}

//To, from ID in the AccountsTrie
type Transaction struct {
	Value   *big.Int // wei amount
	Gas     *big.Int // gasLimit
	To      common.Address
	From    common.Address
	Nonce   uint64
	V, R, S *big.Int // signature values
}

type SolidityTransaction struct {
	Value   []byte // wei amount
	Gas     []byte // gasLimit
	To      common.Address
	From    common.Address
	Nonce   uint64
	V, R, S *big.Int // signature values TODO => make signature verificable on-chain
}

type Account struct {
	Nonce   uint64
	Balance *big.Int //weis
}

type SolidityAccount struct {
	Nonce   uint64
	Balance []byte
}

type Batch struct {
	PrevStateRoot common.Hash
	StateRoot     common.Hash
	Transactions  []Transaction
}

type SolidityBatch struct {
	PrevStateRoot common.Hash
	StateRoot     common.Hash
	Transactions  []SolidityTransaction
}

func (tx *Transaction) encodeTyped(w *bytes.Buffer) error {
	//w.WriteByte(tx.Type())
	return rlp.Encode(w, tx)
}

func (bt *Batch) encodeTyped(w *bytes.Buffer) error {
	//w.WriteByte(tx.Type())
	return rlp.Encode(w, bt)
}

func (bt *Batch) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	err := bt.encodeTyped(&buf)
	return buf.Bytes(), err
}

func UnMarshalBatch(b []byte) (*Batch, error) {
	var data Batch
	err := rlp.DecodeBytes(b, &data)
	return &data, err
}

func (tx *Transaction) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	err := tx.encodeTyped(&buf)
	return buf.Bytes(), err
}

func (account *Account) MarshalBinary() []byte {
	//Uint64 will occupy a byte array of length 8
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, account.Nonce)
	//big.Int.Bytes returns BigEndian array
	bb := account.Balance.Bytes()
	return append(b, bb...)
}

func (account *Account) SolidityFormat() interface{} {
	return SolidityAccount{Nonce: account.Nonce, Balance: math.U256Bytes(account.Balance)}
}

func (account *SolidityAccount) ToGolangFormat() (Account, error) {
	return Account{Nonce: account.Nonce, Balance: new(big.Int).SetBytes(account.Balance)}, nil
}

func (b *Batch) SolidityFormat() SolidityBatch {
	sb := SolidityBatch{
		PrevStateRoot: b.PrevStateRoot,
		StateRoot:     b.StateRoot,
	}
	for _, tx := range b.Transactions {
		sb.Transactions = append(sb.Transactions, SolidityTransaction{
			Gas:   math.U256Bytes(tx.Gas),
			Value: math.U256Bytes(tx.Value),
			To:    tx.To,
			From:  tx.From,
			Nonce: tx.Nonce,
			V:     tx.V,
			R:     tx.R,
			S:     tx.S,
		})
	}
	return sb
}

func (sb *SolidityBatch) ToGolangFormat() (Batch, error) {
	b := Batch{
		PrevStateRoot: sb.PrevStateRoot,
		StateRoot:     sb.StateRoot,
	}
	for _, tx := range sb.Transactions {
		b.Transactions = append(b.Transactions, Transaction{
			Gas:   new(big.Int).SetBytes(tx.Gas),
			Value: new(big.Int).SetBytes(tx.Value),
			To:    tx.To,
			From:  tx.From,
			Nonce: tx.Nonce,
			V:     tx.V,
			R:     tx.R,
			S:     tx.S,
		})
	}
	return b, nil
}

func (account *Account) UnMarshalBinary(abytes []byte) (*Account, error) {
	b := abytes[:8]
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.BigEndian, &account.Nonce)
	if err != nil {
		return nil, err
	}
	i := new(big.Int)
	i.SetBytes(abytes[8:])
	account.Balance = i
	return account, nil
}

func (tx *Transaction) setSignatureValues(v, r, s *big.Int) {
	tx.V = v
	tx.R = r
	tx.S = s
}

func (tx *Transaction) copy() *Transaction {
	return &Transaction{
		Value: tx.Value,
		Gas:   tx.Gas,
		To:    tx.To,
		From:  tx.From,
		Nonce: tx.Nonce,
		V:     new(big.Int),
		R:     new(big.Int),
		S:     new(big.Int),
	}
}

func (tx *Transaction) WithSignature(signer Signer, sig []byte) (*Transaction, error) {
	r, s, v, err := signer.SignatureValues(sig)
	if err != nil {
		return nil, err
	}
	signedTx := tx.copy()
	signedTx.setSignatureValues(v, r, s)
	return signedTx, nil
}
