package main

import (
	"core"
	"fmt"
	"strconv"
)

func main() {
	bc := core.NewBlockChain()
	bc.AddBlock("send 1 coin to lwq")
	bc.AddBlock("send 2 coin to lwq")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev.hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()

		pow := core.NewProofOfWork(block)
		fmt.Printf("Pow:%s\n", strconv.FormatBool(pow.Validata()))
		fmt.Println()
	}

}
