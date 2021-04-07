package optimisticrp

import (
	"math/big"

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

//Additional helpers not linked to interface so you can use them as you wish
func (ot *OptimisticTrie) AddFunds(account common.Address, value *big.Int) error {
	acc, err := ot.GetAccount(account)
	switch err.(type) {
	case nil:
	case *AccountNotFound:
		newAcc := Account{Balance: value, Nonce: 0}
		ot.UpdateAccount(account, newAcc)
		return nil
	default:
		return err
	}
	acc.Balance.Add(acc.Balance, value)
	ot.UpdateAccount(account, acc)
	return nil
}

func (ot *OptimisticTrie) ProcessTx(transaction Transaction) (common.Hash, error) {
	fromAcc, err := ot.GetAccount(transaction.From)
	if err != nil {
		return common.Hash{}, err
	}
	toAcc, err := ot.GetAccount(transaction.To)
	switch err.(type) {
	case nil:
	case *AccountNotFound:
		toAcc = Account{Balance: new(big.Int).SetUint64(0), Nonce: 0}
		ot.UpdateAccount(transaction.To, toAcc)
	default:
		return common.Hash{}, err
	}
	if fromAcc.Balance.Cmp(transaction.Value) == -1 {
		return common.Hash{}, &InvalidBalance{transaction.From, fromAcc.Balance}
	}
	fromAcc.Balance.Sub(fromAcc.Balance, transaction.Value)
	toAcc.Balance.Add(toAcc.Balance, transaction.Value)
	fromAcc.Nonce++
	ot.UpdateAccount(transaction.From, fromAcc)
	return ot.UpdateAccount(transaction.To, toAcc), nil
}
