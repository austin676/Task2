package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Transaction struct {
	ID           int
	Data         string
	PreviousHash string
	Hash         string
}

func NewTransaction(id int, data string, prevHash string) Transaction {
	tx := Transaction{
		ID:           id,
		Data:         data,
		PreviousHash: prevHash,
	}
	tx.Hash = calculateHash(tx)
	fmt.Printf("[LOG] Created transaction ID: %d, Hash: %s\n", tx.ID, tx.Hash)
	return tx
}

func calculateHash(tx Transaction) string {
	record := fmt.Sprintf("%d%s%s", tx.ID, tx.Data, tx.PreviousHash)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

type Ledger struct {
	Transactions []Transaction
}

func (l *Ledger) AddTransaction(data string) {
	var prevHash string
	id := len(l.Transactions)

	if id == 0 {
		prevHash = "GENESIS"
	} else {
		prevHash = l.Transactions[id-1].Hash
	}

	tx := NewTransaction(id, data, prevHash)
	l.Transactions = append(l.Transactions, tx)
}

func (l *Ledger) ValidateLedger() (bool, string) {
	fmt.Println("[LOG] Validation started...")
	defer fmt.Println("[LOG] Validation ended.")

	for i := 0; i < len(l.Transactions); i++ {
		current := l.Transactions[i]

		if calculateHash(current) != current.Hash {
			return false, fmt.Sprintf("Hash mismatch at Tx %d", current.ID)
		}

		if i > 0 {
			previous := l.Transactions[i-1]
			if current.PreviousHash != previous.Hash {
				return false, fmt.Sprintf("Broken chain at Tx %d", current.ID)
			}
		}
	}
	return true, "Valid"
}

func (l *Ledger) GetLedgerHash() string {
	if len(l.Transactions) == 0 {
		return ""
	}
	return l.Transactions[len(l.Transactions)-1].Hash
}