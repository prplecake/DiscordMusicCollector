package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/prplecake/DiscordMusicCollector/dmc"

	_ "modernc.org/sqlite" // sqlite driver
)

// NewSqliteStore opens a SQLite
func NewSqliteStore(config dmc.DatabaseConfig) (*Store, error) {
	dbf := "DMC.db"
	if _, err := os.Stat(dbf); os.IsNotExist(err) {
		dbFile, err := os.Create(dbf)
		if err != nil {
			return nil, err
		}
		dbFile.Close()
	}

	db, err := sql.Open("sqlite", dbf)
	if err != nil {
		return nil, err
	}

	initTableSQL := `CREATE TABLE IF NOT EXISTS tracks (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"Artist" TEXT,
		"Album" TEXT,
		"Service" TEXT
	);`

	log.Print("Initializing sqlite database...")
	statement, err := db.Prepare(initTableSQL)
	if err != nil {
		return nil, err
	}
	statement.Exec()
	log.Print(db)
	return &Store{conn: db}, nil
}
