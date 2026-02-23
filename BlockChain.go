package main

import "github.com/boltdb/bolt"

type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	db          *bolt.DB
}

const dbFile = "blockchain.db"
const blockBucket = "blocks"

func (bc *BlockChain) addBlock(data string) {
	var lastHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		panic(err)
	}
	block := newBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		err := b.Put(block.Hash, block.Serialize())
		if err != nil {
			return err
		}
		err = b.Put([]byte("l"), block.Hash)
		if err != nil {
			return err
		}
		bc.tip = block.Hash
		return nil
	})
	if err != nil {
		panic(err)
	}
}
func newGenesisBlock() *Block {
	return newBlock("genesis block", []byte{})
}

func newBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))

		if b == nil {
			genesis := newGenesisBlock()
			b, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				return err
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				return err
			}
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				return err
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	bc := &BlockChain{tip, db}
	return bc
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	bci := &BlockChainIterator{bc.tip, bc.db}
	return bci
}

func (i *BlockChainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		encodedBlock := b.Get(i.CurrentHash)
		block = DeSerialize(encodedBlock)
		return nil
	})
	if err != nil {
		return nil
	}
	i.CurrentHash = block.PrevHash

	return block
}
