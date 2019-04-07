package core

type BlockChain struct {
	Blocks []*Block
}

func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func (bc *BlockChain) AddBlock(data string) {
	parentBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := NewBlock([]byte(data), parentBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}
