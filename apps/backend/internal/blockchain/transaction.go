package blockchain

import (
	"encoding/json"

	"github.com/0xygengg/air/apps/backend/internal/core"
)

type Transaction struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
	Nonce     int     `json:"nonce"`
	Signature string  `json:"signature"`
}

func (tx *Transaction) Hash() string {
	data, _ := json.Marshal(tx)
	return core.Hash(data)
}

func (tx *Transaction) Sign(privateKey []byte) {
	data, _ := json.Marshal(tx)
	sig := core.Sign(data, privateKey)
	tx.Signature = core.Hash(sig)
}

func (tx *Transaction) VerifySignature(pubKey []byte) bool {
	data, _ := json.Marshal(tx)
	return core.Verify(data, []byte(tx.Signature), pubKey)
}
