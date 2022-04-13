package src

import "fmt"

//print all address

func (cli *CLI) GetAddressList(nodeID string) {

	fmt.Println("Address list:")

	wallets, _ := NewWallets(nodeID)
	for address, _ := range wallets.WalletMap {
		fmt.Println(address)
	}

}
