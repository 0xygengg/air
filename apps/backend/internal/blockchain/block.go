package blockchain

import (
	"encoding/json"
	"time"

	"github.com/0xygengg/air/apps/backend/internal/core"
)

type Block struct {
	Index        int           `json:"index"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash     string        `json:"prev_hash"`
	MerkleRoot   string        `json:"merkle_root"`
	Validator    string        `json:"validator"`
	Signature    string        `json:"signature"`
	Hash         string        `json:"hash"`
}

func (b *Block) CalculateHash() string {
	blockData := map[string]interface{}{
		"index":       b.Index,
		"timestamp":   b.Timestamp,
		"prev_hash":   b.PrevHash,
		"merkle_root": b.MerkleRoot,
		"validator":   b.Validator,
	}
	raw, _ := json.Marshal(blockData)
	return core.Hash(raw)
}

func (b *Block) SignBlock(privKey []byte) {
	data, _ := json.Marshal(b)
	signature := core.Sign(data, privKey)
	b.Signature = core.Hash(signature)
}

func (b *Block) ValidateHash() bool {
	return b.Hash == b.CalculateHash()
}

func ProposeBlock(transactions []Transaction, privKey []byte, pubKey []byte) Block {
	latest := GetLatestBlock()
	validator := core.Hash(pubKey)

	block := Block{
		Index:        latest.Index + 1,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     latest.Hash,
		MerkleRoot:   BuildMerkleRoot(transactions),
		Validator:    validator,
	}

	block.Hash = block.CalculateHash()
	block.SignBlock(privKey)

	return block
}

func ParseBlockFromJSON(jsonStr string) Block {
	var block Block
	_ = json.Unmarshal([]byte(jsonStr), &block)
	return block
}
