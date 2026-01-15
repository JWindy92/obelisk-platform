package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JWindy92/obelisk-platform/libs/store"
	"github.com/JWindy92/obelisk-platform/libs/store/postgres"
	"github.com/JWindy92/obelisk-platform/libs/store/sqlite"
)

func main() {
	ctx := context.Background()

	// Example 1: Using SQLite
	fmt.Println("=== SQLite Example ===")
	sqliteStore := initSQLiteStore()
	if err := sqliteStore.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer sqliteStore.Close()

	// Pass to application - accepts Store interface
	runApp(sqliteStore)

	fmt.Println()

	// Example 2: Using Postgres
	fmt.Println("=== Postgres Example ===")
	postgresStore := initPostgresStore()
	if err := postgresStore.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer postgresStore.Close()

	// Pass to same application - works with any Store implementation!
	runApp(postgresStore)
}

// initSQLiteStore creates and returns a SQLite store instance.
// To switch to SQLite, just call this function instead of initPostgresStore.
func initSQLiteStore() *sqlite.SQLiteStore {
	config := store.Config{
		MaxOpenConns:    25,
		MaxIdleConns:    5,
		ConnMaxLifetime: int64(time.Hour),
	}

	return sqlite.New("./example.db", config)
}

// initPostgresStore creates and returns a Postgres store instance.
// To switch to Postgres, just call this function instead of initSQLiteStore.
func initPostgresStore() *postgres.PostgresStore {
	pgConfig := postgres.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "obelisk",
		Password: "obelisk123",
		DBName:   "obelisk_dev",
		SSLMode:  "disable",
	}

	storeConfig := store.Config{
		MaxOpenConns:    25,
		MaxIdleConns:    5,
		ConnMaxLifetime: int64(time.Hour),
	}

	return postgres.New(pgConfig, storeConfig)
}

// runApp demonstrates the dependency injection pattern.
// It accepts the Store INTERFACE, not a concrete type.
// This means it works with ANY implementation (SQLite, Postgres, etc.)
// without needing to change the application code.
func runApp(s store.Store) {
	ctx := context.Background()

	// Test the connection
	if err := s.DB().PingContext(ctx); err != nil {
		log.Printf("Failed to ping database: %v", err)
		return
	}

	fmt.Println("✓ Successfully connected to database")
	fmt.Printf("✓ Connection is healthy\n")

	// Your application logic here...
	// The code doesn't know or care whether it's using SQLite or Postgres!

	// Example: Get connection stats
	stats := s.DB().Stats()
	fmt.Printf("✓ Max open connections: %d\n", stats.MaxOpenConnections)
	fmt.Printf("✓ Open connections: %d\n", stats.OpenConnections)
}
