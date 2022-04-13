package src

import "fmt"

func (cli *CLI) StartNode(nodeID string, mineAddress string) {

	if IsValidForAddress([]byte(mineAddress)) || mineAddress == "" {

		//start server
		fmt.Printf("start the server, localhost:%s\n", nodeID)
		StartServer(nodeID, mineAddress)

	} else {
		fmt.Println("reward address is invalid")
	}

}
