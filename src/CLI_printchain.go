package BLC

import (
	"fmt"
	"os"
)

func (cli CLI) printChain() {
	if DBExists() == false {
		fmt.Println("database didn't exist")
		os.Exit(1)
	}
	blockchain := BlockChainObject()
	defer blockchain.DB.Close()
	blockchain.PrintChain()
}
