package core

import (
	"bytes"
	"crypto/sha256"
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
func (pow *ProofOfWork) generateHeader(nonce int64) []byte {
	headers := bytes.Join(
		[][]byte{
			Int64ToBytes(pow.block.Timestamp),
			pow.block.ParentHash,
			pow.block.Data,
			Int64ToBytes(nonce),
		}, []byte{})
	return headers
}

// Find the hash
func (pow *ProofOfWork) Run() (int64, []byte) {
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	var hashInt big.Int
	var hashInBytes [32]byte
	var nonce int64 = 0

	for nonce < MaxNonce {
		headers := pow.generateHeader(nonce)
		hashInBytes = sha256.Sum256(headers)
		hashInt.SetBytes(hashInBytes[:])
		fmt.Printf("\r%x", hashInBytes)

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hashInBytes[:]
}

func (pow *ProofOfWork) Validate() bool {
	headers := pow.generateHeader(pow.block.Nonce)
	hashInBytes := sha256.Sum256(headers)

	var hashInt big.Int
	hashInt.SetBytes(hashInBytes[:])

	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
