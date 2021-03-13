package optimisticrp

import (
	"math/big"
	"testing"
)

var Accounts = []Account{
	{1, big.NewInt(2)},
	{3, big.NewInt(4)},
}

func TestMarshallAccount(t *testing.T) {
	for _, acc := range Accounts {
		b := acc.MarshalBinary()
		emptyAcc := Account{}
		_, err := emptyAcc.UnMarshalBinary(b)
		if err != nil {
			t.Error(err)
		}
		if emptyAcc.Nonce != acc.Nonce || (emptyAcc.Balance.Cmp(acc.Balance) != 0) {
			t.Errorf("Nonces must be equal")
		}
	}
}
