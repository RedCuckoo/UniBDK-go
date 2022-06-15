package server

import (
	"context"
	"github.com/RedCuckoo/UniBDK-go/proto/generated/proto"
	"github.com/RedCuckoo/UniBDK-go/unindexer/config"
	"github.com/RedCuckoo/UniBDK-go/unindexer/fdbutils"
	"gitlab.com/distributed_lab/logan/v3"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcServer struct {
	log          *logan.Entry
	FoundationDB *fdbutils.FDB
	addresses    config.GrpcAddressesConfig
	server       *grpc.Server
	proto.IndexerServer
}

func NewGrpcServer(cfg config.Config) *GrpcServer {
	log := cfg.Log().WithField("btc_indexer", "indexer")

	fdbConnect, err := fdbutils.Connect(cfg)
	if err != nil {
		panic(err)
	}

	return &GrpcServer{
		log:          log,
		FoundationDB: fdbConnect,
		addresses:    cfg.GrpcAddresses(),
	}
}

func (server *GrpcServer) Start() {
	listener, err := net.Listen("tcp", server.addresses.BtcIndexerGrpc)
	if err != nil {
		panic(err)
	}
	server.server = grpc.NewServer()

	proto.RegisterIndexerServer(server.server, server)

	go func() {
		server.log.Info("btcIndexer gRPC server is listening on address %s", server.addresses.BtcIndexerGrpc)
		if err = server.server.Serve(listener); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()
}

func (server *GrpcServer) AddressBalance(ctx context.Context, request *proto.AddressBalanceRequest) (*proto.AddressBalanceReply, error) {
	return server.FoundationDB.GetAddressBalance(ctx, request)
}
