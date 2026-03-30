package main

import (
	"fmt"
)

// ReplayLedger recomputes hashes from scratch and validates the chain mapping
func ReplayLedger(txs []Transaction) (bool, string) {
	fmt.Println("[LOG] Replay started...")
	defer fmt.Println("[LOG] Replay ended.")

	if len(txs) == 0 {
		return true, "Empty ledger"
	}

	for i := 0; i < len(txs); i++ {
		current := txs[i]

		// Recompute hash from scratch and compare
		expectedHash := calculateHash(current)
		if expectedHash != current.Hash {
			return false, fmt.Sprintf("Hash mismatch at Tx %d", current.ID)
		}

		// Validate previous hash linkage
		if i == 0 {
			if current.PreviousHash != "GENESIS" {
				return false, fmt.Sprintf("Broken chain at Tx %d", current.ID)
			}
		} else {
			previous := txs[i-1]
			if current.PreviousHash != previous.Hash {
				return false, fmt.Sprintf("Broken chain at Tx %d", current.ID)
			}
		}
	}

	return true, "Valid replay"
}
