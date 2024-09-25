// database/database.go
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// NewDatabase initializes a new in-memory SQLite database.
func NewDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("failed to open in-memory database: %w", err)
	}

	// Ensure the todos table exists
	if err := createTable(db); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// createTable creates the todos table.
func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	)`)
	if err != nil {
		return fmt.Errorf("failed to create todos table: %w", err)
	}
	return nil
}
