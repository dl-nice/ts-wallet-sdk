package wallet

type BitcoinAccount struct{}

func (*BitcoinAccount) GenerateKeyPair(seed string, netWork string) (k *KeyPair, err error) {
	return
}
