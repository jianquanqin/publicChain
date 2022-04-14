package src

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func StartServer(nodeID string, minerAddress string) {

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
	if nodeAddress != knowNodes[0] {
		//knowNodes[0] is main node
		SendVersion(knowNodes[0], bc)
	}

	for {

		//receive data from client
		conn, err1 := ln.Accept()
		if err != nil {
			log.Panic(err1)
		}
		go HandleConnection(conn, bc)
	}
}
func HandleConnection(conn net.Conn, bc *Blockchain) {
	//read data from client
	request, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("received %s command\n", request[:COMMANDLENGTH])
	command := bytesToCommand(request[:COMMANDLENGTH])
	//fmt.Println(command)

	switch command {
	case COMMAND_VERSION:
		HandleVersion(request, bc)
	case COMMAND_ADDR:
		HandleAddr(request, bc)
	case COMMAND_BLOCK:
		HandleBlock(request, bc)
	case COMMAND_GETBLOCKS:
		HandleGetBlocks(request, bc)
	case COMMAND_GETDATA:
		HandleGetData(request, bc)
	case COMMAND_INV:
		HandleInv(request, bc)
	case COMMAND_TX:
		HandleTx(request, bc)
	default:
		fmt.Println("Unknown command!")
	}
	defer conn.Close()
}

func NodeIsKnown(addr string) bool {
	for _, node := range knowNodes {
		if node == addr {
			return true
		}
	}

	return false
}
