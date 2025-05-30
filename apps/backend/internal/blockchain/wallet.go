// internal/blockchain/wallet.go
package blockchain

import (
	"crypto/ed25519"
	"encoding/hex"
)

type Wallet struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
	Address    string
}

func NewWallet() Wallet {
	pub, priv, _ := ed25519.GenerateKey(nil)
	return Wallet{
		PrivateKey: priv,
		PublicKey:  pub,
		Address:    hex.EncodeToString(pub),
	}
}
