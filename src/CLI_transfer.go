package src

import (
	"fmt"
	"os"
)

func (cli CLI) send(from []string, to []string, amount []string) {

	if DBExists() == false {
		fmt.Println("data didn't exist")
		os.Exit(1)
	}

	//mine a new clock
	blockchain := BlockChainObject()
	defer blockchain.DB.Close()
	blockchain.MineNewBlock(from, to, amount)

	//when finished the transaction, update the data
	utxoSet := &UTXOSet{blockchain}
	utxoSet.Update()

}
