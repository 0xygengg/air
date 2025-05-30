// internal/api/nft.go
package api

import (
	"github.com/gofiber/fiber/v2"
)

type MintRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
	Owner       string `json:"owner"`
}

func (s *APIServer) registerNFTRoutes() {
	s.App.Post("/nft/mint", func(c *fiber.Ctx) error {
		var req MintRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
		}

		// TODO: Add NFT to SQLite (future: IPFS, ERC-721-compatible store)
		return c.JSON(fiber.Map{
			"message": "NFT minted",
			"data":    req,
		})
	})
}
