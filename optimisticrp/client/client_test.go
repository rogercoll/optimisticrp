package client

import (
	"math/big"
	"testing"

	"github.com/rogercoll/optimisticrp"
)

var (
	priv1 = "ff10aa6af851c1b49b7d3a94611d7823adbcfae76e153fc2757b4108a1dc402d"
	priv2 = "cd40c0e859b7f6ebf942ee4b2f923acbe54546e9339a025de4b173f442187828"
	priv3 = "482254ce62c1473ccbf354bf33e08d71ff09dd2859e4fb8ae08d228fb8b727a5"
	priv4 = "8d7ebc3af6553667bd122d0a8e41bee25a9191fd493ec89861f6f993f436af80"
	priv5 = "6be7af0159b0f06c078c583df4f262bffc946dbc50c550667225adf1e27b365e"
)

func TestSignTx(t *testing.T) {
	client1, err := New(priv1, nil)
	if err != nil {
		t.Error(err)
	}
	client2, err := New(priv2, nil)
	if err != nil {
		t.Error(err)
	}
	opTx := optimisticrp.Transaction{
		From:  client1.ethAddr,
		To:    client2.ethAddr,
		Value: big.NewInt(1e+18),
	}
	signedTx, err := client1.SignTx(&opTx)
	if err != nil {
		t.Error(err)
	}
	//client2 should be signing the transaction, ONLY for testing purposes
	signedTx2, err := client2.SignTx(&opTx)
	if err != nil {
		t.Error(err)
	}

	if signedTx == signedTx2 {
		t.Errorf("Signed transaction from two different clients must be different")
	}
}
