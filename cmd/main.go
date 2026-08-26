package main

import (
	"flag"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"github.com/selsa-inube/iclient-query-service/config"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP Address")
	dsn := flag.String("dsn", "user=fernando password=Sparkie11 host=localhost port=5432 dbname=xclient sslmode=disable", "Data Source Name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := config.NewApplication(*logger, db)
	server := config.NewServer(*addr, router(&app))

	app.Logger.Info("starting server", slog.String("addr", server.Addr))

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Error(err.Error())
		os.Exit(1)
	}
}
