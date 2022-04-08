package src

import "fmt"

//a struct to collect wallet

type Wallets struct {
	WalletMap map[string]*Wallet
}

//creates a new wallets

func NewWallets() *Wallets {
	//create a struct
	wallets := &Wallets{}
	//access its property and initialize it
	wallets.WalletMap = make(map[string]*Wallet)
	return wallets
}

func (wallets *Wallets) CreatNewWallet() {
	//create a new wallet
	wallet := NewWallet()
	fmt.Printf("Wallet address is: \n%s\n", wallet.GetAddress())
	//throw wallet into wallets (a map)
	wallets.WalletMap[string(wallet.GetAddress())] = wallet
}
