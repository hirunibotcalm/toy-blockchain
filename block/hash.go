package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func CalculateHash(block Block) string {
	data := fmt.Sprintf("%d%d%v%s%d",
		block.Index,
		block.Timestamp,
		block.Transactions,
		block.PreviousHash,
		block.Nonce,
	)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
