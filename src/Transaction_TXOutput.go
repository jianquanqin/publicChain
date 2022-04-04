package src

import "bytes"

type TXOutput struct {

	//1.output amount
	Value int64

	//2.recipient address
	Ripedmd160Hash []byte //username
}

// judge the current output belong to one's

// judge the current output belong to one's

func (txOutput *TXOutput) UnlockScriptPublicKeyWithAddress(address string) bool {

	//decode address
	publicKeyHash := Base58Decode([]byte(address))
	//get ripemdhash160
	hash160 := publicKeyHash[1 : len(publicKeyHash)-4]
	//compare two 哈希160 from different sources
	return bytes.Compare(txOutput.Ripedmd160Hash, hash160) == 0
}

// create a lock to get ripemd160hash

func (txOutput *TXOutput) Lock(address string) {
	publicKey := Base58Decode([]byte(address))
	txOutput.Ripedmd160Hash = publicKey[1 : len(publicKey)-4]
}

//create a new output

func NewTXOutput(value int64, address string) *TXOutput {
	txOutput := &TXOutput{value, nil}

	//set Ripemd160Hash
	txOutput.Lock(address)

	return txOutput
}


