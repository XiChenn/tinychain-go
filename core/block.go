package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Timestamp  int64
	ParentHash string
	Hash       string
	Data       string
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}

func NewBlock(data, parentHash string) *Block {
	block := &Block{
		Timestamp:  time.Now().Unix(),
		ParentHash: parentHash,
		Data:       data,
	}
	block.SetHash()
	return block
}

func (block *Block) SetHash() {
	headers := string(block.Timestamp) + block.ParentHash + block.Data
	hashInBytes := sha256.Sum256([]byte(headers))
	block.Hash = hex.EncodeToString(hashInBytes[:])
}
