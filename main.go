package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := newBlockChain()

	bc.addBlock("ciallo")
	bc.addBlock("0721")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.prevHash)
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Hash: %x\n", block.hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
