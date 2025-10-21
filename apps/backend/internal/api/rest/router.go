package rest

import (
	"log"
	"net/http"
)

func StartServer(cfg map[string]string) {
	port := cfg["port_http"]
	if port == "" {
		port = "8090"
	}

	http.HandleFunc("/api/chain", handleGetChain)
	http.HandleFunc("/api/tx", handlePostTransaction)
	http.HandleFunc("/api/stake", handlePostStake)
	http.HandleFunc("/api/propose", handlePostProposeBlock)

	log.Printf("REST API listening on :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Failed to start REST API:", err)
	}
}
