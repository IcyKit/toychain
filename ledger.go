package main

import (
	"errors"
	"sync"
)

type Transaction struct {
	From      string
	To        string
	Amount    uint64
	Signature []byte `json:",omitempty"`
}

type Ledger struct {
	Balance map[string]uint64
	mx      sync.Mutex
}

func (l *Ledger) Apply(tx Transaction) error {
	l.mx.Lock()
	defer l.mx.Unlock()

	if tx.Amount == 0 {
		return errors.New("zero amount error")
	}
	if l.Balance[tx.From] < tx.Amount {
		return errors.New("insufficient funds error")
	}
	l.Balance[tx.From] -= tx.Amount
	l.Balance[tx.To] += tx.Amount
	return nil
}
