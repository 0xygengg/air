package blockchain

import "sync"

var (
	mempool     []Transaction
	mempoolLock sync.Mutex
)

// AddTransactionToMempool adds a tx to the shared mempool
func AddTransactionToMempool(tx Transaction) {
	mempoolLock.Lock()
	defer mempoolLock.Unlock()
	mempool = append(mempool, tx)
}

// GetMempool returns all pending transactions
func GetMempool() []Transaction {
	mempoolLock.Lock()
	defer mempoolLock.Unlock()
	// return a copy to avoid race conditions
	out := make([]Transaction, len(mempool))
	copy(out, mempool)
	return out
}

// ClearMempool resets the mempool (used after block is mined)
func ClearMempool() {
	mempoolLock.Lock()
	defer mempoolLock.Unlock()
	mempool = []Transaction{}
}
