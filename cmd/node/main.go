package main

import (
	"fmt"
	"os"

	"github.com/Hanningtone03/tiny-blockchain/internal/chain"
	"github.com/Hanningtone03/tiny-blockchain/internal/network"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/node/main.go <port> [peer1:port,peer2:port]")
		os.Exit(1)
	}

	port := os.Args[1]
	network.SetCurrentPort(port)

	var peers []string
	if len(os.Args) >= 3 {
		peers = parsePeers(os.Args[2])
	}

	difficulty := 4
	bc := chain.NewBlockchain(difficulty)

	network.StartServer(port, bc, peers)

	fmt.Printf("Node started on port %s\n", port)
	fmt.Printf("Genesis block mined: %s\n", bc.LatestBlock().Hash)
	fmt.Printf("Peers: %v\n", peers)
	fmt.Println("\nEndpoints:")
	fmt.Println("  GET  /status         - chain status")
	fmt.Println("  GET  /chain          - full chain")
	fmt.Println("  POST /transaction    - add a transaction")
	fmt.Println("  POST /mine           - mine pending transactions")

	select {}
}

func parsePeers(raw string) []string {
	var peers []string
	current := ""
	for _, c := range raw {
		if c == ',' {
			peers = append(peers, current)
			current = ""
		} else {
			current += string(c)
		}
	}
	if current != "" {
		peers = append(peers, current)
	}
	return peers
}