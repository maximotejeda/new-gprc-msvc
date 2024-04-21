package main

import (
	"log/slog"

	"github.com/maximotejeda/new-grpc-msvc/config"
	"github.com/maximotejeda/new-grpc-msvc/internal/adapters/db"
	"github.com/maximotejeda/new-grpc-msvc/internal/adapters/grpc"
	"github.com/maximotejeda/new-grpc-msvc/internal/application/core/api"
)

func main() {
	log := slog.Default()
	log.With("place", "main")
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Error("creating db connection", "error", err)
	}
	app := api.NewApplication(dbAdapter)

	grpcAdapter := grpc.NewAdapter(app, config.GetApplicationPort())
	log.Info("server running")
	grpcAdapter.Run()
}
