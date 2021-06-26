package optimisticrp

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	var acc SolidityAccount
	if len(fBytes) == 0 {
		return Account{}, &AccountNotFound{address}
	}
	err := rlp.DecodeBytes(fBytes, &acc)
	if err != nil {
		return Account{}, err
	}
	return acc.ToGolangFormat()
}

func (ot *OptimisticTrie) UpdateAccount(address common.Address, acc Account) common.Hash {
	//val := acc.MarshalBinary()
	val, err := rlp.EncodeToBytes(acc.SolidityFormat())
	if err != nil {
		panic(err)
	}
	//acc.Balance = new(big.Int).SetUint64(0e+18)
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
	it := trie.NewIterator(ot.NodeIterator(nil))
	accounts := 0
	for it.Next() {
		accounts += 1
	}
	//log.Printf("Number of accounts in Trie: %v\n", accounts)
	proof := memorydb.New()
	formatProof := [][]byte{}
	if it := trie.NewIterator(ot.NodeIterator(address.Bytes())); it.Next() && bytes.Equal(address.Bytes(), it.Key) {
		for _, p := range it.Prove() {
			formatProof = append(formatProof, p)
			proof.Put(crypto.Keccak256(p), p)
		}
	}

	toSend := make([][]byte, 4)
	//key
	toSend[0] = address.Bytes()
	//value
	toSend[1] = fBytes
	//root
	toSend[3] = ot.Hash().Bytes()
	rlpProof, err := rlp.EncodeToBytes(formatProof)
	if err != nil {
		return nil, err
	}
	//rlp proof for onchain data https://github.com/ethereum-optimism/contracts/blob/c39fcc40aec235511a5a161c3e33a6d3bd24221c/test/helpers/trie/trie-test-generator.ts#L170
	toSend[2] = rlpProof
	val, err := trie.VerifyProof(ot.StateRoot(), address.Bytes(), proof)
	if !bytes.Equal(val, fBytes) {
		return nil, fmt.Errorf("Verified value mismatch for key %x: have %x, want %x", address, val, fBytes)
	}
	return toSend, err
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

func (ot *OptimisticTrie) RemoveFunds(account common.Address, value *big.Int) error {
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
	acc.Balance.Sub(acc.Balance, value)
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
	//tx Value must be higher than the account balance (fee are not included)
	if fromAcc.Balance.Cmp(transaction.Value) == -1 {
		return common.Hash{}, &InvalidBalance{transaction.From, fromAcc.Balance}
	}
	fromAcc.Balance.Sub(fromAcc.Balance, transaction.Value)
	toAcc.Balance.Add(toAcc.Balance, transaction.Value)
	fromAcc.Nonce++
	ot.UpdateAccount(transaction.From, fromAcc)
	return ot.UpdateAccount(transaction.To, toAcc), nil
}

func (ot *OptimisticTrie) Copy() (*OptimisticTrie, error) {
	var (
		diskdb = memorydb.New()
		triedb = trie.NewDatabase(diskdb)
	)
	tr, err := trie.New(common.Hash{}, triedb)
	if err != nil {
		return nil, err
	}
	it := trie.NewIterator(ot.NodeIterator(nil))
	for it.Next() {
		tr.Update(it.Key, it.Value)
	}
	return &OptimisticTrie{tr}, nil
}
