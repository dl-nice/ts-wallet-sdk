package wallet

type EthereumAccount struct{}

func (*EthereumAccount) GenerateKeyPair(seed string, netWork string) (k *KeyPair, err error) {
	return
}
