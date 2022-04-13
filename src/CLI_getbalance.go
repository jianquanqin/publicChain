package src

import "fmt"

func (cli CLI) getBalance(address string, nodeID string) {

	//fmt.Println("Address：" + address)

	blockchain := BlockChainObject(nodeID)
	defer blockchain.DB.Close()

	utxoSet := &UTXOSet{blockchain}
	amount := utxoSet.GetBalance(address)

	fmt.Printf("%s has %d token\n", address, amount)
}
