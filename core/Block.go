package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index int64
	Timestamp int64
	PrevBlockHash string
	Hash string
	Data string
}

func calculateHash(block *Block) string {
	blockData := string(block.Index) + string(block.Timestamp) + block.PrevBlockHash + block.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

func GenerateNewBlock(preBlock *Block, data string) Block {
	newBlock := Block{
		Index:         preBlock.Index + 1,
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: preBlock.Hash,
		Data:          data,
	}
	newBlock.Hash = calculateHash(&newBlock)
	return newBlock
}

func GenerateGenesisBlock() Block {
	 preBlock := Block{
		 Index:         -1,
		 Hash:          "",
	 }
	 return GenerateNewBlock(&preBlock, "Genesis Block")
}