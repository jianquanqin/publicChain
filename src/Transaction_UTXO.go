package src

type UTXO struct {
	TxHash []byte
	Index  int
	Output *TXOutput
}
