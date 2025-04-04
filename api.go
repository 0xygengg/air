package main

import (
	"air/blockchain"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func StartAPI(chain *[]blockchain.Block) {
	app := fiber.New()

	// Get account balance
	app.Get("/balance/:addr", func(c *fiber.Ctx) error {
		addr := c.Params("addr")
		bal := blockchain.AccountState[addr]
		return c.JSON(fiber.Map{"address": addr, "balance": bal})
	})

	// Get full blockchain
	app.Get("/chain", func(c *fiber.Ctx) error {
		return c.JSON(chain)
	})

	// Get latest block
	app.Get("/block/latest", func(c *fiber.Ctx) error {
		if len(*chain) == 0 {
			return c.JSON(fiber.Map{"error": "No blocks yet"})
		}
		return c.JSON((*chain)[len(*chain)-1])
	})

	// Submit a transaction
	app.Post("/tx", func(c *fiber.Ctx) error {
		var tx blockchain.Transaction
		if err := json.Unmarshal(c.Body(), &tx); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		// Signature verification (optional — or can require from frontend)
		if blockchain.AccountState[tx.From] < tx.Amount {
			return c.Status(400).JSON(fiber.Map{"error": "Insufficient funds"})
		}

		// For demo: assume signature is trusted
		block := blockchain.GenerateNextBlock((*chain)[len(*chain)-1], []blockchain.Transaction{tx})
		*chain = append(*chain, block)

		return c.JSON(fiber.Map{"message": "Block added", "block": block})
	})

	app.Listen(":8080")
}
