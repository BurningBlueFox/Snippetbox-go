package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"os"
)

type dbSettings struct {
	user     string
	password string
	host     string
	port     uint
}

func (dbSettings *dbSettings) getConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/snippetbox",
		dbSettings.user,
		dbSettings.password,
		dbSettings.host,
		dbSettings.port)
}

func getDbSettings() *dbSettings {
	value := new(dbSettings)
	flag.StringVar(&value.user, "db-user", "web", "User that shall be connected in the database")
	flag.StringVar(&value.password, "db-password", "pass", "Password for the database user")
	flag.StringVar(&value.host, "db-host", "localhost", "Host where the database is running")
	flag.UintVar(&value.port, "db-port", 5432, "Port where the database is listening")
	return value
}

func getDb(settings *dbSettings) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), settings.getConnectionString())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDbOrExit(settings *dbSettings, logger *slog.Logger) *pgxpool.Pool {
	db, err := getDb(settings)
	if err != nil {
		logger.Error("cannot connect to db", "error", err)
		os.Exit(1)
	}

	err = db.Ping(context.Background())
	if err != nil {
		logger.Error("db didn't respond to ping", "error", err)
		os.Exit(1)
	}

	logDbConnection(db, logger)
	return db
}

func logDbConnection(db *pgxpool.Pool, logger *slog.Logger) {
	logger.Info("established connection with db",
		"host", db.Config().ConnConfig.Host,
		"port", db.Config().ConnConfig.Port,
		"user", db.Config().ConnConfig.User,
		"db", db.Config().ConnConfig.Database,
	)
}
