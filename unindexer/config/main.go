package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Config interface {
	comfig.Logger
	comfig.Listenerer

	FDB() FDBConfig
	BtcAuth() BtcAuthConfig
	EthAuth() EthAuthConfig
	GrpcAddresses() GrpcAddressesConfig
}

type config struct {
	comfig.Logger
	comfig.Listenerer
	getter kv.Getter
	once   comfig.Once

	FDBConfig           FDBConfig
	BtcAuthConfig       BtcAuthConfig
	EthAuthConfig       EthAuthConfig
	GrpcAddressesConfig GrpcAddressesConfig
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}
