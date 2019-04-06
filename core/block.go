package core

import (
	"time"
)

type Block struct {
	Timestamp  int64
	ParentHash string
	Hash       string
	Data       string
	Nonce      int64
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}

func NewBlock(data, parentHash string) *Block {
	block := &Block{
		Timestamp:  time.Now().Unix(),
		ParentHash: parentHash,
		Hash:       "",
		Data:       data,
		Nonce:      0,
	}
	block.Nonce, block.Hash = NewProofOfWork(block).Run()
	return block
}
