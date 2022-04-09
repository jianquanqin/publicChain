package src

import "fmt"

func (cli *CLI) createWallet() {

	wallets, _ := NewWallets()
	wallets.CreatNewWallet()

	fmt.Println(len(wallets.WalletMap))
}
