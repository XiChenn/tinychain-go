package core

import (
	"bytes"
	"encoding/gob"
	"log"
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

// Convert a block to a bytes array
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// Deserialize a bytes array to a block
func DeserializeBlock(b []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}