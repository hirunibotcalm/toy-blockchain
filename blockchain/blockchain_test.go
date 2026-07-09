package blockchain

import (
	"testing"

	"toy-blockchain/ledger"
)

func TestBlockchainCreation(t *testing.T) {

	bc := NewBlockchain()

	if len(bc.Blocks) != 1 {
		t.Error("Blockchain should start with genesis block")
	}

}

func TestAddBlock(t *testing.T) {

	bc := NewBlockchain()

	tx := ledger.Transaction{
		Sender:    "Alice",
		Recipient: "Bob",
		Amount:    100,
	}

	bc.AddBlock([]ledger.Transaction{tx})

	if len(bc.Blocks) != 2 {
		t.Error("Block was not added")
	}

	if !bc.IsValid() {
		t.Error("Blockchain should be valid")
	}
}
