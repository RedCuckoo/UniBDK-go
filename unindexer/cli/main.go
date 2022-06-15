package cli

import (
	"context"
	"github.com/RedCuckoo/UniBDK-go/unindexer/config"
	"github.com/RedCuckoo/UniBDK-go/unindexer/server"
	"github.com/RedCuckoo/UniBDK-go/unindexer/service/indexers/btc_indexer"
	"github.com/RedCuckoo/UniBDK-go/unindexer/service/indexers/eth_indexer"
	"github.com/alecthomas/kingpin"
	"github.com/sirupsen/logrus"

	//"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

func Run(args []string) bool {
	log := logan.New().Formatter(&logrus.JSONFormatter{})

	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("app panicked")
		}
	}()

	cfg := config.New(kv.MustFromEnv())
	log = cfg.Log().Formatter(&logrus.JSONFormatter{})

	app := kingpin.New("unindexer", "")

	runCmd := app.Command("run", "run command")
	serviceCmd := runCmd.Command("indexer", "run indexer") // you can insert custom help

	btcIndexer := serviceCmd.Command("btc", "run indexer btc")
	ethIndexer := serviceCmd.Command("eth", "run indexer eth")

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	grpcServer := server.NewGrpcServer(cfg)
	grpcServer.Start()

	switch cmd {
	case btcIndexer.FullCommand():
		svc := btc_indexer.New(cfg)
		svc.Run(context.Background())
	case ethIndexer.FullCommand():
		svc := eth_indexer.New(cfg)
		svc.Run(context.Background())
	default:
		log.Errorf("unknown command %s", cmd)
		return false
	}

	if err != nil {
		log.WithError(err).Error("failed to exec cmd")
		return false
	}

	return true
}
