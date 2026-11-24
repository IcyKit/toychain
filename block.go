package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"maps"
	"sync"
)

type Block struct {
	InitialBalance map[string]uint64
	Transactions   []Transaction
	PrevHash       [32]byte
	Hash           [32]byte
}

func ComputeHash(b Block) ([32]byte, error) {
	val, err := json.Marshal(struct {
		InitialBalance map[string]uint64
		Transaction    []Transaction
		PrevHash       [32]byte
	}{
		InitialBalance: b.InitialBalance,
		Transaction:    b.Transactions,
		PrevHash:       b.PrevHash,
	})
	if err != nil {
		return [32]byte{}, err
	}
	hash := sha256.Sum256(val)
	return hash, nil
}

func Verify(blocks []Block, expectedInitialBalances map[string]uint64) (*Ledger, error) {
	if len(blocks) == 0 {
		return nil, errors.New("empty blocks slice error")
	}
	isEqual := maps.Equal(blocks[0].InitialBalance, expectedInitialBalances)
	if !isEqual {
		return nil, errors.New("slices are not equal error")
	}
	expectedPrevHash := [32]byte{}
	for _, block := range blocks {
		if block.PrevHash != expectedPrevHash {
			return nil, errors.New("block previous hash and expected not equal error")
		}
		hash, err := ComputeHash(block)
		if err != nil {
			return nil, err
		}
		if hash != block.Hash {
			return nil, errors.New("hash not equal block hash")
		}
		expectedPrevHash = block.Hash
	}
	ledger := Ledger{
		Balance: maps.Clone(blocks[0].InitialBalance),
		mx:      sync.Mutex{},
	}
	for _, block := range blocks {
		for _, tx := range block.Transactions {
			if err := ledger.Apply(tx); err != nil {
				return nil, err
			}
		}
	}
	return &ledger, nil
}
