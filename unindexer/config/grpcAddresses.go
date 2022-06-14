package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type GrpcAddressesConfig struct {
	BtcIndexerGrpc string `fig:"btc" required:"true"`
}

func (c *config) GrpcAddresses() GrpcAddressesConfig {
	c.once.Do(func() interface{} {
		var result GrpcAddressesConfig
		err := figure.
			Out(&result).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "grpc-addresses")).
			Please()

		if err != nil {
			panic(err)
		}

		c.GrpcAddressesConfig = result

		return nil
	})

	return c.GrpcAddressesConfig
}
