package verifier

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rogercoll/optimisticrp"
)

type VerifierNode struct {
	accountsTrie optimisticrp.Optimistic
	ethContract  optimisticrp.OptimisticSContract
	privKey      *ecdsa.PrivateKey
	onChainRoot  common.Hash
}

func New(newAccountsTrie optimisticrp.Optimistic, newEthContract optimisticrp.OptimisticSContract, privateKey *ecdsa.PrivateKey) *VerifierNode {
	return &VerifierNode{
		accountsTrie: newAccountsTrie,
		ethContract:  newEthContract,
		privKey:      privateKey,
	}
}

//Generate proof data to be send onchain, a proof proves that key with a certain value exits on the trie
func (v *VerifierNode) generateProof(acc common.Address) {
	_, err := v.accountsTrie.NewProve(acc)
	if err != nil {
		return
	}
}

func (v *VerifierNode) VerifyOnChainData(logs chan<- string) {

}
