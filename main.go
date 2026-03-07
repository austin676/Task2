package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
reader := bufio.NewReader(os.Stdin)
	ledger := Ledger{}

	for {

		fmt.Println("\n------ Deterministic Ledger CLI ------")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. View Ledger")
		fmt.Println("3. Validate Ledger")
		fmt.Println("4. Corrupt Ledger")
		fmt.Println("5. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {

		case 1:

			fmt.Println("Enter transaction data:")
			data, _ := reader.ReadString('\n')
			data = strings.TrimSpace(data)

			ledger.AddTransaction(data)

			fmt.Println("Transaction Added")

		case 2:

			if len(ledger.Transactions) == 0 {
				fmt.Println("Ledger is empty")
				break
			}

			fmt.Println("\n--- Ledger Transactions ---")
			fmt.Println()
			for _, tx := range ledger.Transactions {

				fmt.Printf("ID: %d\n", tx.ID)
				fmt.Printf("Timestamp: %s\n", tx.Timestamp)
				fmt.Printf("Data: %s\n", tx.Data)
				fmt.Printf("Hash: %s\n", tx.Hash)
				fmt.Printf("PrevHash: %s\n", tx.PreviousHash)
				fmt.Println()
			}

		case 3:

			if ledger.ValidateLedger() {
				fmt.Println("Ledger VALID")
			} else {
				fmt.Println("Ledger INVALID")
			}

		case 4:

			if len(ledger.Transactions) > 0 {
				ledger.Transactions[0].Data = "HACKED DATA"
				fmt.Println("Ledger has been intentionally corrupted")
			} else {
				fmt.Println("No transactions to corrupt")
			}

		case 5:
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}