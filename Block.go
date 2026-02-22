package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	timestamp int64
	data      []byte
	hash      []byte
	prevHash  []byte
	nonce     int
}

func (b *Block) SetHash() {
	//strconv.FormatInt(b.timestamp, 10) 进制转换
	timestamp := []byte(strconv.FormatInt(b.timestamp, 10))
	header := bytes.Join([][]byte{b.prevHash, b.data, timestamp}, []byte{})
	hash := sha256.Sum256(header)
	b.hash = hash[:]

}
func newBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), []byte{}, prevHash, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.run()
	block.hash = hash
	block.nonce = nonce

	return block
}
