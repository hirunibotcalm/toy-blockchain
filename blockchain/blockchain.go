package blockchain

import (
	"fmt"
	"time"
	"toy-blockchain/block"
	"toy-blockchain/ledger"
)

const MiningReward = 50

type Blockchain struct {
	Blocks              []block.Block
	PendingTransactions []ledger.Transaction
	Balances            map[string]float64
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

		Balances: make(map[string]float64),
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

	// Update balances after block confirmation

	for _, tx := range transactions {

		if tx.Sender != "SYSTEM" {

			bc.Balances[tx.Sender] -= tx.Amount

		}

		bc.Balances[tx.Recipient] += tx.Amount

	}
}

func (bc *Blockchain) AddTransaction(tx ledger.Transaction) bool {

	// SYSTEM can create coins
	if tx.Sender != "SYSTEM" {

		balance, exists := bc.Balances[tx.Sender]

		if !exists {
			return false
		}

		if balance < tx.Amount {
			return false
		}

	}

	bc.PendingTransactions = append(
		bc.PendingTransactions,
		tx,
	)

	return true
}

func (bc *Blockchain) MineTransactions() {

	if len(bc.PendingTransactions) == 0 {
		fmt.Println("No transactions available")
		return
	}

	bc.AddBlock(bc.PendingTransactions)

	bc.PendingTransactions = []ledger.Transaction{}

	fmt.Println("Transactions mined successfully")
}

func (bc *Blockchain) GenerateReward(miner string) {

	reward := ledger.Transaction{

		Sender: "SYSTEM",

		Recipient: miner,

		Amount: MiningReward,
	}

	bc.AddBlock(
		[]ledger.Transaction{reward},
	)

	fmt.Println("Mining reward generated")
}
