package src

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"github.com/boltdb/bolt"
	"log"
)

func HandleVersion(request []byte, bc *Blockchain) {

	var buff bytes.Buffer
	var payload Version

	//l := len(COMMAND_VERSION)
	dataBytes := request[COMMANDLENGTH:]

	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	bestHeight := bc.GetBestHeight()
	foreignerBestHeight := payload.BestHeight

	if bestHeight > foreignerBestHeight {
		SendVersion(payload.AddrFrom, bc)
	} else if bestHeight < foreignerBestHeight {
		SendGetBlocks(payload.AddrFrom)
	}

	if !NodeIsKnown(payload.AddrFrom) {
		knowNodes = append(knowNodes, payload.AddrFrom)
	}
}
func HandleGetBlocks(request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload GetBlocks

	dataBytes := request[COMMANDLENGTH:]

	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blockHashes := bc.GetBlockHashes()
	SendInv(payload.AddFrom, BLOCK_TYPE, blockHashes)

}
func HandleBlock(request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload BlockData

	dataBytes := request[COMMANDLENGTH:]

	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	block := payload.Block
	bc.AddBlock(block)

	//fmt.Println("reset dataBase...")
	//utxoSet := &UTXOSet{bc}
	//utxoSet.ResetUTXOSet()

	if len(transactionArray) > 0 {

		SendGetData(payload.AddrFrom, BLOCK_TYPE, transactionArray[0])
		transactionArray = transactionArray[1:]
	}
}

func HandleGetData(request []byte, bc *Blockchain) {

	var buff bytes.Buffer
	var payload GetData

	dataBytes := request[COMMANDLENGTH:]
	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)

	if err != nil {
		log.Panic(err)
	}

	if payload.Type == BLOCK_TYPE {
		block, err := bc.GetBlock([]byte(payload.Hash))
		if err != nil {
			return
		}

		SendBlock(payload.AddFrom, block)
	}
	if payload.Type == TX_TYPE {

		tx := memoryTxPool[hex.EncodeToString(payload.Hash)]
		SendTx(payload.AddFrom, tx)
	}

}
func HandleInv(request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload Inv

	dataBytes := request[COMMANDLENGTH:]
	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)

	if err != nil {
		log.Panic(err)
	}
	if payload.Type == BLOCK_TYPE {

		blockHash := payload.Items[0]
		SendGetData(payload.AddFrom, BLOCK_TYPE, blockHash)
		if len(payload.Items) >= 1 {
			transactionArray = payload.Items[1:]
		}
	}
	if payload.Type == TX_TYPE {

		txHash := payload.Items[0]
		if memoryTxPool[hex.EncodeToString(txHash)] == nil {
			SendGetData(payload.AddFrom, TX_TYPE, txHash)
		}

	}
}

func HandleTx(request []byte, bc *Blockchain) {

	var buff bytes.Buffer
	var payload Tx

	dataBytes := request[COMMANDLENGTH:]

	// 反序列化
	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	//-----

	tx := payload.Tx
	memoryTxPool[hex.EncodeToString(tx.TxHash)] = tx

	// 说明主节点自己
	if nodeAddress == knowNodes[0] {
		// 给矿工节点发送交易hash
		for _, nodeAddr := range knowNodes {

			if nodeAddr != nodeAddress && nodeAddr != payload.AddrFrom {
				SendInv(nodeAddr, TX_TYPE, [][]byte{tx.TxHash})
			}

		}
	}

	// 矿工进行挖矿验证
	// "" | 1DVFvyCK8qTQkLBTZ5fkh5eDSbcZVoHAsj
	if len(minerAddress) > 0 {

		utxoSet := &UTXOSet{bc}
		//
		//
		txs := []*Transaction{tx}

		//奖励
		coinbaseTx := NewCoinbaseTransAction(minerAddress)
		txs = append(txs, coinbaseTx)

		var _txs []*Transaction

		//fmt.Println("开始进行数字签名验证.....")

		for _, tx := range txs {

			//fmt.Printf("开始第%d次验证...\n",index)

			// 作业，数字签名失败
			if bc.VerifyTransaction(tx, _txs) != true {
				log.Panic("ERROR: Invalid transaction")
			}

			//fmt.Printf("第%d次验证成功\n",index)
			_txs = append(_txs, tx)
		}

		//fmt.Println("数字签名验证成功.....")

		//1. 通过相关算法建立Transaction数组
		var block *Block

		bc.DB.View(func(tx *bolt.Tx) error {

			b := tx.Bucket([]byte(blockTableName))
			if b != nil {

				hash := b.Get([]byte("l"))

				blockBytes := b.Get(hash)

				block = DeserializeBlock(blockBytes)

			}

			return nil
		})

		//2. 建立新的区块
		block = Newblock(block.Height+1, block.Hash, txs)

		//将新区块存储到数据库
		bc.DB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blockTableName))
			if b != nil {

				b.Put(block.Hash, block.Serialize())

				b.Put([]byte("l"), block.Hash)

				bc.Tip = block.Hash

			}
			return nil
		})
		utxoSet.Update()
		SendBlock(knowNodes[0], block) //changed
	}
}

func HandleAddr(request []byte, bc *Blockchain) {}
