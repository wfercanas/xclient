package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	Logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Address")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		Logger: logger,
	}

	app.Logger.Info("starting server", slog.String("addr", *addr))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)

	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		app.Logger.Error(err.Error())
		os.Exit(1)
	}
}
