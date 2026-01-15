package sqlite

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/JWindy92/obelisk-platform/libs/store"
)

func TestSQLiteStore_Connect(t *testing.T) {
	tests := []struct {
		name      string
		filepath  string
		storeConf store.Config
		setup     func() error
		cleanup   func() error
		wantErr   bool
	}{
		{
			name:     "successful connection with file database",
			filepath: filepath.Join(t.TempDir(), "test.db"),
			storeConf: store.Config{
				MaxOpenConns:    10,
				MaxIdleConns:    5,
				ConnMaxLifetime: int64(time.Hour),
			},
			wantErr: false,
		},
		{
			name:      "successful connection with in-memory database",
			filepath:  ":memory:",
			storeConf: store.Config{},
			wantErr:   false,
		},
		{
			name:     "successful connection creates directory if needed",
			filepath: filepath.Join(t.TempDir(), "subdir", "test.db"),
			storeConf: store.Config{
				MaxOpenConns: 5,
			},
			wantErr: false,
		},
		{
			name:     "connection failure with invalid path",
			filepath: "/invalid/path/that/cannot/be/created/\x00/test.db",
			storeConf: store.Config{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				if err := tt.setup(); err != nil {
					t.Fatalf("Setup failed: %v", err)
				}
			}
			if tt.cleanup != nil {
				defer func() {
					if err := tt.cleanup(); err != nil {
						t.Errorf("Cleanup failed: %v", err)
					}
				}()
			}

			st := New(tt.filepath, tt.storeConf)
			
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			
			err := st.Connect(ctx)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("Connect() expected error but got nil")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Connect() unexpected error: %v", err)
				return
			}
			
			// Verify the connection is actually established
			if st.DB() == nil {
				t.Error("Connect() succeeded but DB() returned nil")
			}
			
			// Test that we can ping the database
			if err := st.DB().PingContext(ctx); err != nil {
				t.Errorf("DB ping failed after Connect(): %v", err)
			}
			
			// Verify we can execute a simple query
			var result int
			err = st.DB().QueryRowContext(ctx, "SELECT 1").Scan(&result)
			if err != nil {
				t.Errorf("Simple query failed: %v", err)
			}
			if result != 1 {
				t.Errorf("Query returned %d, want 1", result)
			}
			
			// Clean up
			if err := st.Close(); err != nil {
				t.Errorf("Close() unexpected error: %v", err)
			}
		})
	}
}

func TestSQLiteStore_Close(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "test.db")
	st := New(dbPath, store.Config{})
	
	ctx := context.Background()
	if err := st.Connect(ctx); err != nil {
		t.Fatalf("Connect() failed: %v", err)
	}
	
	// Close should work
	if err := st.Close(); err != nil {
		t.Errorf("Close() unexpected error: %v", err)
	}
	
	// Multiple closes should be safe
	if err := st.Close(); err != nil {
		t.Errorf("Second Close() unexpected error: %v", err)
	}
}

func TestSQLiteStore_DB(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "test.db")
	st := New(dbPath, store.Config{})
	
	// DB should be nil before Connect
	if st.DB() != nil {
		t.Error("DB() should return nil before Connect()")
	}
	
	ctx := context.Background()
	if err := st.Connect(ctx); err != nil {
		t.Fatalf("Connect() failed: %v", err)
	}
	defer st.Close()
	
	// DB should be non-nil after Connect
	if st.DB() == nil {
		t.Error("DB() should return non-nil after Connect()")
	}
}

func TestSQLiteStore_ConnectionPoolSettings(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "test.db")
	storeConf := store.Config{
		MaxOpenConns:    20,
		MaxIdleConns:    10,
		ConnMaxLifetime: int64(30 * time.Minute),
	}
	
	st := New(dbPath, storeConf)
	
	ctx := context.Background()
	if err := st.Connect(ctx); err != nil {
		t.Fatalf("Connect() failed: %v", err)
	}
	defer st.Close()
	
	// Verify connection pool settings are applied
	stats := st.DB().Stats()
	if stats.MaxOpenConnections != 20 {
		t.Errorf("MaxOpenConnections = %d, want 20", stats.MaxOpenConnections)
	}
}

func TestSQLiteStore_InMemory(t *testing.T) {
	st := New(":memory:", store.Config{})
	
	ctx := context.Background()
	if err := st.Connect(ctx); err != nil {
		t.Fatalf("Connect() failed: %v", err)
	}
	defer st.Close()
	
	// Create a table and insert data
	_, err := st.DB().ExecContext(ctx, `
		CREATE TABLE test (
			id INTEGER PRIMARY KEY,
			name TEXT
		)
	`)
	if err != nil {
		t.Fatalf("Create table failed: %v", err)
	}
	
	_, err = st.DB().ExecContext(ctx, "INSERT INTO test (name) VALUES (?)", "test_value")
	if err != nil {
		t.Fatalf("Insert failed: %v", err)
	}
	
	// Verify data was inserted
	var name string
	err = st.DB().QueryRowContext(ctx, "SELECT name FROM test WHERE id = 1").Scan(&name)
	if err != nil {
		t.Fatalf("Select failed: %v", err)
	}
	
	if name != "test_value" {
		t.Errorf("Got name %q, want %q", name, "test_value")
	}
}

func TestSQLiteStore_FilePersistence(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "persistent.db")
	
	// Create and populate database
	{
		st := New(dbPath, store.Config{})
		ctx := context.Background()
		
		if err := st.Connect(ctx); err != nil {
			t.Fatalf("Connect() failed: %v", err)
		}
		
		_, err := st.DB().ExecContext(ctx, `
			CREATE TABLE test (
				id INTEGER PRIMARY KEY,
				value TEXT
			)
		`)
		if err != nil {
			t.Fatalf("Create table failed: %v", err)
		}
		
		_, err = st.DB().ExecContext(ctx, "INSERT INTO test (value) VALUES (?)", "persistent_data")
		if err != nil {
			t.Fatalf("Insert failed: %v", err)
		}
		
		st.Close()
	}
	
	// Verify file exists
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		t.Fatalf("Database file was not created at %s", dbPath)
	}
	
	// Reopen and verify data persisted
	{
		st := New(dbPath, store.Config{})
		ctx := context.Background()
		
		if err := st.Connect(ctx); err != nil {
			t.Fatalf("Second Connect() failed: %v", err)
		}
		defer st.Close()
		
		var value string
		err := st.DB().QueryRowContext(ctx, "SELECT value FROM test WHERE id = 1").Scan(&value)
		if err != nil {
			t.Fatalf("Select after reopen failed: %v", err)
		}
		
		if value != "persistent_data" {
			t.Errorf("Got value %q, want %q", value, "persistent_data")
		}
	}
}