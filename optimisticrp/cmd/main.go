package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rogercoll/optimisticrp"
	"log"
	"math/big"
)

func main() {
	opr, err := optimisticrp.New("hello")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("State Root: %v", opr.StateRoot)
	address1 := common.HexToAddress("0x8B503cA1beF55A904276138f2EA60906d2c58781")
	err = opr.AddAccount(address1)
	if err != nil {
		log.Fatal(err)
	}
	address2 := common.HexToAddress("0x048C82fe2C85956Cf2872FBe32bE4AD06de3Db1E")
	err = opr.AddAccount(address2)
	if err != nil {
		log.Fatal(err)
	}
	/*
		opTx := optimisticrp.Transaction{
			From: address1,
			To:   address2,
		}
		log.Println(opTx.MarshalBinary())
	*/
	//Sending 1 eth
	err = opr.NewOptimisticTx(address1, address2, big.NewInt(1e+18), big.NewInt(5e+10))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("State Root: %v", opr.StateRoot)
	acc1, _ := opr.GetAccount(address1)
	log.Printf("Account Balance: %v", acc1.Balance)
	acc2, _ := opr.GetAccount(address2)
	log.Printf("Account Balance: %v", acc2.Balance)
	opr.SendBatch()

	r, _ := optimisticrp.Encode(opr.StateRoot)
	log.Println(r)
	tr, _ := optimisticrp.Encode(opr.NewBatch.Transactions[0])
	log.Println(tr)
	tob, _ := optimisticrp.Encode(opr.NewBatch.Transactions[0].To)
	log.Println(tob)
}
