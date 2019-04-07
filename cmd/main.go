package main

import (
	"fmt"
	"strconv"
	"tinychian-go/core"
)

func main() {
	bc := core.NewBlockchain()
	bc.AddBlock("Send 1 BTC to Jacky")
	bc.AddBlock("Send 1 EOS to Jack")

	for index, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", index)
		fmt.Printf("Parent Hash: %x\n", block.ParentHash)
		fmt.Printf("Current Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)

		pow := core.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Println()
	}
}
