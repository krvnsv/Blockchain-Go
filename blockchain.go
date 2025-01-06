package main

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	blocks []*Block // A slice of pointers to Block structs
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// Genesis block is the firsst block in the blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) // Passes empty PrevBlockHash []byte{}, since its the first block
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
