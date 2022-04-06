package src

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
	"math/big"
)

//UTXO
type Transaction struct {

	//1 transaction hash
	TxHash []byte

	//2.input
	//a slice to collect all TXInput
	Vins []*TXInput

	//3.output
	//a slice to collect all TXOutput
	Vouts []*TXOutput
}

//create transaction for genesis block

func NewCoinbaseTransAction(address string) *Transaction {

	//consume

	txInput := &TXInput{[]byte{}, -1, nil, []byte{}}

	txOutput := NewTXOutput(10, address)

	txCoinbase := &Transaction{[]byte{}, []*TXInput{txInput}, []*TXOutput{txOutput}}

	// set hash
	txCoinbase.HashTransaction()

	return txCoinbase
}

//SerializeBlock

func (tx *Transaction) HashTransaction() {

	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash := sha256.Sum256(result.Bytes())

	tx.TxHash = hash[:]
}

func NewSimpleTransaction(from, to string, amount int, blockchain *Blockchain, txs []*Transaction) *Transaction {

	money, spendableUTXODic := blockchain.FindSpendableUTXOS(from, amount, txs)

	//get a wallets
	//find the specific wallet with from
	wallets, _ := NewWallets()
	wallet := wallets.WalletsMap[from]

	//consume
	var txInputs []*TXInput
	for txHash, indexSlice := range spendableUTXODic {
		for _, index := range indexSlice {
			txHashBytes, _ := hex.DecodeString(txHash)
			txInput := &TXInput{txHashBytes, index, nil, wallet.PublicKey}
			txInputs = append(txInputs, txInput)
		}
	}

	//transfer
	var txOutputs []*TXOutput
	txOutput := NewTXOutput(int64(amount), to)
	txOutputs = append(txOutputs, txOutput)

	//the rest amount
	txOutput = NewTXOutput(int64(money)-int64(amount), from)
	txOutputs = append(txOutputs, txOutput)

	tx := &Transaction{[]byte{}, txInputs, txOutputs}

	// set hash
	tx.HashTransaction()

	//signature
	blockchain.SignTransaction(tx, wallet.PrivateKey)

	return tx
}

//judge the current transaction belongs to Coinbase

func (tx *Transaction) IsCoinbaseTransaction() bool {

	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1

}

func (tx *Transaction) Sign(privateKey ecdsa.PrivateKey, preTXs map[string]Transaction) {
	if tx.IsCoinbaseTransaction() {
		return
	}

	for _, vin := range tx.Vins {
		if preTXs[hex.EncodeToString(vin.TxHash)].TxHash == nil {
			log.Panic("ERROR:Previous transaction is not correct")
		}
	}

	txCopy := tx.TrimmedCopy()

	for inID, vin := range txCopy.Vins {
		prevTX := preTXs[hex.EncodeToString(vin.TxHash)]
		txCopy.Vins[inID].Signature = nil
		txCopy.Vins[inID].PublicKey = prevTX.Vouts[vin.Vout].Ripemd160Hash
		txCopy.TxHash = txCopy.Hash()
		txCopy.Vins[inID].PublicKey = nil

		r, s, err := ecdsa.Sign(rand.Reader, &privateKey, txCopy.TxHash)
		if err != nil {
			log.Panic(err)
		}
		signature := append(r.Bytes(), s.Bytes()...)
		tx.Vins[inID].Signature = signature
	}
}

//get a copy

func (tx *Transaction) TrimmedCopy() Transaction {

	var inputs []*TXInput
	var outputs []*TXOutput

	for _, vin := range tx.Vins {
		//no need signature and publicKey
		inputs = append(inputs, &TXInput{vin.TxHash, vin.Vout, nil, nil})
	}
	for _, vout := range tx.Vouts {
		outputs = append(outputs, &TXOutput{vout.Value, vout.Ripemd160Hash})
	}
	txCopy := Transaction{tx.TxHash, inputs, outputs}

	return txCopy
}

func (tx *Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	return encoded.Bytes()
}

func (tx *Transaction) Hash() []byte {
	txCopy := tx
	txCopy.TxHash = []byte{}
	hash := sha256.Sum256(txCopy.Serialize())
	return hash[:]
}

//verify signature

func (tx *Transaction) Verify(prevTXs map[string]Transaction) bool {
	if tx.IsCoinbaseTransaction() {
		return true
	}

	//if txHash from parameters didn't in current transaction, terminate
	for _, vin := range tx.Vins {
		if prevTXs[hex.EncodeToString(vin.TxHash)].TxHash == nil {
			log.Panic("ERROR:Previous transaction is not correct")
		}
	}

	txCopy := tx.TrimmedCopy()

	//get a curve
	curve := elliptic.P256()

	for inID, vin := range tx.Vins {
		//get a transaction
		prevTX := prevTXs[hex.EncodeToString(vin.TxHash)]
		//1.signature still is nil
		txCopy.Vins[inID].Signature = nil
		//2.assign Ripemd160Hash of output to PublicKey of input
		txCopy.Vins[inID].PublicKey = prevTX.Vouts[vin.Vout].Ripemd160Hash
		//3.get a new txHash
		txCopy.TxHash = txCopy.Hash()
		//4.reset PublicKey as nil
		txCopy.Vins[inID].PublicKey = nil

		//signature code

		r := big.Int{}
		s := big.Int{}
		sigLen := len(vin.Signature)
		//get the first half
		r.SetBytes(vin.Signature[:(sigLen / 2)])
		//get the second half
		s.SetBytes(vin.Signature[(sigLen / 2):])

		x := big.Int{}
		y := big.Int{}
		keyLen := len(vin.PublicKey)
		//get the first half
		x.SetBytes(vin.PublicKey[:(keyLen / 2)])
		//get the first half
		y.SetBytes(vin.PublicKey[(keyLen / 2):])

		rawPubKey := ecdsa.PublicKey{curve, &x, &y}
		if ecdsa.Verify(&rawPubKey, txCopy.TxHash, &r, &s) == false {
			return false
		}
	}
	return true
}
