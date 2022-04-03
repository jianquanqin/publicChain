package src

type TXOutput struct {

	//1.output amount
	Value int64

	//2.recipient address
	ScriptPublicKey string //pubKey
}

// judge the current output belong to one's

func (txOutput *TXOutput) UnlockScriptPublicKeyWithAddress(address string) bool {
	return txOutput.ScriptPublicKey == address
}
