package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block keeps block headers
type Block struct {
	Timestamp     int64  // Time when block was created
	Data          []byte // Valuable information
	PrevBlockHash []byte
	Hash          []byte
}

// SetHash calculates and sets block hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))                       // Convert timestamp to string and to byte slice
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) // Concatinate byte slices
	hash := sha256.Sum256(headers)                                                // Computes hash, returns array of 32 bytes

	b.Hash = hash[:] // Converts hash into slice and assigns it to the b.Hash field in the block
}

// NewBlock creates and returns Block
// Golang convention to use constructor functions New<Type> to create and return a pointer to a new instance of a struct
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}
