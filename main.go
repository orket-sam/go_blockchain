package main

import (
	"encoding/json"
	"strconv"
	"time"
)

func main() {
	genesisBlock := Block{Index: 0, TimeStamp: time.Now(), Data: "Genesis Block", PreviousHash: ""}
	genesisBlock.CalculateHash()

	blockChain := BlockChain{genesisBlock}
	for i := 0; i < 9; i++ {
		blockChain.AddBlock(*GenerateBlock(blockChain[i], "Data "+strconv.Itoa(i+1)))

	}
	blockBytes, _ := json.MarshalIndent(blockChain, "", " ")
	println(string(blockBytes))
}
