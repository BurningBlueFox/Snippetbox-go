package main

import (
	"flag"
	"github.com/BurningBlueFox/letsgo/internal/models"
	"html/template"
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
	templateCache map[string]*template.Template
}

func main() {
	logger := getLogger()
	addr := getAddr()
	dbSetting := getDbSettings()
	flag.Parse()

	db := getDbOrExit(dbSetting, logger)
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	app := &application{
		logger:        logger,
		snippetsModel: &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	logger.Info("starting server", slog.String("addr", *addr))
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
