# Deterministic Ledger Implementation (Go)

## Overview

This project implements a deterministic, replayable, self-verifying ledger system using Go.
The system records transactions and links them together using SHA256 hash chaining to preserve data integrity.

The objective of this project is to demonstrate the core principles behind blockchain systems:
- Deterministic hashing
- Hash chaining
- Tamper detection
- Ledger validation and replay

---

## Determinism Explanation

In a fully deterministic system, the same sequence of inputs must always produce the exact same sequence of outputs, hashes, and final states. 
To achieve this:
- **No Randomness**: We removed all sources of entropy (such as `time.Now()` or external timestamping).
- **Pure Functions**: The `NewTransaction` function constructs transactions without relying on any external application state.
- **Strict Hash Formula**: Every hash is computed solely based on `ID + Data + PreviousHash`.

By making the ledger completely deterministic, we can guarantee that if two programs process the exact same sequence of transactions (e.g., `A → B → C`), their final state hashes will always be perfectly identical.

---

## Replay Explanation

The system features a **Replay Engine** (`ReplayLedger`) which allows the entire transaction history to be recalculated from scratch.

During a replay sequence, the engine iterates through the provided chain of transactions from index `0`. For each transaction, it:
1. Recalculates the transaction hash dynamically based strictly on its `ID`, `Data`, and `PreviousHash`.
2. Compares the derived hash against the transaction's stored `Hash`.
3. Verifies that the transaction's `PreviousHash` identically matches the `Hash` of the immediately preceding transaction (or `"GENESIS"` if it is the first node).

If the engine can successfully replay the entire ledger without discrepancies, the ledger's integrity is guaranteed.

---

## Failure Scenarios

The validation and replay engine will intercept and clearly detail exact causes of tampering:

### Hash Mismatch
If a transaction's `Data`, `ID`, or `PreviousHash` is artificially modified after creation, the recalculated SHA256 hash will not reflect its stored `.Hash` value. The system will throw: `"Hash mismatch at Tx X"`.

### Broken Chain
If a transaction's `PreviousHash` link is altered to point to a different hash entirely, or if a previous transaction's payload was modified (causing its hash to change), the contiguous continuity is severed. The system will throw: `"Broken chain at Tx X"`.

---

## CLI Usage

The program provides an interactive, terminal-based CLI for manipulating and verifying the immutable structures.

**Available Commands:**
1. **Add Transaction**: Provide string data to securely append a new transaction onto the ledger.
2. **View Ledger**: Print the sequential history of transactions alongside their precise hex-encoded Hashes.
3. **Validate Ledger**: Run the self-verification integrity check on the live chain.
4. **Corrupt Ledger**: Intentionally attack individual transactions to test the tamper detection threshold. Format `corrupt <index> <field>`. Valid fields are `data`, `hash`, or `prevhash`. Example: `corrupt 1 data`.
5. **Replay Ledger**: Extract the live transactions and process them through the scratch-built Replay Engine.
6. **Get Ledger Hash**: Snapshot and print the terminal transaction's hash (the cumulative identifier for the ledger).
7. **Determinism Proof**: Automatically spin up parallel ledgers, feed them identical requests (`A → B → C`), and prove convergence on an identical terminal hash.
8. **Replay Proof**: Automatically submit valid entries, verify positive structural replay, actively attack the chain state with targeted corruption, and prove the successful validation failure mechanism.
9. **Exit**: Terminate session.

---

## Technologies Used

- **Go (Golang)**
- **SHA256** cryptographic hashing
- **Command Line Interface (CLI)**

> No external blockchain libraries were used.
