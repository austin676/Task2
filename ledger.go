package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Transaction struct {
	ID           int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
}

type Ledger struct {
	Transactions []Transaction
}
func calculateHash(tx Transaction) string {

	record :=
		strconv.Itoa(tx.ID) +
			tx.Timestamp +
			tx.Data +
			tx.PreviousHash

	hash := sha256.Sum256([]byte(record))

	return hex.EncodeToString(hash[:])
}

func (l *Ledger) AddTransaction(data string) {

	var prevHash string
	var id int
    id = len(l.Transactions)

	if id == 0 {
		prevHash = "GENESIS"
	} else {
		prevHash = l.Transactions[id-1].Hash
	}

	tx := Transaction{
		ID:           id,
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Data:         data,
		PreviousHash: prevHash,
	}

	tx.Hash = calculateHash(tx)

	l.Transactions = append(l.Transactions, tx)
}
func (l *Ledger) ValidateLedger() bool {

	for i := 0; i < len(l.Transactions); i++ {

		current := l.Transactions[i]

		if calculateHash(current) != current.Hash {
			return false
		}
		if i > 0 {
			previous := l.Transactions[i-1]

			if current.PreviousHash != previous.Hash {
				return false
			}
		}
	}
	return true
}