package block

import (
	"toy-blockchain/ledger"
)

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []ledger.Transaction
	PreviousHash string
	Hash         string // generate using SHA256
	Nonce        int
}
