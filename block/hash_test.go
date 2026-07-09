package block

import "testing"

func TestCalculateHash(t *testing.T) {

	block := Block{
		Index:        1,
		Timestamp:    100,
		PreviousHash: "abc",
		Nonce:        0,
	}

	hash1 := CalculateHash(block)
	hash2 := CalculateHash(block)

	if hash1 != hash2 {
		t.Error("Hash should be consistent")
	}

	if hash1 == "" {
		t.Error("Hash should not be empty")
	}
}
