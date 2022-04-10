package src

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

//now we just want to have a sheet to store all available TXOutput
//so when we check or transfer,just need to look for this sheet,it's a good way to save resources

const utxoTableName = "utxoTableName"

type UTXOSet struct {
	BlockChain *Blockchain
}

func (utxoSet *UTXOSet) ResetUTXOSet() {
	//update the table
	err := utxoSet.BlockChain.DB.Update(func(tx *bolt.Tx) error {
		//get the table
		b := tx.Bucket([]byte(utxoTableName))
		if b != nil {
			//if you find the table, delete it
			err := tx.DeleteBucket([]byte(utxoTableName))
			if err != nil {
				log.Panic("delete utxoTable failed")
			}
		}
		//and create a new one
		b, _ = tx.CreateBucket([]byte(utxoTableName))
		if b != nil {
			//get txOutput
			txOutputsMap := utxoSet.BlockChain.FindUTXOMap()
			//get key and value
			for keyHash, txOutputs := range txOutputsMap {
				txHash, _ := hex.DecodeString(keyHash)

				//serialize value and put it in db with key
				b.Put(txHash, txOutputs.Serialize())

			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

//find available outputs

func (utxoSet *UTXOSet) FindUTXOForAddress(address string) []*UTXO {

	var utxos []*UTXO

	utxoSet.BlockChain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoTableName))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key = %x, \nvalue = %x\n", k, v)

			txOutputs := DeserializeTXOutputs(v)
			for _, utxo := range txOutputs.UTXOS {
				if utxo.Output.UnlockScriptPublicKeyWithAddress(address) {
					utxos = append(utxos, utxo)
				}
			}
		}
		return nil
	})
	return utxos
}

//get balance from utxo

func (utxoSet *UTXOSet) GetBalance(address string) int64 {

	UTXOS := utxoSet.FindUTXOForAddress(address)
	var amount int64

	for _, utxo := range UTXOS {
		amount += utxo.Output.Value
	}
	return amount
}
