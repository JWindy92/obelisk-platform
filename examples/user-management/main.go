package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JWindy92/obelisk-platform/libs/store"
	"github.com/JWindy92/obelisk-platform/libs/store/sqlite"
	usermgmt "github.com/JWindy92/obelisk-platform/libs/user-management"
)

func main() {
	ctx := context.Background()

	// Step 1: Initialize the database store
	// This can be SQLite, Postgres, or any other store.Store implementation
	dbStore := initStore(ctx)
	defer dbStore.Close()

	// Step 2: Create the user repository
	// The repository accepts the store.Store interface
	config := usermgmt.DefaultConfig()
	repo := usermgmt.NewRepository(dbStore, config)

	// Step 3: Create auth provider and password hasher
	// In a real app, you'd use actual implementations (JWT, bcrypt, etc.)
	authProvider := &mockAuthProvider{}
	passwordHasher := &mockPasswordHasher{}

	// Step 4: Create the user service
	// The service accepts all dependencies via interfaces
	userService := usermgmt.NewService(repo, authProvider, passwordHasher, config)

	// Step 5: Use the service
	demonstrateUsage(ctx, userService)
}

// initStore creates and connects a database store.
// In a real app, you'd choose SQLite, Postgres, etc. based on your needs.
func initStore(ctx context.Context) store.Store {
	storeConfig := store.Config{
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: int64(time.Hour),
	}

	st := sqlite.New(":memory:", storeConfig)
	if err := st.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("✓ Database connected")
	return st
}

// demonstrateUsage shows basic operations with the user service.
func demonstrateUsage(ctx context.Context, svc usermgmt.Service) {
	fmt.Println("\n=== User Management Example ===\n")

	// Example: Signup
	fmt.Println("1. Creating a new user...")
	// user, err := svc.Signup(ctx, usermgmt.CreateUserRequest{
	// 	Email:    "user@example.com",
	// 	Password: "securepassword123",
	// })
	// if err != nil {
	// 	log.Printf("Signup failed: %v", err)
	// } else {
	// 	fmt.Printf("✓ User created: %s\n", user.Email)
	// }

	// Example: Login
	fmt.Println("\n2. Logging in...")
	// loginResp, err := svc.Login(ctx, usermgmt.LoginRequest{
	// 	Email:    "user@example.com",
	// 	Password: "securepassword123",
	// })
	// if err != nil {
	// 	log.Printf("Login failed: %v", err)
	// } else {
	// 	fmt.Printf("✓ Login successful. Token: %s\n", loginResp.Token)
	// }

	// Example: Get user
	fmt.Println("\n3. Getting user by ID...")
	// user, err = svc.GetUser(ctx, user.ID)
	// if err != nil {
	// 	log.Printf("Get user failed: %v", err)
	// } else {
	// 	fmt.Printf("✓ Found user: %s\n", user.Email)
	// }

	fmt.Println("\n(Methods not yet implemented - coming soon!)")
	fmt.Println("\n=== Key Points ===")
	fmt.Println("• Service accepts Repository, AuthProvider, and PasswordHasher interfaces")
	fmt.Println("• Repository accepts store.Store interface")
	fmt.Println("• Easy to swap implementations (SQLite → Postgres, JWT → Sessions, etc.)")
	fmt.Println("• All dependencies injected at initialization")
}

// mockAuthProvider is a simple mock for demonstration.
// In production, implement with JWT, sessions, etc.
type mockAuthProvider struct{}

func (m *mockAuthProvider) GenerateToken(ctx context.Context, user *usermgmt.User) (string, error) {
	return "mock-token-" + user.ID, nil
}

func (m *mockAuthProvider) ValidateToken(ctx context.Context, token string) (*usermgmt.User, error) {
	return &usermgmt.User{ID: "123", Email: "user@example.com"}, nil
}

func (m *mockAuthProvider) RevokeToken(ctx context.Context, token string) error {
	return nil
}

// mockPasswordHasher is a simple mock for demonstration.
// In production, implement with bcrypt, argon2, etc.
type mockPasswordHasher struct{}

func (m *mockPasswordHasher) Hash(password string) (string, error) {
	return "hashed-" + password, nil
}

func (m *mockPasswordHasher) Compare(password, hash string) error {
	if "hashed-"+password == hash {
		return nil
	}
	return fmt.Errorf("password mismatch")
}
