package chain

import (
	"fmt"
	"sync"

	"github.com/Hanningtone03/tiny-blockchain/internal/block"
)

type Blockchain struct {
	mu         sync.Mutex
	Blocks     []*block.Block
	Difficulty int
	Pending    []block.Transaction
}

func NewBlockchain(difficulty int) *Blockchain {
	genesis := block.GenesisBlock(difficulty)
	return &Blockchain{
		Blocks:     []*block.Block{genesis},
		Difficulty: difficulty,
		Pending:    []block.Transaction{},
	}
}

func (bc *Blockchain) AddTransaction(tx block.Transaction) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.Pending = append(bc.Pending, tx)
}

func (bc *Blockchain) MineBlock() *block.Block {
	bc.mu.Lock()
	pending := bc.Pending
	bc.Pending = []block.Transaction{}
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	bc.mu.Unlock()

	if len(pending) == 0 {
		pending = []block.Transaction{{From: "system", To: "system", Amount: 0}}
	}

	newBlock := block.NewBlock(prevBlock.Index+1, pending, prevBlock.Hash, bc.Difficulty)
	newBlock.Mine()

	bc.mu.Lock()
	bc.Blocks = append(bc.Blocks, newBlock)
	bc.mu.Unlock()

	return newBlock
}

func (bc *Blockchain) LatestBlock() *block.Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) Length() int {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	return len(bc.Blocks)
}

func (bc *Blockchain) GetBlocks() []*block.Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	result := make([]*block.Block, len(bc.Blocks))
	copy(result, bc.Blocks)
	return result
}

func (bc *Blockchain) ReplaceChain(newBlocks []*block.Block) error {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	if len(newBlocks) <= len(bc.Blocks) {
		return fmt.Errorf("received chain is not longer than current chain")
	}

	bc.Blocks = newBlocks
	return nil
}