package blockchain

import "github.com/0xygengg/air/apps/backend/internal/core"

func BuildMerkleRoot(transactions []Transaction) string {
	if len(transactions) == 0 {
		return ""
	}

	var hashes []string
	for _, tx := range transactions {
		hashes = append(hashes, tx.Hash())
	}

	for len(hashes) > 1 {
		var next []string
		for i := 0; i < len(hashes); i += 2 {
			if i+1 < len(hashes) {
				combined := hashes[i] + hashes[i+1]
				next = append(next, core.Hash([]byte(combined)))
			}
		}
		hashes = next
	}

	return hashes[0]
}
