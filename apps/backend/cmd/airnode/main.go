// apps/backend/cmd/airnode/main.go
package main

import (
	"log"

	"github.com/0xygengg/air/backend/internal/api"
)

func main() {
	server := api.NewServer()
	log.Println("🌐 Air node is running on http://localhost:8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
