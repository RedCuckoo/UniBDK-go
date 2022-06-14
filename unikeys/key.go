package unikeys

import (
	"crypto/ecdsa"
	"errors"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/crypto"
)

func NewKey(mnemonic string) (*Key, error) {
	seed, err := createHashedSeed(mnemonic)
	if err != nil {
		return nil, err
	}
	master, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	result := Key{extended: master}

	result.privateBTC, err = master.ECPrivKey()
	if err != nil {
		return nil, err
	}
	result.publicBTC, err = master.ECPubKey()
	if err != nil {
		return nil, err
	}
	result.privateETH = result.privateBTC.ToECDSA()
	result.publicETH = result.publicBTC.ToECDSA()

	return &result, nil
}

type Key struct {
	extended *hdkeychain.ExtendedKey

	privateBTC *btcec.PrivateKey
	publicBTC  *btcec.PublicKey

	privateETH *ecdsa.PrivateKey
	publicETH  *ecdsa.PublicKey
}

func (key *Key) WalletByBlockchain(blockchain string) (Wallet, error) {
	if blockchain == "" {
		return nil, errors.New("empty blockchain name")
	}

	switch Blockchain(blockchain) {
	case BTC:
		return BtcWallet{key}, nil
	case ETH:
		return EthWallet{key}, nil
	default:
		return nil, errors.New("not supported blockchain name")
	}
}

func (key *Key) BtcAddress() (string, error) {
	pubKey, err := btcutil.NewAddressPubKey(key.publicBTC.SerializeCompressed(), &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	return pubKey.EncodeAddress(), nil
}

func (key *Key) EthAddress() (string, error) {
	return crypto.PubkeyToAddress(*key.publicETH).Hex(), nil
}
