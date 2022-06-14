package unikeys

import "github.com/ethereum/go-ethereum/crypto"

type Wallet interface {
	GetBlockchainName() Blockchain
	GetAddress() (string, error)
	GetKey() *Key
	SignData([]byte) ([]byte, error)
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

func (wallet BtcWallet) SignData(data []byte) ([]byte, error) {
	// TODO
	return crypto.Sign(data, wallet.key.privateETH)
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

func (wallet EthWallet) SignData(data []byte) ([]byte, error) {
	return crypto.Sign(data, wallet.key.privateETH)
}
