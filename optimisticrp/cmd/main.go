package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rogercoll/optimisticrp"
	"log"
)

func main() {
	address1 := common.HexToAddress("0x8B503cA1beF55A904276138f2EA60906d2c58781")
	address2 := common.HexToAddress("0x048C82fe2C85956Cf2872FBe32bE4AD06de3Db1E")
	opTx := optimisticrp.Transaction{
		From: address1,
		To:   address2,
	}
	log.Println(opTx.MarshalBinary())
}
