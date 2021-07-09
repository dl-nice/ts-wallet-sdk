package wallet

import "github.com/tyler-smith/go-bip39"

type KeyPair struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Address    string `json:"address"`
	NetWork    string `json:"netWork"`
	ChainType  string `json:"chainType"`
}
type Wallet struct {
	KeyPair  []KeyPair
	Mnemonic string `json:"mnemonic"`
}

type Accounter interface {
	GenerateKeyPair(seed string, netWork string) (*KeyPair, error)
}

func GenerateMnemonic(password string) (seed []byte, err error) {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed = bip39.NewSeed(mnemonic, password)
	return
}
