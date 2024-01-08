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

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	logger.Info("starting server", slog.String("addr", *addr))
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
