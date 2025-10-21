package blockchain

import (
	"sync"
	"time"
)

var (
	blockchain []Block
	lock       sync.Mutex
)

func InitChain() {
	genesis := Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		Transactions: []Transaction{},
		PrevHash:     "",
		MerkleRoot:   "",
		Validator:    "genesis",
	}
	genesis.Hash = genesis.CalculateHash()
	blockchain = append(blockchain, genesis)
}

func GetLatestBlock() Block {
	lock.Lock()
	defer lock.Unlock()
	return blockchain[len(blockchain)-1]
}

func AddBlock(block Block) {
	lock.Lock()
	defer lock.Unlock()
	blockchain = append(blockchain, block)
}

func GetChain() []Block {
	lock.Lock()
	defer lock.Unlock()
	return blockchain
}
