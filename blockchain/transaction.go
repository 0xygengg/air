package blockchain

import "math/big"

type Transaction struct {
	From      string
	To        string
	Amount    float64
	Signature Signature
}

type Signature struct {
	R *big.Int
	S *big.Int
}
