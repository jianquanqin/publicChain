package src

import "fmt"

func (cli *CLI) TestMethod() {
	blockchain := BlockChainObject()

	utxoMap := blockchain.FindUTXOMap()
	defer blockchain.DB.Close()

	utxoSet := &UTXOSet{blockchain}
	utxoSet.ResetUTXOSet()

	var x []*TXOutputs
	for _, v := range utxoMap {
		x = append(x, v)
	}
	fmt.Println(x)
}
