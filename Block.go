package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Data      []byte
	Hash      []byte
	PrevHash  []byte
	Nonce     int
}

func (b *Block) SetHash() {
	//strconv.FormatInt(b.Timestamp, 10) 进制转换
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	header := bytes.Join([][]byte{b.PrevHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(header)
	b.Hash = hash[:]

}
func newBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), []byte{}, prevHash, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

func DeSerialize(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewBuffer(d))
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}
	return &block
}
