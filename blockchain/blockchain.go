package blockchain

import (
	"fmt"
	"time"
	"toy-blockchain/block"
	"toy-blockchain/ledger"
)

type Blockchain struct {
	Blocks              []block.Block
	PendingTransactions []ledger.Transaction
}

func NewBlockchain() *Blockchain {
	genesisBlock := block.Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		Transactions: []ledger.Transaction{},
		PreviousHash: "0x0000000000000000",
		Nonce:        0,
	}

	genesisBlock.Hash = block.CalculateHash(genesisBlock)

	return &Blockchain{
		Blocks: []block.Block{genesisBlock},
	}
}

func (bc *Blockchain) AddBlock(transactions []ledger.Transaction) {

	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := block.Block{
		Index:        lastBlock.Index + 1,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PreviousHash: lastBlock.Hash,
		Nonce:        0,
	}

	newBlock = block.MineBlock(newBlock)

	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) AddTransaction(tx ledger.Transaction) {
	bc.PendingTransactions = append(bc.PendingTransactions, tx)
}

func (bc *Blockchain) MinePendingTransactions() {

	if len(bc.PendingTransactions) == 0 {
		fmt.Println("No transactions available")
		return
	}

	bc.AddBlock(bc.PendingTransactions)

	bc.PendingTransactions = []ledger.Transaction{}
}