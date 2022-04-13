package src

type Version struct {
	Version    int64
	BestHeight int64  //the block height of current node
	AddrFrom   string //the address of current node
}
