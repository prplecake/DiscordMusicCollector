package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/prplecake/DiscordMusicCollector/dmc"

	_ "github.com/lib/pq" // Postgres driver
)

// NewPostgresStore creates storage against a Postgres database
func NewPostgresStore(config dmc.DatabaseConfig) (*Store, error) {
	if err := validateConfig(config); err != nil {
		log.Fatal("database config error: ", err)
	}

	dbinfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		config.Username, config.Password, config.Host, config.Name, config.SSLMode)
	log.Print("dbinfo: ", dbinfo)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}

	var initTableSQL = `CREATE TABLE IF NOT EXISTS tracks (
		"id" integer NOT NULL PRIMARY KEY,
		"title" TEXT,
		"artist" TEXT,
		"album" TEXT,
		"service" TEXT
	);`

	log.Print("Initializing postgres database...")
	statement, err := db.Prepare(initTableSQL)
	if err != nil {
		return nil, err
	}
	statement.Exec()
	log.Print(db)
	return &Store{conn: db}, nil
}

func validateConfig(config dmc.DatabaseConfig) error {
	if config.Username == "" {
		return errors.New("database username not set")
	}
	if config.Password == "" {
		return errors.New("database password not set")
	}
	if config.Name == "" {
		return errors.New("database name not set")
	}
	return nil
}
