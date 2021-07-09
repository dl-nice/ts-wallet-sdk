package wallet

type ChainType = int

const (
	CHAIN_BITCOIN ChainType = iota
	CHAIN_ETHEREUM
	CHAIN_NEM
	CHAIN_SUBSTRATE
)
const (
	HD_BTC = "m/44'/0'/0'/0/0"
	HD_ETH = "m/44'/60'/0'/0/0"
)
