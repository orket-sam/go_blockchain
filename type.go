package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index        int
	TimeStamp    time.Time
	Data         string
	PreviousHash string
	Hash         string
}

type BlockChain []Block

func (b *Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) + b.PreviousHash + b.Data + b.TimeStamp.String()
	hash := sha256.New()
	hash.Write([]byte(record))

	hashed := hash.Sum(nil)
	b.Hash = hex.EncodeToString(hashed)
	return b.Hash

}

func GenerateBlock(prevBlock Block, data string) *Block {
	newBlock := Block{Index: prevBlock.Index + 1, TimeStamp: time.Now(), Data: data, PreviousHash: prevBlock.PreviousHash}
	newBlock.CalculateHash()
	return &newBlock
}

func IsBlockValid(prevBlock, newBlock Block) bool {
	if prevBlock.Index+1 != newBlock.Index {
		return false
	}
	if prevBlock.Hash != newBlock.PreviousHash {
		return false
	}
	if newBlock.CalculateHash() != newBlock.Hash {
		return false
	}
	return true
}

func (blockChain *BlockChain) AddBlock(newBlock Block) {
	// prevBlock := (*blockChain)[newBlock.Index-1]
	// if IsBlockValid(prevBlock, newBlock) {
	*blockChain = append(*blockChain, newBlock)
	// }
}
