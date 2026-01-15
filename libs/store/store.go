package store

import (
	"context"
	"database/sql"
)

// Store defines the interface that all database implementations must satisfy.
// This allows applications to remain database-agnostic by programming against
// the interface rather than concrete implementations.
type Store interface {
	// Connect establishes a connection to the database.
	// It should handle connection pooling and validation internally.
	Connect(ctx context.Context) error

	// Close gracefully closes the database connection and cleans up resources.
	Close() error

	// DB returns the underlying *sql.DB for cases where direct access is needed.
	// Use sparingly - prefer adding methods to the Store interface instead.
	DB() *sql.DB
}

// Config holds common configuration options for database connections.
type Config struct {
	// MaxOpenConns sets the maximum number of open connections to the database.
	MaxOpenConns int

	// MaxIdleConns sets the maximum number of idle connections.
	MaxIdleConns int

	// ConnMaxLifetime sets the maximum time a connection can be reused.
	// Use time.Duration (e.g., time.Hour)
	ConnMaxLifetime int64
}
