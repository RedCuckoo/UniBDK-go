package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type EthAuthConfig struct {
	Url       string            `json:"url" required:"true"`
	Key       string            `json:"key" required:"true"`
	EthClient *ethclient.Client `json:"ethClient"`
}

func (c *config) EthAuth() EthAuthConfig {
	c.once.Do(func() interface{} {
		var result EthAuthConfig

		err := envconfig.Process("ETH_INDEXER", &result)

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out eth indexer info"))
		}

		c.EthAuthConfig = result
		dial, err := ethclient.Dial(fmt.Sprintf("%s%s", result.Url, result.Key))
		if err != nil {
			panic(errors.Wrap(err, "failed to dial eth client"))
		}
		c.EthAuthConfig.EthClient = dial
		return nil
	})

	return c.EthAuthConfig
}
