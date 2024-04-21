package api

import (
	"log/slog"

	"github.com/maximotejeda/new-grpc-msvc/internal/ports"
)

type Application struct {
	db  ports.DBPort
	log *slog.Logger
}

func NewApplication(db ports.DBPort) *Application {
	log := slog.Default()
	log = log.With("adapter", "application")
	return &Application{
		db:  db,
		log: log,
	}
}

// todo Export differents Methods to map to grpc contracts
