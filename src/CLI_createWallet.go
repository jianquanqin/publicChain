package src

import "fmt"

func (cli *CLI) createWallet() {
	wallets := NewWallets()
	wallets.CreatNewWallet()

	fmt.Println(wallets.WalletMap)
}
