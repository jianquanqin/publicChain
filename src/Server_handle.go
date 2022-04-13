package src

import (
	"bytes"
	"encoding/gob"
	"log"
)

func handleAddr(request []byte, bc *Blockchain) {}
func handleVersion(request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload Version

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

		//sendGetBlocks(payload.AddrFrom)
	}
}
func handleBlock(request []byte, bc *Blockchain)     {}
func handleInv(request []byte, bc *Blockchain)       {}
func handleGetblocks(request []byte, bc *Blockchain) {}
func handleTx(request []byte, bc *Blockchain)        {}
func handleGetData(request []byte, bc *Blockchain)   {}
