package db

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"
	"time"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schema string

type Adapter struct {
	db  *sql.DB
	log *slog.Logger
}

func NewAdapter(dataSourceURL string) (*Adapter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log := slog.Default().With("adapter", "db")
	db, err := sql.Open("sqlite", dataSourceURL)
	if err != nil {
		return nil, fmt.Errorf("connecion error: %w", err)
	}
	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping error: %w", err)
	}
	db.SetConnMaxIdleTime(10 * time.Second)
	CreateTables(db)
	// TODO create tables and trigers on first run
	return &Adapter{db: db, log: log}, nil
}

// Schema
func CreateTables(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		panic(err)
	}
}
