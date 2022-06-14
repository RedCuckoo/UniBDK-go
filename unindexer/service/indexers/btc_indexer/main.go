package btc_indexer

import (
	"github.com/RedCuckoo/UniBDK-go/unindexer/config"
	"github.com/RedCuckoo/UniBDK-go/unindexer/fdbutils"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
	"time"
)

type Service struct {
	log           *logan.Entry
	FoundationDB  *fdbutils.FDB
	StartBlock    int64
	EndBlock      int64
	httpClient    *http.Transport
	requestUrl    string
	requestApiKey string
}

func New(cfg config.Config) *Service {
	log := cfg.Log().WithField("btc_indexer", "indexer")

	fdbConnect, err := fdbutils.Connect(cfg)
	if err != nil {
		panic(err)
	}

	btcAuth := cfg.BtcAuth()

	return &Service{
		log: log,
		httpClient: &http.Transport{
			ReadBufferSize:        1024 * 1024,
			WriteBufferSize:       1024 * 1024,
			ResponseHeaderTimeout: time.Second * 30,
			MaxIdleConnsPerHost:   5000,
			MaxIdleConns:          5000,
		},
		FoundationDB:  fdbConnect,
		requestUrl:    btcAuth.Url,
		requestApiKey: btcAuth.Key,
	}
}
