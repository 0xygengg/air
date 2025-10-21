package main

import (
	"github.com/0xygengg/air/apps/backend/internal/api/rest"
	"github.com/0xygengg/air/apps/backend/internal/blockchain"
	"github.com/0xygengg/air/apps/backend/internal/config"
	"github.com/0xygengg/air/apps/backend/internal/p2p"
)

func main() {
	cfg := config.Load()

	go p2p.StartPeerDiscovery(cfg)

	blockchain.InitChain()

	rest.StartServer(cfg)
}
