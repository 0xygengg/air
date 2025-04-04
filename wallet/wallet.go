package wallet

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/scrypt"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
}

const keyFile = "wallet/keys/private.enc"
const saltSize = 16

func NewWallet() *Wallet {
	os.MkdirAll(filepath.Dir(keyFile), os.ModePerm)

	password := promptPassword("🔐 Enter wallet password: ")

	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		return createAndSaveWallet(password)
	}
	return loadWallet(password)
}

func createAndSaveWallet(password string) *Wallet {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	privBytes, _ := x509.MarshalECPrivateKey(priv)

	encrypted, err := encryptWithPassword(privBytes, password)
	if err != nil {
		panic("Failed to encrypt wallet: " + err.Error())
	}

	ioutil.WriteFile(keyFile, encrypted, 0600)

	return &Wallet{
		PrivateKey: priv,
		PublicKey:  priv.PublicKey,
	}
}

func loadWallet(password string) *Wallet {
	encData, err := ioutil.ReadFile(keyFile)
	if err != nil {
		panic("Failed to read wallet file: " + err.Error())
	}

	decrypted, err := decryptWithPassword(encData, password)
	if err != nil {
		panic("Failed to decrypt wallet: " + err.Error())
	}

	priv, err := x509.ParseECPrivateKey(decrypted)
	if err != nil {
		panic("Failed to parse EC private key: " + err.Error())
	}

	return &Wallet{
		PrivateKey: priv,
		PublicKey:  priv.PublicKey,
	}
}

func PublicKeyToAddress(pub ecdsa.PublicKey) string {
	pubKeyBytes := append(pub.X.Bytes(), pub.Y.Bytes()...)
	hash := sha256Hash(pubKeyBytes)
	return hex.EncodeToString(hash)
}

func sha256Hash(data []byte) []byte {
	h := make([]byte, 32)
	copy(h, data)
	return h[:32]
}

func promptPassword(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	password, _ := reader.ReadString('\n')
	return strings.TrimSpace(password)
}

// ---- Encryption Helpers ----

func encryptWithPassword(data []byte, password string) ([]byte, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	key, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nil, nonce, data, nil)

	return append(salt, append(nonce, ciphertext...)...), nil
}

func decryptWithPassword(data []byte, password string) ([]byte, error) {
	salt := data[:saltSize]
	key, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	nonce := data[saltSize : saltSize+nonceSize]
	ciphertext := data[saltSize+nonceSize:]

	return aesGCM.Open(nil, nonce, ciphertext, nil)
}
