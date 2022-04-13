package src

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
)

type getdata struct {
	AddFrom string
	Type    string
	ID      []byte
}

type inv struct {
	AddFrom string
	Type    string
	Items   [][]byte
}

type tx struct {
	AddFrom     string
	Transaction []byte
}

//main node address

var konwNodes = []string{"localhost:3000"}
var nodeAddress string

func StartServer(nodeID string, mineAddress string) {

	//current node address
	nodeAddress = fmt.Sprintf("localhost:%s", nodeID)

	ln, err := net.Listen(PROTOCOL, nodeAddress)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()

	bc := BlockChainObject(nodeID)

	//when the node is not main node
	//send some info to main node
	if nodeAddress != konwNodes[0] {
		//konwNodes[0] is main node
		SendVersion(konwNodes[0], bc)
	}

	for {

		//receive data from client
		conn, err1 := ln.Accept()
		if err != nil {
			log.Panic(err1)
		}
		go HandleConnection(conn)
	}
}
func HandleConnection(conn net.Conn) {
	//read data from client
	request, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Panic(err)
	}

	command := bytesToCommand(request[:COMMANDLENGTH])
	fmt.Printf("received %s command\n", command)

	bc := BlockChainObject("test")

	switch command {
	case COMMAND_VERSION:
		handleVersion(request, bc)
	case COMMAND_ADDR:
		handleAddr(request, bc)
	case COMMAND_BLOCK:
		handleBlock(request, bc)
	case COMMAND_GETBLOCKS:
		handleGetblocks(request, bc)
	case COMMAND_GETDATA:
		handleGetData(request, bc)
	case COMMAND_INV:
		handleInv(request, bc)
	case COMMAND_TX:
		handleTx(request, bc)
	default:
		fmt.Println("Unknown command!")
	}
	defer conn.Close()

}

//get and pack the info,then send to main node

func SendVersion(toAddress string, bc *Blockchain) {
	bestHeight := bc.GetBestHeight()

	payload := GobEncode(Version{NODE_VERSION, bestHeight, nodeAddress})

	request := append(CommandTOBytes(COMMAND_VERSION), payload...)

	SendData(toAddress, request)
}

func SendData(to string, data []byte) {

	fmt.Println("a client sends message to the server...")
	conn, err := net.Dial("tcp", to)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	//attach message
	_, err1 := io.Copy(conn, bytes.NewReader(data))
	if err1 != nil {
		log.Panic(err1)
	}
}
