package src

import (
	"fmt"
	"os"
)

func (cli *CLI) StartNode(nodeID string, minerAddress string) {

	if IsValidForAddress([]byte(minerAddress)) || minerAddress == "" {

		//start server
		fmt.Printf("start the server, localhost:%s\n", nodeID)
		StartServer(nodeID, minerAddress)

	} else {
		fmt.Println("reward address is invalid")
		os.Exit(0)
	}
}
