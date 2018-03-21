package core

import (
	"time"
	"github.com/cbergoon/merkletree"
)

type Block struct {
	Index        int64
	CreationTime time.Time
	CommitTime   time.Time
	Transactions merkletree.MerkleTree
}

func NewBlock() *Block {
	b := &Block{}

	return b
}

func (b *Block) GenerateGenesis() {

}

func (b *Block) Generate() {

}

func (b *Block) Validate() bool {

	return true
}