package chain

import (
	"strings"

	"github.com/Hanningtone03/tiny-blockchain/internal/block"
)

func IsBlockValid(newBlock, prevBlock *block.Block) bool {
	if newBlock.Index != prevBlock.Index+1 {
		return false
	}
	if newBlock.PrevHash != prevBlock.Hash {
		return false
	}
	if newBlock.ComputeHash() != newBlock.Hash {
		return false
	}

	target := strings.Repeat("0", newBlock.Difficulty)
	if len(newBlock.Hash) < newBlock.Difficulty || newBlock.Hash[:newBlock.Difficulty] != target {
		return false
	}

	expectedRoot := block.ComputeMerkleRoot(newBlock.Transactions)
	if newBlock.MerkleRoot != expectedRoot {
		return false
	}

	return true
}

func IsChainValid(blocks []*block.Block) bool {
	if len(blocks) == 0 {
		return false
	}

	for i := 1; i < len(blocks); i++ {
		if !IsBlockValid(blocks[i], blocks[i-1]) {
			return false
		}
	}

	return true
}