package config

import "os"

func Load() map[string]string {
	return map[string]string{
		"port":      getEnv("P2P_PORT", "8080"),
		"peers":     getEnv("P2P_PEERS", ""),
		"port_http": getEnv("HTTP_PORT", "8090"),
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
