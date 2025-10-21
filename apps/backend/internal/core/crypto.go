package core

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/hex"
)

func Hash(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func GenerateKeypair() (ed25519.PublicKey, ed25519.PrivateKey) {
	pub, priv, _ := ed25519.GenerateKey(nil)
	return pub, priv
}

func Sign(data []byte, priv ed25519.PrivateKey) []byte {
	return ed25519.Sign(priv, data)
}

func Verify(data, sig []byte, pub ed25519.PublicKey) bool {
	return ed25519.Verify(pub, data, sig)
}

func DecodeHexPrivateKey(hexStr string) ed25519.PrivateKey {
	raw, _ := hex.DecodeString(hexStr)
	return ed25519.PrivateKey(raw)
}
