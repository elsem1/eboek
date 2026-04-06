package storage

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
)

func NewStore(path string) *Store {
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		path = filepath.Join(home, ".local", "share", "eboek", "eboek.db")
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatal(err)
	}

	store := &Store{db: db}
	store.migrate()

	return store
}

func (s *Store) migrate() {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS documents (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			type TEXT NOT NULL,
			parent_id TEXT,
			checksum TEXT,
			sync_status TEXT NOT NULL DEFAULT 'pending_download',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
			)
		`)
	if err != nil {
		log.Fatal(err)
	}
}

type Store struct {
	db *sql.DB
}
