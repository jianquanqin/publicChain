package src

import "fmt"

//print all address

func (cli *CLI) GetAddressList() {

	fmt.Println("Address list:")

	wallets, _ := NewWallets()
	for address, _ := range wallets.WalletMap {
		fmt.Println(address)
	}

}
