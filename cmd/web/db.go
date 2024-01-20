package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

func getDb() (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), "postgresql://web:pass@localhost:5432/snippetbox")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func logDbConnection(db *pgxpool.Pool, logger *slog.Logger) {
	logger.Info("established connection with db",
		"host", db.Config().ConnConfig.Host,
		"port", db.Config().ConnConfig.Port,
		"user", db.Config().ConnConfig.User,
		"db", db.Config().ConnConfig.Database,
	)
}
