package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/JWindy92/obelisk-platform/libs/store"
)

func TestPostgresStore_Connect(t *testing.T) {
	tests := []struct {
		name      string
		config    Config
		storeConf store.Config
		wantErr   bool
	}{
		{
			name: "successful connection with valid config",
			config: Config{
				Host:     "localhost",
				Port:     5432,
				User:     "obelisk",
				Password: "obelisk123",
				DBName:   "obelisk_dev",
				SSLMode:  "disable",
			},
			storeConf: store.Config{
				MaxOpenConns:    10,
				MaxIdleConns:    5,
				ConnMaxLifetime: int64(time.Hour),
			},
			wantErr: false,
		},
		{
			name: "connection failure with invalid host",
			config: Config{
				Host:     "invalid-host-that-does-not-exist",
				Port:     5432,
				User:     "obelisk",
				Password: "obelisk123",
				DBName:   "obelisk_dev",
				SSLMode:  "disable",
			},
			storeConf: store.Config{},
			wantErr:   true,
		},
		{
			name: "connection failure with wrong credentials",
			config: Config{
				Host:     "localhost",
				Port:     5432,
				User:     "wrong_user",
				Password: "wrong_password",
				DBName:   "obelisk_dev",
				SSLMode:  "disable",
			},
			storeConf: store.Config{},
			wantErr:   true,
		},
		{
			name: "connection failure with invalid database name",
			config: Config{
				Host:     "localhost",
				Port:     5432,
				User:     "obelisk",
				Password: "obelisk123",
				DBName:   "nonexistent_db",
				SSLMode:  "disable",
			},
			storeConf: store.Config{},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := New(tt.config, tt.storeConf)

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

			// Clean up
			if err := st.Close(); err != nil {
				t.Errorf("Close() unexpected error: %v", err)
			}
		})
	}
}

func TestPostgresStore_Close(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "obelisk",
		Password: "obelisk123",
		DBName:   "obelisk_dev",
		SSLMode:  "disable",
	}

	st := New(config, store.Config{})

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

func TestPostgresStore_DB(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "obelisk",
		Password: "obelisk123",
		DBName:   "obelisk_dev",
		SSLMode:  "disable",
	}

	st := New(config, store.Config{})

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

func TestPostgresStore_ConnectionPoolSettings(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "obelisk",
		Password: "obelisk123",
		DBName:   "obelisk_dev",
		SSLMode:  "disable",
	}

	storeConf := store.Config{
		MaxOpenConns:    20,
		MaxIdleConns:    10,
		ConnMaxLifetime: int64(30 * time.Minute),
	}

	st := New(config, storeConf)

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
