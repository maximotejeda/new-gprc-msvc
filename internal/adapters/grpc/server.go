package grpc

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/maximotejeda/new-grpc-msvc/config"
	"github.com/maximotejeda/new-grpc-msvc/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	log  *slog.Logger
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	log := slog.Default()
	log = log.With("adapter", "dolar-grpc")
	return &Adapter{api: api, port: port, log: log}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		a.log.Error("failed to listen", "port", a.port, "error", err)
		panic(err)
	}
	grpcServer := grpc.NewServer()

	// TODO: register service here
	// Service.RegisterDollarServer(grpcServer, a)

	if config.GetEnv() == "development" {
		// This is to use GRPCurl to test services
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listen); err != nil {
		a.log.Error("failed to serve grpc on port", "port", a.port)
	}
}
