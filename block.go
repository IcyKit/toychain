package main

import (
	"crypto/sha256"
	"encoding/json"
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
