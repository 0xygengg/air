package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"air/blockchain"
	"air/wallet"
)

func main() {
	// Load or create a wallet
	userWallet := wallet.NewWallet()
	myAddress := wallet.PublicKeyToAddress(userWallet.PublicKey)
	fmt.Println("🔑 Your wallet address:", myAddress)

	// Seed your wallet with an initial balance
	blockchain.AccountState[myAddress] = 1000.0

	var chain []blockchain.Block

	// Create the genesis block
	genesis := blockchain.CreateGenesisBlock()
	chain = append(chain, genesis)
	fmt.Println("🟢 Genesis block created.")

	// 🔥 Start the JSON API (Fiber)
	go StartAPI(&chain)

	// Optional: keep CLI interactive mode for manual txs
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n📝 New transaction:")
		fmt.Print("To (or type 'exit'): ")
		to, _ := reader.ReadString('\n')
		to = strings.TrimSpace(to)
		if to == "exit" {
			break
		}

		fmt.Print("Amount: ")
		amountStr, _ := reader.ReadString('\n')
		amountStr = strings.TrimSpace(amountStr)
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Println("❌ Invalid amount. Try again.")
			continue
		}

		if blockchain.AccountState[myAddress] < amount {
			fmt.Println("❌ You don't have enough balance.")
			continue
		}

		tx := blockchain.Transaction{
			From:   myAddress,
			To:     to,
			Amount: amount,
		}

		blockchain.SignTransaction(&tx, userWallet.PrivateKey)
		if !blockchain.VerifyTransaction(tx, userWallet.PublicKey) {
			fmt.Println("❌ Signature verification failed.")
			continue
		}

		prevBlock := chain[len(chain)-1]
		newBlock := blockchain.GenerateNextBlock(prevBlock, []blockchain.Transaction{tx})
		chain = append(chain, newBlock)

		fmt.Printf("✅ Block #%d added! Hash: %s\n", newBlock.Index, newBlock.Hash)
	}

	// Print the blockchain
	fmt.Println("\n🔗 Full chain:")
	for _, block := range chain {
		fmt.Printf("\n🔸 Block #%d\nHash: %s\nPrev: %s\n", block.Index, block.Hash, block.PrevHash)
		for _, tx := range block.Transactions {
			fmt.Printf("  → %s sent %.2f to %s\n", tx.From, tx.Amount, tx.To)
		}
	}

	// Show final balances
	fmt.Println("\n🧾 Account Balances:")
	for addr, bal := range blockchain.AccountState {
		fmt.Printf("  %s: %.2f\n", addr, bal)
	}
}
