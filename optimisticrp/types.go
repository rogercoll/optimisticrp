package optimisticrp

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

const MAX_TRANSACTIONS_BATCH = 10

//To, from ID in the AccountsTrie
type Transaction struct {
	Value *big.Int // wei amount
	Gas   *big.Int // gasLimit
	To    common.Address
	From  common.Address
	Nonce uint64
}

type Account struct {
	Nonce   uint64
	Balance *big.Int //weis
}

type Batch struct {
	StateRoot    common.Hash
	Transactions []Transaction
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

func (tx *Transaction) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	err := tx.encodeTyped(&buf)
	return buf.Bytes(), err
}
