package src

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"log"
)

//create a wallet

//1.Produce a pubKey and privateKey
//2.for pubKey: first hash256 and then hash160
//3.select the last four nums(add check digit)
//4.add version(1) \ hash160(20) \ add check digit(4)
//5.Base58Encode(25),then you can get address
//6.Send address to sender,the sender decoded address to pubKey and sign
//7.You decrypt transaction with your privateKey

const version = byte(0x00)
const addressChecksumLen = 4

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

//produce a publicKey through privateKey

func NewKeyPair() (ecdsa.PrivateKey, []byte) {

	//get a parameter
	curve := elliptic.P256()

	//use parameter to privateKey
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	//because privateKey is a struct, pubKey is one if its properties
	//we can get pubKey through privateKey
	//
	pubKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	return *privateKey, pubKey
}

// create a wallet

func NewWallet() *Wallet {

	privateKey, pubKey := NewKeyPair()

	return &Wallet{privateKey, pubKey}

}

//first hash256 then hash160

func Ripemd160(publicKey []byte) []byte {

	//1.hash256
	hash256 := sha256.New()
	hash256.Write(publicKey)
	hash := hash256.Sum(nil)

	//2.ripemd160
	ripeMd160 := ripemd160.New()
	ripeMd160.Write(hash)

	return ripeMd160.Sum(nil)
}

//CheckSum

func CheckSum(rpm160 []byte) []byte {

	hash1 := sha256.Sum256(rpm160)
	hash2 := sha256.Sum256(hash1[:])

	return hash2[:addressChecksumLen]

}

//get address

func (wallet *Wallet) GetAddress() []byte {
	//1 hash256.160 publicKey
	rpm160 := Ripemd160(wallet.PublicKey)
	versionRipedmd160hash := append([]byte{version}, rpm160...)

	checkSumBytes := CheckSum(versionRipedmd160hash)
	bytes := append(versionRipedmd160hash, checkSumBytes...)

	return Base58Encode(bytes)
}

//verify an address

func (wallet *Wallet) VerifyAddress(address []byte) bool {

	//first decode
	versionPubKeyCheckSumBytes := Base58Decode(address)
	//fmt.Println(versionPubKeyCheckSumBytes)

	CheckSumBytes := versionPubKeyCheckSumBytes[len(versionPubKeyCheckSumBytes)-addressChecksumLen:]
	versionRipemd160 := versionPubKeyCheckSumBytes[:len(versionPubKeyCheckSumBytes)-addressChecksumLen]

	//verify
	checkBytes := CheckSum(versionRipemd160)
	if bytes.Compare(CheckSumBytes, checkBytes) == 0 {
		return true
	}
	return false
}
