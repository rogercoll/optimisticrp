package optimisticrp

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

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

func (tx *Transaction) encodeTyped(w *bytes.Buffer) error {
	//w.WriteByte(tx.Type())
	return rlp.Encode(w, tx)
}

func (tx *Transaction) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	err := tx.encodeTyped(&buf)
	return buf.Bytes(), err
}
