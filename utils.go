package optimisticrp

import (
	"github.com/ethereum/go-ethereum/rlp"
)

func Encode(i interface{}) (result []byte, err error) {
	result, err = rlp.EncodeToBytes(i)
	return
}
