package storage

import (
	"testing"

	_ "modernc.org/sqlite"
)

func TestNewStore(t *testing.T) {
	store := NewStore(":memory:")
	if store.db == nil {
		t.Fatal("database connection is nil")
	}
}

func TestMigrateCreatesDocumentsTable(t *testing.T) {
	store := NewStore(":memory:")

	_, err := store.db.Exec(`
		INSERT INTO documents (id, name, type) VALUES ('test-1', 'My Book', 'pdf')
	`)
	if err != nil {
		t.Fatal("failed to insert into documents table:", err)
	}

	var name string
	err = store.db.QueryRow("SELECT name FROM documents WHERE id = 'test-1'").Scan(&name)
	if err != nil {
		t.Fatal("failed to select from documents table:", err)
	}

	if name != "My Book" {
		t.Fatalf("expected 'My Book', got '%s'", name)
	}
}
