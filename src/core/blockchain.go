package core

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]   //取数组最后一元素
	newBlock := NewBlock(data, prevBlock.Hash) //调用NewBlock方法创建新的区块
	bc.Blocks = append(bc.Blocks, newBlock)    //将新区快添加
}

func NewBlockChain() *Blockchain {
	return &Blockchain{Blocks: []*Block{NewGenesisBlock()}}
}
