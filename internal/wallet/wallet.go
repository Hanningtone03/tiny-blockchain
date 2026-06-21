package wallet

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type Wallet struct {
	Address string
	balance float64
}

func NewWallet() *Wallet {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	hash := sha256.Sum256(bytes)
	address := hex.EncodeToString(hash[:])[:24]

	return &Wallet{
		Address: address,
		balance: 0,
	}
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func (w *Wallet) Credit(amount float64) {
	w.balance += amount
}

func (w *Wallet) Debit(amount float64) bool {
	if amount > w.balance {
		return false
	}
	w.balance -= amount
	return true
}