package main

import "encoding/json"

func signingBytes(tx Transaction) ([]byte, error) {
	marshStruct := struct {
		From   string
		To     string
		Amount uint64
	}{
		From:   tx.From,
		To:     tx.To,
		Amount: tx.Amount,
	}
	txJson, err := json.Marshal(marshStruct)
	if err != nil {
		return nil, err
	}
	return txJson, err
}
