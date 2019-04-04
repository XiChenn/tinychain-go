package main

import (
	"fmt"
	"tinychian-go/core"
)

func main() {
	bc := core.NewBlockchain()
	bc.AddBlock("Send 1 BTC to Jacky")
	bc.AddBlock("Send 1 EOS to Jack")

	for index, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", index)
		fmt.Printf("Prev.Hash: %s\n", block.PrevBlockHash)
		fmt.Printf("Curr.Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Println()
	}
}
