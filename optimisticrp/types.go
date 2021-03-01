package optimisticrp

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

//To, from ID in the AccountsTrie
type Transaction struct {
	Value uint64
	Gas   uint64
	To    uint32
	From  uint32
	Nonce uint64
}

type Account struct {
	Addr    common.Address //64-bit
	Nonce   uint64
	Balance uint64 //weis
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
