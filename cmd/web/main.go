package main

import (
	"flag"
	"github.com/BurningBlueFox/letsgo/internal/models"
	"log/slog"
	"net/http"
	"os"
)

func getAddr() *string {
	value := flag.String("addr", ":4000", "HTTP network address")
	return value
}

func getLogger() *slog.Logger {
	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(loggerHandler)
	return logger
}

type application struct {
	logger        *slog.Logger
	snippetsModel *models.SnippetModel
}

func main() {
	logger := getLogger()
	addr := getAddr()
	dbSetting := getDbSettings()
	flag.Parse()

	db := getDbOrExit(dbSetting, logger)
	defer db.Close()

	app := &application{
		logger:        logger,
		snippetsModel: &models.SnippetModel{DB: db},
	}

	logger.Info("starting server", slog.String("addr", *addr))
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
