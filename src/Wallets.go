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

const walletFile = "Wallets.dat"

//a struct to collect wallet

type Wallets struct {
	WalletMap map[string]*Wallet
}

//creates a new wallets

func NewWallets() (*Wallets, error) {

	//to find the file,if there is no file,create a one
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		//create a struct
		wallets := &Wallets{}
		//access its property and initialize it
		wallets.WalletMap = make(map[string]*Wallet)
		return wallets, err
	}

	//read the file
	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	//Deserialize
	var wallets Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}
	return &wallets, nil
}

func (wallets *Wallets) CreatNewWallet() {
	//create a new wallet
	wallet := NewWallet()
	fmt.Printf("Wallet address: %s\n", wallet.GetAddress())
	//throw wallet into wallets (a map),instantiate structs' element
	wallets.WalletMap[string(wallet.GetAddress())] = wallet
	wallets.SaveWallets()
}

func (wallets *Wallets) SaveWallets() {

	//serialize

	var content bytes.Buffer //define a buffer

	gob.Register(elliptic.P256()) //register with a curve,for serializing all types

	encoder := gob.NewEncoder(&content) //create a encoder
	err := encoder.Encode(&wallets)     //serialize wallets and throw it into buffer
	if err != nil {
		log.Panic(err)
	}
	//write serialized file to dat
	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
