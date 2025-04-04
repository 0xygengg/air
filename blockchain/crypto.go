package blockchain

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
)

func HashTransaction(tx Transaction) []byte {
	txCopy := tx
	txCopy.Signature = Signature{}
	bytes, _ := json.Marshal(txCopy)
	hash := sha256.Sum256(bytes)
	return hash[:]
}

func SignTransaction(tx *Transaction, priv *ecdsa.PrivateKey) {
	hash := HashTransaction(*tx)
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash)
	if err != nil {
		panic(err)
	}
	tx.Signature = Signature{R: r, S: s}
}

func VerifyTransaction(tx Transaction, pub ecdsa.PublicKey) bool {
	hash := HashTransaction(tx)
	return ecdsa.Verify(&pub, hash, tx.Signature.R, tx.Signature.S)
}
