# Phase 3: Validation, Replay, and CLI Integration

## 1. Entry Point
Our system is driven by `main.go`, which implements an interactive Command Line Interface. It uses `bufio.NewReader` for clean input handling and provides a menu loop to let users add transactions, run the validation tool, simulate network replays, and intentionally corrupt the ledger to test its security.

## 2. Execution Flow
The standard lifecycle of our ledger is as follows:
- **Add**: The user inputs data, and `ledger.AddTransaction(data)` securely links it to the previous hash.
- **Validate**: The system recalculates each node to mathematically prove the chain is contiguous.
- **Replay**: `ReplayLedger()` processes the history from scratch, computationally rebuilding the hashes to verify structure parity.
- **Hash**: The terminal node acts as the `Final Ledger Hash`, representing the accumulated state of all previous entries.

## 3. Replay Proof
When running a replay on a clean ledger, the engine recalculates all hashes and confirms the chain matches perfectly:
```text
--- Replay Proof ---
[LOG] Created transaction ID: 0, Hash: 39c1464f7e7b73f88ebc26724f07c1a43b9baa0a823de7b5dad1186eefce9450        
[LOG] Created transaction ID: 1, Hash: da7c06fe8d0463cb96f9b024b0fae9a907b2744ec23feecbd0aca6e490936eea        
[LOG] Replay started...
[LOG] Replay ended.
Replay valid ledger -> true (Valid replay)
```
If we corrupt data on a previously established block and force a replay, the recalculation catches the spoof:
```text
[LOG] Replay started...
[LOG] Replay ended.
Replay corrupted ledger -> false (Hash mismatch at Tx 1)
```

## 4. Failure Trace
Our validation logic is designed to pinpoint exactly where and how a chain failed.

- **Hash Mismatch**: If an attacker tries to change the `.Data` of a transaction but leaves everything else alone, `calculateHash(current)` will not equal `current.Hash`. The system catches this state drift and explicitly flags `"Hash mismatch at Tx X"`.
- **Broken Chain**: If an attacker tries to recalculate the spoofed transaction's hash to hide the mismatch, the *next* transaction in the list will still be holding the *old* `.PreviousHash`. When validating, checking `current.PreviousHash != previous.Hash` exposes the discontinuity, throwing `"Broken chain at Tx X"`.

## 5. Sample Outputs

**Adding Transactions:**
```text
Choice: 1
Enter transaction data: A
[LOG] Created transaction ID: 0, Hash: d0c306afe0d6aa7bd475d95fc51c8e09777184a0d95c81db4e8a1ecd918eec92        
Transaction Added

Choice: 1
Enter transaction data: B
[LOG] Created transaction ID: 1, Hash: ae5a0d359de8aeb57db581db0877274b50acf8ee083b2943669fad0dc26b5b5e        
Transaction Added
```

**Validation Case:**
```text
Choice: 3
[LOG] Validation started...
[LOG] Validation ended.
Ledger VALID: Valid
```
