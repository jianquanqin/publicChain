package src

func (cli *CLI) creatGenesisBlockChain(address string) {

	// create coinbase transaction
	blockchain := CreatBlockchainWithGenesisBlock(address)
	blockchain.DB.Close()

}
