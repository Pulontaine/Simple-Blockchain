package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()

	return block
}

type BlockChain struct {
	blocks []*Block
}

func (b *BlockChain) AddBlock(data string) {
	prevBlock := b.blocks[len(b.blocks)-1]
	block := NewBlock(data, prevBlock.Hash)
	b.blocks = append(b.blocks, block)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockchain()
	bc.AddBlock("First block")
	bc.AddBlock("Second block")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %X\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

}
