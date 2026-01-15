# Store Library

A database-agnostic store library that provides a common interface for different database implementations.

## Design Philosophy

This library follows the "Accept Interfaces, Return Structs" pattern, allowing applications to remain database-agnostic. You can switch between database implementations (e.g., SQLite to PostgreSQL) by simply changing the concrete type passed to your application.

## Usage

### SQLite Example

```go
package main

import (
    "context"
    "log"
    "time"
    
    "github.com/JWindy92/obelisk-platform/libs/store"
    "github.com/JWindy92/obelisk-platform/libs/store/sqlite"
)

func main() {
    // Create SQLite store
    config := store.Config{
        MaxOpenConns: 25,
        MaxIdleConns: 5,
        ConnMaxLifetime: int64(time.Hour),
    }
    
    st := sqlite.New("./data.db", config)
    
    // Connect to database
    ctx := context.Background()
    if err := st.Connect(ctx); err != nil {
        log.Fatal(err)
    }
    defer st.Close()
    
    // Pass to your application
    runApp(st)
}

// Your application accepts the interface
func runApp(st store.Store) {
    // Use the store
}
```

### PostgreSQL Example

```go
package main

import (
    "context"
    "log"
    "time"
    
    "github.com/JWindy92/obelisk-platform/libs/store"
    "github.com/JWindy92/obelisk-platform/libs/store/postgres"
)

func main() {
    // Create Postgres store
    pgConfig := postgres.Config{
        Host:     "localhost",
        Port:     5432,
        User:     "myuser",
        Password: "mypassword",
        DBName:   "mydb",
        SSLMode:  "disable",
    }
    
    storeConfig := store.Config{
        MaxOpenConns: 25,
        MaxIdleConns: 5,
        ConnMaxLifetime: int64(time.Hour),
    }
    
    st := postgres.New(pgConfig, storeConfig)
    
    // Connect to database
    ctx := context.Background()
    if err := st.Connect(ctx); err != nil {
        log.Fatal(err)
    }
    defer st.Close()
    
    // Pass to your application - same interface!
    runApp(st)
}

// Your application accepts the interface
func runApp(st store.Store) {
    // Use the store - works with any implementation
}
```

## Switching Implementations

To switch from SQLite to PostgreSQL (or vice versa), you only need to change the initialization code in your `main()` function. Your application code remains unchanged.

## Current Status

✅ Store interface defined  
✅ Connect() method implemented  
⏳ CRUD operations (coming next)  
⏳ Transaction support (coming next)  
⏳ Migration support (coming next)
