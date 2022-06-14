package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type FDBConfig struct {
	ClusterFile string `json:"cluster_file" required:"true"`
}

func (c *config) FDB() FDBConfig {
	c.once.Do(func() interface{} {
		var result FDBConfig

		err := envconfig.Process("BTC_INDEXER", &result)

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out last block info"))
		}

		c.FDBConfig = result

		return nil
	})

	return c.FDBConfig
}
