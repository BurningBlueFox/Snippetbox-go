package main

import (
	"flag"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
	"os"
)

func getAddr() *string {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	return addr
}

func getLogger() *slog.Logger {
	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(loggerHandler)
	return logger
}

type application struct {
	logger *slog.Logger
	db     *pgxpool.Pool
}

func main() {
	logger := getLogger()
	addr := getAddr()
	db, err := getDb()
	defer db.Close()

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	} else {
		logDbConnection(db, logger)
	}

	app := &application{
		logger: logger,
		db:     db,
	}

	logger.Info("starting server", slog.String("addr", *addr))
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
