package blockchain

import "toy-blockchain/block"

func (bc *Blockchain) IsValid() bool {

	for i := 1; i < len(bc.Blocks); i++ {

		current := bc.Blocks[i]
		previous := bc.Blocks[i-1]

		// 1. Check hash is correct
		if current.Hash != block.CalculateHash(current) {
			return false
		}

		// 2. Check chain link
		if current.PreviousHash != previous.Hash {
			return false
		}

	}

	return true
}
