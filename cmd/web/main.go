package main

import (
	"flag"
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
}

func main() {
	logger := getLogger()
	addr := getAddr()

	app := &application{
		logger: logger,
	}

	logger.Info("starting server", slog.String("addr", *addr))
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
