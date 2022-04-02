package db

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/prplecake/DiscordMusicCollector/dmc"
)

// A Store implements storage against a database
type Store struct {
	conn *sql.DB
}

// NewStore creates a database store
func NewStore(config dmc.DatabaseConfig) (*Store, error) {
	log.Print(config)
	if strings.ToLower(config.Type) == "sqlite" {
		Store, err := NewSqliteStore(config)
		return Store, err
	}
	if strings.ToLower(config.Type) == "postgres" {
		Store, err := NewPostgresStore(config)
		return Store, err
	}

	return nil, errors.New("no database")
}

// AddTrackToDB adds a track to the database
func AddTrackToDB(db *Store, track dmc.Track) error {
	if trackInDB(db, track) {
		return errors.New("track with service already exists in database")
	}

	statement, err := db.conn.Prepare("INSERT INTO tracks (title, artist, album, service) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(
		track.Title, strings.Join(track.Artists, ", "),
		track.Album, track.Service,
	)
	if err != nil {
		return err
	}
	return nil
}

func trackInDB(db *Store, track dmc.Track) bool {
	rows, err := db.conn.Query("SELECT COUNT(*) FROM tracks WHERE title = $1 AND service = $2",
		track.Title, track.Service)
	if err != nil {
		log.Fatal(err)
	}
	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}
	if count > 0 {
		return true
	}

	return false
}
