package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/0xygengg/air/backend/internal/blockchain"
	"github.com/0xygengg/air/backend/internal/storage"
)

type APIServer struct {
	App    *fiber.App
	Ledger *blockchain.Ledger
}

func NewServer() *APIServer {
	app := fiber.New()

	// Enable CORS so frontend can access this backend
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	// Initialize SQLite DB and Ledger
	db := storage.InitDB("air.db")
	ledger := blockchain.NewLedger(db)

	server := &APIServer{
		App:    app,
		Ledger: ledger,
	}

	server.setupRoutes()
	server.registerNFTRoutes()

	return server
}

func (s *APIServer) setupRoutes() {
	s.App.Get("/blocks", func(c *fiber.Ctx) error {
		return c.JSON(s.Ledger.Blocks)
	})

	s.App.Post("/transaction", func(c *fiber.Ctx) error {
		var tx blockchain.Transaction
		if err := c.BodyParser(&tx); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid transaction payload",
			})
		}
		s.Ledger.AddBlock([]blockchain.Transaction{tx})
		return c.JSON(fiber.Map{
			"message": "transaction added",
			"tx":      tx,
		})
	})

	s.App.Get("/wallet/create", func(c *fiber.Ctx) error {
		wallet := blockchain.NewWallet()
		return c.JSON(wallet)
	})
}

func (s *APIServer) Run() error {
	return s.App.Listen(":8080")
}
