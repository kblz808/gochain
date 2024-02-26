package main

import (
	"fmt"
	"gochain/blockchain"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("block after genesis")
	chain.AddBlock("another block")
	chain.AddBlock("third block in the chain")

	for _, block := range chain.Blocks {
		fmt.Printf("prev hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
