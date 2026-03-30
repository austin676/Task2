# Phase 2: Achieving True Determinism

## The Determinism Issue
In our initial ledger implementation, we used `time.Now()` to record the exact moment a transaction was created. While this seemed useful for record-keeping, it introduced a major flaw: randomness (or statefulness). Because the timestamp was included in the SHA256 hashing formula, running the exact same sequence of transactions twice would produce completely different hashes due to the different execution times. A true deterministic system requires that the same inputs *always* produce the exact same outputs.

## Timestamp Removal
To fix this, we removed the `Timestamp` field entirely from the `Transaction` struct and the `calculateHash()` function. The hashing formula was strictly reduced to `ID + Data + PreviousHash`.

## Introduction of NewTransaction()
We introduced `NewTransaction(id int, data string, prevHash string) Transaction` as a pure function. A pure function means it has no side effects and doesn't rely on any external state (like a system clock). It simply takes the inputs, builds the struct, calculates the hash, and returns it.

## Why Controlled Construction Matters
By isolating transaction creation into a pure function, we guarantee that the rest of the application cannot accidentally inject non-deterministic data. The transaction builder is locked down, ensuring that every transaction is strictly a product of its explicit inputs.

## The Result
Our ledger system is now strictly deterministic. We can reliably prove that if Node A and Node B both process transactions "Alice pays Bob" and "Bob pays Charlie", their final ledger hashes will perfectly align every single time, enabling true consensus without synchronization issues.
