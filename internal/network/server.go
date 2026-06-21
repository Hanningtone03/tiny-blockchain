package network

import (
	"encoding/json"
	"net/http"

	"github.com/Hanningtone03/tiny-blockchain/internal/block"
	"github.com/Hanningtone03/tiny-blockchain/internal/chain"
)

func StartServer(port string, bc *chain.Blockchain, peers []string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/chain", func(w http.ResponseWriter, r *http.Request) {
		blocks := bc.GetBlocks()
		json.NewEncoder(w).Encode(blocks)
	})

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		latest := bc.LatestBlock()
		json.NewEncoder(w).Encode(map[string]interface{}{
			"length":      bc.Length(),
			"latest_hash": latest.Hash,
			"difficulty":  bc.Difficulty,
		})
	})

	mux.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		var tx block.Transaction
		json.NewDecoder(r.Body).Decode(&tx)
		bc.AddTransaction(tx)
		json.NewEncoder(w).Encode(map[string]string{"status": "added to pending pool"})
	})

	mux.HandleFunc("/mine", func(w http.ResponseWriter, r *http.Request) {
		newBlock := bc.MineBlock()
		BroadcastBlock(peers, newBlock)
		json.NewEncoder(w).Encode(newBlock)
	})

	mux.HandleFunc("/receiveChain", func(w http.ResponseWriter, r *http.Request) {
		var blocks []*block.Block
		json.NewDecoder(r.Body).Decode(&blocks)

		if len(blocks) > bc.Length() {
			err := bc.ReplaceChain(blocks)
			if err == nil {
				json.NewEncoder(w).Encode(map[string]string{"status": "chain replaced"})
				return
			}
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "chain rejected, not longer"})
	})

	go func() {
		err := http.ListenAndServe(":"+port, mux)
		if err != nil {
			panic(err)
		}
	}()
}