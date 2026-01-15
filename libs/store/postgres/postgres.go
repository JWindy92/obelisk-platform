package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/JWindy92/obelisk-platform/libs/store"
)

// PostgresStore implements the store.Store interface for PostgreSQL databases.
type PostgresStore struct {
	db     *sql.DB
	dsn    string
	config store.Config
}

// Config holds PostgreSQL-specific configuration.
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string // disable, require, verify-ca, verify-full
}

// New creates a new PostgresStore instance.
// The config parameter contains PostgreSQL connection details.
func New(cfg Config, storeConfig store.Config) *PostgresStore {
	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	return &PostgresStore{
		dsn:    dsn,
		config: storeConfig,
	}
}

// Connect establishes a connection to the PostgreSQL database.
func (p *PostgresStore) Connect(ctx context.Context) error {
	db, err := sql.Open("postgres", p.dsn)
	if err != nil {
		return fmt.Errorf("failed to open postgres database: %w", err)
	}

	// Set connection pool settings
	if p.config.MaxOpenConns > 0 {
		db.SetMaxOpenConns(p.config.MaxOpenConns)
	}
	if p.config.MaxIdleConns > 0 {
		db.SetMaxIdleConns(p.config.MaxIdleConns)
	}
	if p.config.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(p.config.ConnMaxLifetime))
	}

	// Verify the connection works
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return fmt.Errorf("failed to ping postgres database: %w", err)
	}

	p.db = db
	return nil
}

// Close gracefully closes the database connection.
func (p *PostgresStore) Close() error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

// DB returns the underlying *sql.DB instance.
func (p *PostgresStore) DB() *sql.DB {
	return p.db
}
