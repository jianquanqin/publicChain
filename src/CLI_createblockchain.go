package src

func (cli *CLI) creatGenesisBlockChain(address string) {

	// create coinbase transaction
	blockchain := CreatBlockchainWithGenesisBlock(address)
	//remember to close db
	defer blockchain.DB.Close()

	utxoSet := &UTXOSet{blockchain}
	utxoSet.ResetUTXOSet()
}
