package block

const Difficulty = 4

func MineBlock(block Block) Block {
	block.Nonce = 0

	for {
		hash := CalculateHash(block)

		if isValidHash(hash) {
			block.Hash = hash
			return block
		}

		block.Nonce++
	}
}

func isValidHash(hash string) bool {
	prefix := ""

	for i := 0; i < Difficulty; i++ {
		prefix += "0"
	}

	return hash[:Difficulty] == prefix
}
