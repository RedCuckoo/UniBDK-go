package api

import (
	"github.com/RedCuckoo/UniBDK-go/unindexer/config"
	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
)

type service struct {
	log *logan.Entry
	cfg config.Config
}

func (s *service) run() error {
	r := NewRouter(s.cfg)

	return http.Serve(s.cfg.Listener(), r)
}

func newService(cfg config.Config) *service {
	return &service{
		log: cfg.Log().WithField("json-rpc requests", "handler").Formatter(&logrus.JSONFormatter{}),
		cfg: cfg,
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
