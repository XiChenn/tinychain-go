package main

import (
	"encoding/json"
	"log"
	"net/http"
	"tinychian-go/core"
)

var blockchain *core.BlockChain

func main() {
	blockchain = core.NewBlockchain()

	http.HandleFunc("/blockchain/get", getHandler)
	http.HandleFunc("/blockchain/write", writeHandler)

	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(blockchain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockchain.AddBlock(blockData)
	getHandler(w, r)
}
