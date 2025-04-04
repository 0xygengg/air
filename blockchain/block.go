package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

func (b *Block) GenerateHash() {
	input := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash
	hash := sha256.Sum256([]byte(input))
	b.Hash = hex.EncodeToString(hash[:])
}

func CreateGenesisBlock() Block {
	block := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
	}
	block.GenerateHash()
	return block
}
