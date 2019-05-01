package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	// Block is Create
	Timestamp int64
	// Valuable Information
	Data []byte
	// Hash of the Previous Block
	PrevBlockHash []byte
	// Hash of the Block
	Hash []byte
}

type Blockchain struct {
	blocks []*Block
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

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("제네시스 블록", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// main()
func main() {
	bc := NewBlockchain()

	bc.AddBlock("'1' 코인을 Haru에게 전송")
	bc.AddBlock("'2' 코인을 Haru에게 전송")

	for _, block := range bc.blocks {
		fmt.Printf("이전 해쉬 : %x\n", block.PrevBlockHash)
		fmt.Printf("데이터 : %s\n", block.Data)
		fmt.Printf("해쉬 : %x\n", block.Hash)
		fmt.Println()
	}
}
