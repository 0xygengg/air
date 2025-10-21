package p2p

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/0xygengg/air/apps/backend/internal/blockchain"
)

// StartPeerDiscovery launches the TCP listener and connects to initial peers
func StartPeerDiscovery(cfg map[string]string) {
	port := cfg["port"]
	peerList := cfg["peers"]

	go listen(port)

	for _, peer := range strings.Split(peerList, ",") {
		if peer != "" {
			go connect(peer)
		}
	}
}

// listen starts a TCP server on the configured port
func listen(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	fmt.Println("🌐 Listening on port", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

// connect dials a remote peer and adds it to the peer list
func connect(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("❌ Failed to connect to %s: %v\n", addr, err)
		return
	}
	AddPeer(addr, conn)
	go handleConn(conn)
}

// handleConn reads incoming messages line by line and handles them
func handleConn(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	AddPeer(addr, conn)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("📩 Received from %s: %s\n", addr, line)

		// Handle block broadcast
		if strings.HasPrefix(line, "BLOCK:") {
			jsonData := strings.TrimPrefix(line, "BLOCK:")
			var block blockchain.Block
			err := json.Unmarshal([]byte(jsonData), &block)
			if err != nil {
				fmt.Println("⛔ Failed to unmarshal block:", err)
				continue
			}
			if block.ValidateHash() {
				blockchain.AddBlock(block)
				fmt.Println("✅ Block accepted:", block.Index)
			} else {
				fmt.Println("🚫 Invalid block rejected")
			}
		}
	}

	RemovePeer(addr)
}
