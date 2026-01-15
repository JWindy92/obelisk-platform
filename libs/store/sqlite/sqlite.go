package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/JWindy92/obelisk-platform/libs/store"
)

// SQLiteStore implements the store.Store interface for SQLite databases.
type SQLiteStore struct {
	db       *sql.DB
	filepath string
	config   store.Config
}

// New creates a new SQLiteStore instance.
// The filepath parameter specifies the location of the SQLite database file.
// Use ":memory:" for an in-memory database.
func New(filepath string, config store.Config) *SQLiteStore {
	return &SQLiteStore{
		filepath: filepath,
		config:   config,
	}
}

// Connect establishes a connection to the SQLite database.
func (s *SQLiteStore) Connect(ctx context.Context) error {
	db, err := sql.Open("sqlite3", s.filepath)
	if err != nil {
		return fmt.Errorf("failed to open sqlite database: %w", err)
	}

	// Set connection pool settings
	if s.config.MaxOpenConns > 0 {
		db.SetMaxOpenConns(s.config.MaxOpenConns)
	}
	if s.config.MaxIdleConns > 0 {
		db.SetMaxIdleConns(s.config.MaxIdleConns)
	}
	if s.config.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(s.config.ConnMaxLifetime))
	}

	// Verify the connection works
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return fmt.Errorf("failed to ping sqlite database: %w", err)
	}

	s.db = db
	return nil
}

// Close gracefully closes the database connection.
func (s *SQLiteStore) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

// DB returns the underlying *sql.DB instance.
func (s *SQLiteStore) DB() *sql.DB {
	return s.db
}
