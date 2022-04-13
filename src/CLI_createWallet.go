package src

import "fmt"

func (cli *CLI) createWallet(nodeID string) {

	wallets, _ := NewWallets(nodeID)
	wallets.CreatNewWallet(nodeID)

	fmt.Println(len(wallets.WalletMap))
}
