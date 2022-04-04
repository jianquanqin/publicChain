package src

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
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

	return tx
}

//judge the current transaction belongs to Coinbase

func (tx *Transaction) IsCoinbaseTransaction() bool {

	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1

}
