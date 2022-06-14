package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type BtcAuthConfig struct {
	Url string `json:"url" required:"true"`
	Key string `json:"key" required:"true"`
}

func (c *config) BtcAuth() BtcAuthConfig {
	c.once.Do(func() interface{} {
		var result BtcAuthConfig

		err := envconfig.Process("BTC_INDEXER", &result)

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out last block info"))
		}

		c.BtcAuthConfig = result

		return nil
	})

	return c.BtcAuthConfig
}
