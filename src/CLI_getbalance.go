package src

import "fmt"

func (cli CLI) getBalance(address string) {

	fmt.Println("Address：" + address)

	blockchain := BlockChainObject()
	defer blockchain.DB.Close()

	amount := blockchain.GetBalance(address)

	fmt.Printf("%s has %d token\n", address, amount)
}
