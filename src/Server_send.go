package src

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

//get and pack the info,then send to main node

func SendVersion(toAddress string, bc *Blockchain) {

	bestHeight := bc.GetBestHeight()
	payload := GobEncode(Version{NODE_VERSION, bestHeight, nodeAddress})

	request := append(CommandTOBytes(COMMAND_VERSION), payload...)

	SendData(toAddress, request)
}
func SendGetBlocks(toAddress string) {

	payload := GobEncode(GetBlocks{nodeAddress})
	request := append(CommandTOBytes(COMMAND_GETBLOCKS), payload...)
	SendData(toAddress, request)

}
func SendData(toAddress string, request []byte) {

	fmt.Println("the client sends message to the server...")
	conn, err := net.Dial("tcp", toAddress)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	//attach message
	_, err1 := io.Copy(conn, bytes.NewReader(request))
	if err1 != nil {
		log.Panic(err1)
	}
}
func SendInv(toAddress, kind string, hashes [][]byte) {

	payload := GobEncode(Inv{nodeAddress, kind, hashes})
	request := append(CommandTOBytes(COMMAND_INV), payload...)
	SendData(toAddress, request)
}

func SendGetData(toAddress, kind string, blockHash []byte) {

	payload := GobEncode(GetData{nodeAddress, kind, blockHash})
	request := append(CommandTOBytes(COMMAND_GETDATA), payload...)

	SendData(toAddress, request)
}

func SendBlock(toAddress string, block *Block) {
	payload := GobEncode(BlockData{nodeAddress, block})
	request := append(CommandTOBytes(COMMAND_BLOCK), payload...)
	SendData(toAddress, request)
}

func SendTx(toAddress string, tx *Transaction) {

	payload := GobEncode(Tx{nodeAddress, tx})

	request := append(CommandTOBytes(COMMAND_TX), payload...)

	SendData(toAddress, request)
}
