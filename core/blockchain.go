package core

type BlockChain struct {
	Blocks []Block
}

func NewBlockChain() *BlockChain {
	ch := &BlockChain{}

	return ch
}

func (bc *BlockChain) AddBlock() *BlockChain {


	return bc
}

func (bc *BlockChain) Validate() bool {

	return true
}