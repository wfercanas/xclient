package main

import (
	"database/sql"
	"fmt"
)

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
