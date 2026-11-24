package main

import (
	"fmt"
	"maps"
	"sync"
)

var lg Ledger = Ledger{
	Balance: map[string]uint64{
		"alice": 5,
		"bob":   0,
		"katya": 0,
	},
	mx: sync.Mutex{},
}

func main() {
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
	blocks := []Block{
		{
			InitialBalance: maps.Clone(lg.Balance),
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
	fmt.Println(blocks[0].Hash)
	fmt.Println(blocks[0].InitialBalance)
	fmt.Println(blocks)
	fmt.Println(lg.Balance)
}
