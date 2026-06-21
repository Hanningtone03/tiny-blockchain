package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func hashTransaction(tx Transaction) string {
	data, _ := json.Marshal(tx)
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}

func hashPair(a, b string) string {
	h := sha256.Sum256([]byte(a + b))
	return hex.EncodeToString(h[:])
}

func ComputeMerkleRoot(transactions []Transaction) string {
	if len(transactions) == 0 {
		h := sha256.Sum256([]byte(""))
		return hex.EncodeToString(h[:])
	}

	var layer []string
	for _, tx := range transactions {
		layer = append(layer, hashTransaction(tx))
	}

	for len(layer) > 1 {
		var nextLayer []string
		for i := 0; i < len(layer); i += 2 {
			if i+1 < len(layer) {
				nextLayer = append(nextLayer, hashPair(layer[i], layer[i+1]))
			} else {
				nextLayer = append(nextLayer, hashPair(layer[i], layer[i]))
			}
		}
		layer = nextLayer
	}

	return layer[0]
}