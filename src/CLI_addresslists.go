package src

import "fmt"

//print all wallet address

func (cli *CLI) AddressLists() []string {

	fmt.Println("print all wallets' address")
	wallets, _ := NewWallets()

	for address, _ := range wallets.WalletsMap {
		fmt.Println(address)
	}
	return nil
}
