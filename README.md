# Toy Blockchain and Ledger Simulator

A simple blockchain and ledger simulator implemented from scratch using **Go (Golang)**.

This project demonstrates the fundamental concepts behind blockchain technology, including blocks, transactions, cryptographic hashing, proof-of-work mining, blockchain validation, and a command-line interface.

The system allows users to create transactions, mine new blocks, view the blockchain history, and verify the integrity of the chain.

---

# Project Overview

A blockchain is a distributed ledger system where data is stored in a sequence of blocks. Each block contains transaction data, a unique hash, and the hash of the previous block. This creates a secure chain where any modification can be detected.

This project implements a simplified blockchain model for learning and demonstration purposes.

---

# Technologies Used

* **Programming Language:** Go (Golang 1.22+)
* **Cryptographic Algorithm:** SHA-256
* **Mining Algorithm:** Proof of Work
* **Application Type:** Command Line Application (CLI)

---

# Features

## Block Management

* Create blocks containing transactions
* Store block index, timestamp, transactions, previous hash, hash, and nonce
* Link blocks together using previous hashes

## Transaction Management

* Create transactions between users
* Store pending transactions before mining
* Validate transaction details

## Cryptographic Hashing

* Generate SHA-256 hashes for blocks
* Ensure block integrity
* Detect any modification in blockchain data

## Genesis Block

* Create the first block of the blockchain
* Initialize the blockchain with a starting block

## Proof of Work Mining

* Implement a simple mining algorithm
* Change nonce values until a valid hash is found
* Use difficulty level to control mining complexity

## Blockchain Validation

The system verifies:

* Block hash correctness
* Previous hash connections
* Data integrity

## Command Line Interface

Users can interact with the blockchain through a menu:

```
===== TOY BLOCKCHAIN MENU =====

1. Add Transaction
2. Mine Block
3. View Blockchain
4. Check Validity
5. Exit
```

---

# Project Structure

```
toy-blockchain/

│
├── block/
│   ├── block.go
│   ├── hash.go
│   └── mining.go
│
├── blockchain/
│   ├── blockchain.go
│   └── validation.go
│
├── ledger/
│   ├── transaction.go
│   └── validation.go
│
├── cli/
│   └── menu.go
│
├── storage/
│   └── storage.go
│
├── main.go
├── go.mod
└── README.md
```

---

# Architecture Overview

The system follows a modular package-based architecture.

```
                 CLI Layer
                    |
                    |
                    v

             Blockchain Layer
                    |
                    |
                    v

              Block Layer
                    |
                    |
                    v

             Ledger Layer
```

### CLI Package

Responsible for:

* User interaction
* Receiving transactions
* Displaying blockchain information

### Blockchain Package

Responsible for:

* Managing blocks
* Creating genesis block
* Adding new blocks
* Validating blockchain

### Block Package

Responsible for:

* Block structure
* Hash calculation
* Mining process

### Ledger Package

Responsible for:

* Transaction structure
* Transaction validation

---

# How To Run

## 1. Clone Repository

```bash
git clone <repository-url>
```

## 2. Navigate to Project Folder

```bash
cd toy-blockchain
```

## 3. Run Application

```bash
go run .
```

---

# Example Usage

## 1. Add Transaction

Select:

```
1
```

Example:

```
Sender: Alice
Recipient: Bob
Amount: 100
```

Output:

```
Transaction added to pool
```

---

## 2. Mine Block

Select:

```
2
```

The system performs proof-of-work mining.

Example:

```
Block mined successfully!
```

---

## 3. View Blockchain

Select:

```
3
```

Example output:

```
--------------------
Index: 0
Hash: 8e6a160fe5e9...
PrevHash: 0x0000000000000000
Transactions: []

--------------------
Index: 1
Hash: 00005c929ac7...
PrevHash: 8e6a160fe5e9...
Transactions: [{Alice Bob 100}]
```

---

## 4. Validate Blockchain

Select:

```
4
```

Output:

```
Blockchain is VALID
```

---

# Blockchain Concepts Implemented

## 1. Blocks

A block stores:

* Index
* Timestamp
* Transactions
* Previous block hash
* Current block hash
* Nonce

Example:

```
Block
|
|-- Index
|-- Timestamp
|-- Transactions
|-- PreviousHash
|-- Hash
|-- Nonce
```

---

## 2. Hashing

Each block generates a SHA-256 hash based on its contents.

The hash works like a digital fingerprint.

If any data changes:

```
Original Block
       |
       v
Hash: abc123


Modified Block
       |
       v
Hash: xyz789
```

The changed hash reveals tampering.

---

## 3. Previous Hash Linking

Each block stores the hash of the previous block.

Example:

```
Genesis Block

Hash:
AAA111


        |
        v


Block 1

PreviousHash:
AAA111
```

This creates the blockchain connection.

---

## 4. Proof of Work Mining

Mining searches for a valid hash by changing the nonce.

Example:

```
Nonce 1  -> Hash: 8abc123 ❌

Nonce 2  -> Hash: 45def67 ❌

Nonce 5000 -> Hash: 0000abcd ✅
```

The block is accepted after finding a valid hash.

---

## 5. Blockchain Validation

The blockchain checks:

* Whether hashes are correct
* Whether previous hash references match
* Whether data has been changed

If someone modifies a block:

```
Blockchain becomes INVALID
```

---

# Testing

The project can be checked using:

## Format Code

```bash
go fmt ./...
```

## Run Tests

```bash
go test ./...
```

## Build Project

```bash
go build ./...
```

---

# Future Improvements

Possible future enhancements:

* Persistent storage using JSON/database
* Digital signatures for transactions
* Multiple user wallets
* Balance calculation
* Network communication between nodes
* More advanced consensus algorithms

---

# Learning Outcomes

Through this project, the following concepts were practiced:

* Go package organization
* Struct-based design
* Cryptographic hashing
* Blockchain architecture
* Proof-of-work algorithms
* Command-line application development
* Data integrity validation

---
