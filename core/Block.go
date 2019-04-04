package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Timestamp int64
	PrevBlockHash string
	Hash string
	Data string
}

func NewBlock(data, prevHash string) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevHash,
		Data:          data,
	}
	block.SetHash()
	return block
}

func (block *Block) SetHash() {
	headers := string(block.Timestamp) + block.PrevBlockHash + block.Data
	hashInBytes := sha256.Sum256([]byte(headers))
	block.Hash = hex.EncodeToString(hashInBytes[:])
}

func NewGenesisBlock() *Block {
	 return NewBlock("Genesis Block", "")
}