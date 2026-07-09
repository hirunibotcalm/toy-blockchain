package ledger

import "testing"

func TestValidTransaction(t *testing.T) {

	tx := Transaction{
		Sender:    "Alice",
		Recipient: "Bob",
		Amount:    100,
	}

	if !ValidateTransaction(tx) {
		t.Error("Transaction should be valid")
	}
}

func TestInvalidTransaction(t *testing.T) {

	tx := Transaction{
		Sender:    "",
		Recipient: "Bob",
		Amount:    100,
	}

	if ValidateTransaction(tx) {
		t.Error("Transaction should be invalid")
	}
}
