// internal/blockchain/ledger.go
package blockchain

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Ledger struct {
	Blocks []Block
	DB     *sqlx.DB
}

func NewLedger(db *sqlx.DB) *Ledger {
	genesis := NewBlock([]Transaction{}, "")
	l := &Ledger{
		Blocks: []Block{genesis},
		DB:     db,
	}
	l.saveBlock(genesis)
	return l
}

func (l *Ledger) AddBlock(transactions []Transaction) {
	prevHash := l.Blocks[len(l.Blocks)-1].Hash
	newBlock := NewBlock(transactions, prevHash)
	l.Blocks = append(l.Blocks, newBlock)
	l.saveBlock(newBlock)
}

func (l *Ledger) saveBlock(b Block) {
	tx := l.DB.MustBegin()
	_, err := tx.Exec(`INSERT INTO blocks (timestamp, hash, prev_hash, nonce) VALUES (?, ?, ?, ?)`,
		b.Timestamp, b.Hash, b.PrevHash, b.Nonce)
	if err != nil {
		log.Println("block save error:", err)
		return
	}

	for _, t := range b.Transactions {
		_, err := tx.Exec(`INSERT INTO transactions (block_hash, sender, recipient, amount) VALUES (?, ?, ?, ?)`,
			b.Hash, t.Sender, t.Recipient, t.Amount)
		if err != nil {
			log.Println("tx save error:", err)
		}
	}
	tx.Commit()
}
