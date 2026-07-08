package ledger

func ValidateTransaction(tx Transaction) bool {

	if tx.Sender == "" {
		return false
	}

	if tx.Recipient == "" {
		return false
	}

	if tx.Amount <= 0 {
		return false
	}

	return true
}