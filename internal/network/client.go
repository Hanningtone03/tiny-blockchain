package network

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Hanningtone03/tiny-blockchain/internal/block"
)

var currentPort string

func SetCurrentPort(port string) {
	currentPort = port
}

func BroadcastBlock(peers []string, newBlock *block.Block) {
	for _, peer := range peers {
		go sendChainToPeer(peer)
	}
	_ = newBlock
}

func sendChainToPeer(peerURL string) {
	resp, err := http.Get("http://localhost:" + currentPort + "/chain")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var blocks []*block.Block
	json.NewDecoder(resp.Body).Decode(&blocks)

	data, _ := json.Marshal(blocks)
	client := &http.Client{Timeout: 2 * time.Second}
	client.Post("http://"+peerURL+"/receiveChain", "application/json", bytes.NewReader(data))
}