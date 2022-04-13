package src

func (cli CLI) printChain(nodeID string) {

	blockchain := BlockChainObject(nodeID)
	defer blockchain.DB.Close()
	blockchain.PrintChain()

}
