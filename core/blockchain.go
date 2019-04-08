package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const (
	DbFile = "blockchain.db"
	BlocksBucket = "blocks"
)

type BlockChain struct {
	db            *bolt.DB
	lastBlockHash []byte // Hash of the last block in the blockchain
}

func NewBlockchain() *BlockChain {
	var lastBlockHash []byte
	db, err := bolt.Open(DbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BlocksBucket))
		if bucket == nil { // blockchain does not exist in the DB
			fmt.Println("No exiting blockchian found. Creating a new one ...")
			genesis := NewGenesisBlock()

			bucket, err := tx.CreateBucket([]byte(BlocksBucket))
			if err !=  nil {
				log.Panic(err)
			}

			// Store the genesis block
			err = bucket.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			// Store the last bock hash
			err = bucket.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			lastBlockHash = genesis.Hash
		} else { // blockchain exist. Read the last bock
			lastBlockHash = bucket.Get([]byte("l"))
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := BlockChain{db, lastBlockHash}
	return &bc
}

func (bc *BlockChain) AddBlock(data string) {
	var parentBlockHash []byte

	// Read db to retrieve the last block hash
	err := bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BlocksBucket))
		parentBlockHash = bucket.Get([]byte("l"))
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock([]byte(data), parentBlockHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BlocksBucket))

		// Write the new block to DB
		err := bucket.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		// Write the last block hash to DB
		err = bucket.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.lastBlockHash = newBlock.Hash

		return nil
	})
}

func (bc *BlockChain) GetDb() *bolt.DB {
	return bc.db
}

// BlockchainIterator is used to iterate over blockchain blocks
type BlockChainIterator struct {
	blockchain *BlockChain
	currentHash []byte
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc, bc.lastBlockHash}
}

// Next returns the parent block starting from the last block of the blockchain
func (i *BlockChainIterator) Next() *Block {
	var block *Block

	err := i.blockchain.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BlocksBucket))
		encodedBlock := bucket.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.ParentHash

	return block
}
