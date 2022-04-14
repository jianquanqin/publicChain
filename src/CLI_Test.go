package src

import "fmt"

func (cli *CLI) Test(nodeID string) {
	fmt.Println("testing")

	blockchain := BlockChainObject(nodeID)
	defer blockchain.DB.Close()

	utxoSet := &UTXOSet{blockchain}
	utxoSet.ResetUTXOSet()
}
