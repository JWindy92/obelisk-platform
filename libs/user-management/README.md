# User Management Library

A reusable, database-agnostic user management framework that provides common user operations across applications.

## Design Philosophy

This library follows the same dependency injection pattern as the store library, using "Accept Interfaces, Return Structs". This allows applications to:

- Share a common User model across services
- Extend the base User model via composition
- Choose their preferred auth strategy (JWT, sessions, etc.)
- Remain database-agnostic by using the store.Store interface

## Architecture

```
┌─────────────────┐
│   Application   │
└────────┬────────┘
         │
    ┌────▼─────┐
    │ Service  │ - Business logic (signup, login, etc.)
    └────┬─────┘
         │
    ┌────▼────────┐
    │ Repository  │ - Data access (CRUD operations)
    └────┬────────┘
         │
    ┌────▼────────┐
    │ store.Store │ - Database abstraction
    └─────────────┘
```

## Components

### User Model
Core entity with email and password. Applications can extend via composition.

### Repository
Handles database operations (CRUD). Accepts `store.Store` interface.

### Service
Business logic layer (signup, login, validation). Accepts `Repository` and `AuthProvider`.

### AuthProvider (Interface)
Pluggable authentication strategy. Implement your own JWT, session, or custom auth.

### PasswordHasher (Interface)
Pluggable password hashing. Implement with bcrypt, argon2, etc.

## Usage (Planned)

```go
// Initialize dependencies
dbStore := postgres.New(pgConfig, storeConfig)
dbStore.Connect(ctx)

repo := usermgmt.NewRepository(dbStore, usermgmt.DefaultConfig())
authProvider := jwt.NewProvider(jwtSecret)  // Your implementation
passwordHasher := bcrypt.NewHasher()        // Your implementation

svc := usermgmt.NewService(repo, authProvider, passwordHasher, usermgmt.DefaultConfig())

// Use the service
user, err := svc.Signup(ctx, usermgmt.CreateUserRequest{
    Email:    "user@example.com",
    Password: "secretpassword",
})
```

## Extension Example

```go
// Extend the base User model
type AppUser struct {
    usermgmt.User        // Embed base model
    Role          string // App-specific field
    Department    string // App-specific field
}
```

## Current Status

✅ Core data structures defined  
✅ Repository interface and structure  
✅ Service interface and structure  
✅ Pluggable auth provider interface  
✅ Pluggable password hasher interface  
⏳ Implementation of methods (coming next)  
⏳ JWT auth provider implementation  
⏳ Bcrypt password hasher implementation  
⏳ Database migrations
