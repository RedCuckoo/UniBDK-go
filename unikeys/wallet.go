package unikeys

type Wallet interface {
	GetBlockchainName() Blockchain
	GetAddress() (string, error)
	GetKey() *Key
}

type BtcWallet struct {
	key *Key
}

func (wallet BtcWallet) GetBlockchainName() Blockchain {
	return BTC
}

func (wallet BtcWallet) GetAddress() (string, error) {
	return wallet.key.BtcAddress()
}

func (wallet BtcWallet) GetKey() *Key {
	return wallet.key
}

type EthWallet struct {
	key *Key
}

func (wallet EthWallet) GetBlockchainName() Blockchain {
	return ETH
}

func (wallet EthWallet) GetAddress() (string, error) {
	return wallet.key.EthAddress()
}

func (wallet EthWallet) GetKey() *Key {
	return wallet.key
}
