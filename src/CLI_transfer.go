package BLC

import (
	"fmt"
	"os"
)

func (cli CLI) send(from []string, to []string, amount []string) {

	if DBExists() == false {
		fmt.Println("data didn't exist")
		os.Exit(1)
	}

	blockchain := BlockChainObject()
	defer blockchain.DB.Close()
	blockchain.MineNewBlock(from, to, amount)
}
