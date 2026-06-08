package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type application struct {
	Logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Address")
	dsn := flag.String("dsn", "user=iclient password=iclient host=localhost port=5432 dbname=iclient sslmode=disable", "Data Source Name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		Logger: logger,
	}

	app.Logger.Info("starting server", slog.String("addr", *addr))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)

	err = http.ListenAndServe(*addr, mux)
	if err != nil {
		app.Logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
