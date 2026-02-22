package main

type BlockChain struct {
	blocks []*Block
}

func (bc *BlockChain) addBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := newBlock(data, prevBlock.hash)
	bc.blocks = append(bc.blocks, block)
}
func newGenesisBlock() *Block {
	return newBlock("genesis block", []byte{})
}

func newBlockChain() *BlockChain {
	return &BlockChain{[]*Block{newGenesisBlock()}}
}
