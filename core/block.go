package core

import (
	"time"
)

type Block struct {
	Timestamp  int64
	ParentHash []byte
	Hash       []byte
	Data       []byte
	Nonce      int64
}

func NewGenesisBlock() *Block {
	return NewBlock([]byte("Genesis Block"), []byte{})
}

func NewBlock(data []byte, parentHash []byte) *Block {
	block := &Block{
		Timestamp:  time.Now().Unix(),
		ParentHash: parentHash,
		Hash:       []byte{},
		Data:       data,
		Nonce:      0,
	}
	block.Nonce, block.Hash = NewProofOfWork(block).Run()
	return block
}
