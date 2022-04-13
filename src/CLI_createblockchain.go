package src

func (cli *CLI) creatGenesisBlockChain(address string, nodeID string) {

	// create coinbase transaction
	blockchain := CreatBlockchainWithGenesisBlock(address, nodeID)
	//remember to close db
	defer blockchain.DB.Close()

	utxoSet := &UTXOSet{blockchain}
	utxoSet.ResetUTXOSet()
}
