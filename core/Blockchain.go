package core

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	blockchain := Blockchain{
		Blocks: []*Block{NewGenesisBlock()},
	}
	return &blockchain
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}