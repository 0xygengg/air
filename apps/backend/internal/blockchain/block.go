// internal/blockchain/block.go
package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"` // Always present
	PrevHash     string        `json:"prevHash"`
	Hash         string        `json:"hash"`
	Nonce        int           `json:"nonce"`
}

func NewBlock(transactions []Transaction, prevHash string) Block {
	block := Block{
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevHash,
	}
	block.Mine()
	return block
}

func (b *Block) CalculateHash() string {
	data := string(b.Timestamp) + b.PrevHash + serializeTransactions(b.Transactions) + string(b.Nonce)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (b *Block) Mine() {
	for {
		b.Hash = b.CalculateHash()
		if b.Hash[:4] == "0000" {
			break
		}
		b.Nonce++
	}
}
