package rest

import (
	"crypto/ed25519"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/0xygengg/air/apps/backend/internal/blockchain"
	"github.com/0xygengg/air/apps/backend/internal/core"
	"github.com/0xygengg/air/apps/backend/internal/p2p"
)

func handleGetChain(w http.ResponseWriter, r *http.Request) {
	chain := blockchain.GetChain()
	json.NewEncoder(w).Encode(chain)
}

func handlePostTransaction(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var tx blockchain.Transaction
	if err := json.Unmarshal(body, &tx); err != nil {
		http.Error(w, "Invalid tx format", 400)
		return
	}
	// TODO: Validate tx
	blockchain.AddTransactionToMempool(tx) // <-- You can implement a mempool later
	json.NewEncoder(w).Encode(map[string]string{"status": "tx received"})
}

func handlePostStake(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var req struct {
		Address string `json:"address"`
		Amount  int    `json:"amount"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid stake format", 400)
		return
	}
	blockchain.Stake(req.Address, req.Amount)
	json.NewEncoder(w).Encode(map[string]string{"status": "stake registered"})
}

func handlePostProposeBlock(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var req struct {
		PrivateKey string `json:"privkey_hex"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid key", 400)
		return
	}

	priv := core.DecodeHexPrivateKey(req.PrivateKey)
	pub := priv.Public().(ed25519.PublicKey)

	txs := blockchain.GetMempool()
	block := blockchain.ProposeBlock(txs, priv, pub)
	blockchain.AddBlock(block)
	blockchain.ClearMempool()

	blockBytes, _ := json.Marshal(block)
	msg := []byte("BLOCK:" + string(blockBytes))
	p2p.Broadcast(msg)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "block proposed",
		"block":  block,
	})
}
