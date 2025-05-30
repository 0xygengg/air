// internal/blockchain/transaction.go
package blockchain

import "fmt"

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
}

func serializeTransactions(txns []Transaction) string {
	result := ""
	for _, tx := range txns {
		result += tx.Sender + tx.Recipient + fmt.Sprintf("%f", tx.Amount)
	}
	return result
}
