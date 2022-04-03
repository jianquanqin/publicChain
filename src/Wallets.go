package src

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const walletFile = "Wallet.dat"

type Wallets struct {
	WalletsMap map[string]*Wallet
}

//create collection of wallet

func NewWallets() (*Wallets, error) {

	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		wallets := &Wallets{}
		wallets.WalletsMap = make(map[string]*Wallet)
		return wallets, err
	}
	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	var wallets Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}

	return &wallets, nil
}

// create a  new wallet

func (w *Wallets) CreateNewWallet() {

	wallet := NewWallet()
	fmt.Printf("Address：%s\n", wallet.GetAddress())
	w.WalletsMap[string(wallet.GetAddress())] = wallet

	w.SaveWallets()
}

//write the info of wallets to file

func (w *Wallets) SaveWallets() {
	var content bytes.Buffer

	//the purpose of register is to Serialization
	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(&w)

	if err != nil {
		log.Panic(err)
	}

	// write the serialized data to file
	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
