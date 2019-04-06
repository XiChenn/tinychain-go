package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
)

const (
	MaxNonce  = math.MaxInt64
	TargetBit = 20 // number of bits with zeros
)

type ProofOfWork struct {
	block *Block
	target *big.Int // the mined hash should < target
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - TargetBit))
	return &ProofOfWork{block, target}
}

// Used by PoW:
func (pow *ProofOfWork) generateHeader(nonce int64) string {
	headers := string(pow.block.Timestamp) + pow.block.ParentHash + pow.block.Data + string(nonce)
	return headers
}

// Find the hash
func (pow *ProofOfWork) Run() (int64, string) {
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	var hashInt big.Int
	var hashInBytes [32]byte
	var nonce int64 = 0

	for nonce < MaxNonce {
		headers := pow.generateHeader(nonce)
		hashInBytes = sha256.Sum256([]byte(headers))
		hashInt.SetBytes(hashInBytes[:])
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r%x", hashInBytes)
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hex.EncodeToString(hashInBytes[:])
}

func (pow *ProofOfWork) Validate() bool {
	headers := pow.generateHeader(pow.block.Nonce)
	hashInBytes := sha256.Sum256([]byte(headers))

	var hashInt big.Int
	hashInt.SetBytes(hashInBytes[:])

	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
