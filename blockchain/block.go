package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	PrevHash     string
	Hash         string
}

func (b *Block) GenerateHash() {
	txData := ""
	for _, tx := range b.Transactions {
		txData += tx.From + tx.To + fmt.Sprintf("%f", tx.Amount)
	}
	input := strconv.Itoa(b.Index) + b.Timestamp + txData + b.PrevHash
	hash := sha256.Sum256([]byte(input))
	b.Hash = hex.EncodeToString(hash[:])
}

func CreateGenesisBlock() Block {
	genesisTx := Transaction{From: "genesis", To: "network", Amount: 0}
	block := Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transactions: []Transaction{genesisTx},
		PrevHash:     "",
	}
	block.GenerateHash()
	return block
}

func GenerateNextBlock(prev Block, txs []Transaction) Block {
	validTxs := []Transaction{}
	for _, tx := range txs {
		if AccountState[tx.From] < tx.Amount {
			fmt.Printf("❌ Rejected tx: %s has insufficient funds\n", tx.From)
			continue
		}
		AccountState[tx.From] -= tx.Amount
		AccountState[tx.To] += tx.Amount
		validTxs = append(validTxs, tx)
	}

	block := Block{
		Index:        prev.Index + 1,
		Timestamp:    time.Now().String(),
		Transactions: validTxs,
		PrevHash:     prev.Hash,
	}
	block.GenerateHash()
	return block
}
