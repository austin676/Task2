package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		fmt.Println("5. Replay Ledger")
		fmt.Println("6. Get Ledger Hash")
		fmt.Println("7. Determinism Proof")
		fmt.Println("8. Replay Proof")
		fmt.Println("9. Exit")

		fmt.Print("Choice: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid choice, try again.")
			continue
		}

		switch choice {

		case 1:
			fmt.Print("Enter transaction data: ")
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
			for _, tx := range ledger.Transactions {
				fmt.Printf("ID: %d\n", tx.ID)
				fmt.Printf("Data: %s\n", tx.Data)
				fmt.Printf("Hash: %s\n", tx.Hash)
				fmt.Printf("PrevHash: %s\n\n", tx.PreviousHash)
			}

		case 3:
			valid, msg := ledger.ValidateLedger()
			if valid {
				fmt.Println("Ledger VALID:", msg)
			} else {
				fmt.Println("Ledger INVALID:", msg)
			}

		case 4:
			if len(ledger.Transactions) == 0 {
				fmt.Println("No transactions to corrupt.")
				continue
			}
			
			fmt.Print("Enter corruption command (e.g., 'corrupt 1 data' or 'corrupt 0 hash'): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			args := strings.Fields(input)
			if len(args) != 3 || args[0] != "corrupt" {
				fmt.Println("Invalid corrupt command format.")
				continue
			}

			index, err := strconv.Atoi(args[1])
			if err != nil || index < 0 || index >= len(ledger.Transactions) {
				fmt.Println("Invalid transaction index.")
				continue
			}

			field := strings.ToLower(args[2])
			switch field {
			case "data":
				ledger.Transactions[index].Data = "CORRUPTED_DATA"
				fmt.Printf("Tx %d data corrupted.\n", index)
			case "hash":
				ledger.Transactions[index].Hash = "CORRUPTED_HASH"
				fmt.Printf("Tx %d hash corrupted.\n", index)
			case "prevhash":
				ledger.Transactions[index].PreviousHash = "CORRUPTED_PREVHASH"
				fmt.Printf("Tx %d previous hash corrupted.\n", index)
			default:
				fmt.Println("Invalid field. Use 'data', 'hash', or 'prevhash'.")
				continue
			}

		case 5:
			valid, msg := ReplayLedger(ledger.Transactions)
			if valid {
				fmt.Println("Replay Status: VALID -", msg)
			} else {
				fmt.Println("Replay Status: FAILED -", msg)
			}

		case 6:
			hash := ledger.GetLedgerHash()
			if hash == "" {
				fmt.Println("Ledger is empty, no hash.")
			} else {
				fmt.Printf("Final Ledger Hash: %s\n", hash)
			}

		case 7:
			fmt.Println("\n--- Determinism Proof ---")
			l1 := Ledger{}
			l1.AddTransaction("A")
			l1.AddTransaction("B")
			l1.AddTransaction("C")
			hash1 := l1.GetLedgerHash()
			fmt.Printf("Run 1 (A -> B -> C) Final Hash: %s\n", hash1)

			l2 := Ledger{}
			l2.AddTransaction("A")
			l2.AddTransaction("B")
			l2.AddTransaction("C")
			hash2 := l2.GetLedgerHash()
			fmt.Printf("Run 2 (A -> B -> C) Final Hash: %s\n", hash2)

			if hash1 == hash2 {
				fmt.Println("Determinism Proof: SUCCESS (Hashes are identical)")
			} else {
				fmt.Println("Determinism Proof: FAILED")
			}

		case 8:
			fmt.Println("\n--- Replay Proof ---")
			l := Ledger{}
			l.AddTransaction("Valid 1")
			l.AddTransaction("Valid 2")
			valid, msg := ReplayLedger(l.Transactions)
			fmt.Printf("Replay valid ledger -> %v (%s)\n", valid, msg)

			l.Transactions[1].Data = "Corrupted Data"
			valid, msg = ReplayLedger(l.Transactions)
			fmt.Printf("Replay corrupted ledger -> %v (%s)\n", valid, msg)

		case 9:
			return

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}