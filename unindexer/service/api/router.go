package api

import (
	"github.com/RedCuckoo/UniBDK-go/unindexer/config"
	"github.com/deliveroo/jsonrpc-go"

	//"github.com/RedCuckoo/UniBDK-go/unindexer/internal/data/pg"
	"github.com/RedCuckoo/UniBDK-go/unindexer/service/api/helpers"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewRouter(cfg config.Config) http.Handler {
	server := jsonrpc.New()

	server.Use(
		LoggingMiddleware(cfg.Log().WithField("json-rpc requests", "handler").Formatter(&logrus.JSONFormatter{})),
		CtxMiddleware(
			helpers.CtxLog(cfg.Log().WithField("json-rpc requests", "handler").Formatter(&logrus.JSONFormatter{})),
			//helpers.CtxDepositsDB(pg.NewDepositQ(cfg.DB())),
			//helpers.CtxAccount(cfg.Account()),
		),
	)

	server.Register(jsonrpc.Methods{
		//"getbalance":     handlers.GetBalance,
		//"gettransaction": handlers.GetTransaction,
		//"sendtoaddress":  handlers.SendToAddress,
		//"listsinceblock": handlers.ListSinceBlock,
	})

	//authInfo := cfg.BasicAuth()

	return NewBasicAuthRpc(server, "jsonrpc", map[string]string{
		//authInfo.Login: authInfo.Password,
	})
}
