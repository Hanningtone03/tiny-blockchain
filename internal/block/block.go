package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"
)

type Transaction struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount float64 `json:"amount"`
}

type Block struct {
	Index        int           `json:"index"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash     string        `json:"prev_hash"`
	Hash         string        `json:"hash"`
	Nonce        int           `json:"nonce"`
	MerkleRoot   string        `json:"merkle_root"`
	Difficulty   int           `json:"difficulty"`
}

func NewBlock(index int, transactions []Transaction, prevHash string, difficulty int) *Block {
	b := &Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevHash,
		Nonce:        0,
		Difficulty:   difficulty,
	}
	b.MerkleRoot = ComputeMerkleRoot(transactions)
	return b
}

func (b *Block) ComputeHash() string {
	record := strconv.Itoa(b.Index) +
		strconv.FormatInt(b.Timestamp, 10) +
		b.PrevHash +
		b.MerkleRoot +
		strconv.Itoa(b.Nonce)

	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func (b *Block) Mine() {
	target := make([]byte, b.Difficulty)
	for i := range target {
		target[i] = '0'
	}
	prefix := string(target)

	for {
		b.Hash = b.ComputeHash()
		if len(b.Hash) >= b.Difficulty && b.Hash[:b.Difficulty] == prefix {
			break
		}
		b.Nonce++
	}
}

func (b *Block) ToJSON() string {
	data, _ := json.MarshalIndent(b, "", "  ")
	return string(data)
}

func GenesisBlock(difficulty int) *Block {
	genesisTx := []Transaction{
		{From: "genesis", To: "network", Amount: 0},
	}
	b := NewBlock(0, genesisTx, "0", difficulty)
	b.Mine()
	return b
}