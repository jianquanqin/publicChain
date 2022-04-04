package src

import "bytes"

type TXInput struct {
	//1 transaction hash
	TxHash []byte

	//2.index of TXOutput
	Vout int

	//3.signature and pubKey
	//if a person has satisfied the following conditions,it indicates that he is the owner of the address
	Signature []byte
	PublicKey []byte
}

// judge the current money belong to one's

func (txInput *TXInput) UnlockRipedmd160Hash(ripemd160Hash []byte) bool {

	publicKey := Ripemd160(txInput.PublicKey)
	//ripemd160Hash comes from output
	//pubKey comes from user who wants to spend money
	//Compare two public keys from different sources
	return bytes.Compare(publicKey, ripemd160Hash) == 0
}
