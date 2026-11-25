package main

import (
	"crypto/ed25519"
	"errors"
	"fmt"
	"os"
)

func main() {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	transaction := Transaction{
		From:   "alice",
		To:     "bob",
		Amount: 3,
	}
	bytes, err := signingBytes(transaction)
	if err != nil {
		fmt.Println(err)
		return
	}
	transaction.Signature = ed25519.Sign(privateKey, bytes)
	verifiedSign := ed25519.Verify(publicKey, bytes, transaction.Signature)
	fmt.Println(verifiedSign)
	blocks, err := LoadBlocks("chain.json")
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			fmt.Println(err)
			return
		}
	} else {
		lg, err := Verify(blocks, map[string]uint64{
			"alice": 5,
			"bob":   0,
			"katya": 0,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(lg.Balance)
		return
	}
	lg := Ledger{
		Balance: map[string]uint64{
			"alice": 5,
			"bob":   0,
			"katya": 0,
		},
	}

	blocks = []Block{
		{
			InitialBalance: map[string]uint64{
				"alice": 5,
				"bob":   0,
				"katya": 0,
			},
		},
	}

	tx := []Transaction{
		{
			From:   "alice",
			To:     "bob",
			Amount: 3,
		},
		{
			From:   "bob",
			To:     "katya",
			Amount: 1,
		},
	}
	hash, err := ComputeHash(blocks[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	blocks[0].Hash = hash
	for _, t := range tx {
		if err := lg.Apply(t); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Успешный перевод %v -> %v, сумма: %d\n", t.From, t.To, t.Amount)
		lastBlock := blocks[len(blocks)-1]
		newBlock := Block{
			Transactions: []Transaction{
				t,
			},
			PrevHash: lastBlock.Hash,
		}
		hash, err := ComputeHash(newBlock)
		if err != nil {
			fmt.Println(err)
			return
		}
		newBlock.Hash = hash
		blocks = append(blocks, newBlock)
	}

	err = SaveBlocks("chain.json", blocks)
	if err != nil {
		fmt.Println(err)
		return
	}
}
