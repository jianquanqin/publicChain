package src

type TXInput struct {
	//1 transaction hash
	TxHash []byte

	//2.index of TXOutput
	Vout int

	//3.signature
	ScriptSig string
}

// judge the current money belong to one's

func (txInput *TXInput) UnlockWithAddress(address string) bool {
	return txInput.ScriptSig == address
}
