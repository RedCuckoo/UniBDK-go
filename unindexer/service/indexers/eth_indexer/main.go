package eth_indexer

import (
	"github.com/RedCuckoo/UniBDK-go/unindexer/config"
	"github.com/RedCuckoo/UniBDK-go/unindexer/fdbutils"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
)

type Service struct {
	log           *logan.Entry
	FoundationDB  *fdbutils.FDB
	StartBlock    int64
	EndBlock      int64
	requestUrl    string
	requestApiKey string
	client        *ethclient.Client
	batchSize     int
}

func New(cfg config.Config) *Service {
	log := cfg.Log().WithField("eth_indexer", "indexer")

	fdbConnect, err := fdbutils.Connect(cfg)
	if err != nil {
		panic(err)
	}

	ethAuth := cfg.EthAuth()

	return &Service{
		log:    log,
		client: ethAuth.EthClient,

		FoundationDB:  fdbConnect,
		requestUrl:    ethAuth.Url,
		requestApiKey: ethAuth.Key,
		batchSize:     1,
	}
}
