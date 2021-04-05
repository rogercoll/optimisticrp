package optimisticrp

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

type OptimisticTrie struct {
	*trie.Trie
}

func NewTrie(triedb *trie.Database) (*OptimisticTrie, error) {
	tr, err := trie.New(common.Hash{}, triedb)
	if err != nil {
		return nil, err
	}
	return &OptimisticTrie{tr}, nil
}

func (ot *OptimisticTrie) GetAccount(address common.Address) (Account, error) {
	fBytes := ot.Get(address.Bytes())
	var acc Account
	if len(fBytes) == 0 {
		return acc, &AccountNotFound{address}
	}
	_, err := acc.UnMarshalBinary(fBytes)
	if err != nil {
		return Account{}, err
	}
	return acc, nil
}

func (ot *OptimisticTrie) UpdateAccount(address common.Address, acc Account) common.Hash {
	val := acc.MarshalBinary()
	ot.Update(address.Bytes(), val)
	return ot.Hash()
}

func (ot *OptimisticTrie) StateRoot() common.Hash {
	return ot.Hash()
}

func (ot *OptimisticTrie) NewProve(address common.Address) ([][]byte, error) {
	fBytes := ot.Get(address.Bytes())
	if len(fBytes) == 0 {
		return nil, &AccountNotFound{address}
	}
	proof := memorydb.New()
	err := ot.Prove(address.Bytes(), 0, proof)
	if err != nil {
		return nil, err
	}
	toSend := make([][]byte, 4)
	//key
	toSend[0] = address.Bytes()
	//value
	toSend[1] = fBytes
	//root
	toSend[3] = ot.Hash().Bytes()
	formatProof := [][]byte{}
	it := proof.NewIterator(nil, nil)
	for it.Next() {
		formatProof = append(formatProof, it.Value())
	}
	rlpProof, err := rlp.EncodeToBytes(formatProof)
	if err != nil {
		return nil, err
	}
	//rlp proof for onchain data https://github.com/ethereum-optimism/contracts/blob/c39fcc40aec235511a5a161c3e33a6d3bd24221c/test/helpers/trie/trie-test-generator.ts#L170
	toSend[2] = rlpProof
	return toSend, nil
}
