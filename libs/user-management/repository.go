package usermgmt

import (
	"context"

	"github.com/JWindy92/obelisk-platform/libs/store"
)

// Repository defines the data access interface for user operations.
// This interface abstracts the database operations, allowing the business
// logic to remain independent of the underlying storage implementation.
type Repository interface {
	// Create inserts a new user into the database
	Create(ctx context.Context, user *User) error

	// GetByID retrieves a user by their unique identifier
	GetByID(ctx context.Context, id string) (*User, error)

	// GetByEmail retrieves a user by their email address
	GetByEmail(ctx context.Context, email string) (*User, error)

	// Update modifies an existing user's data
	Update(ctx context.Context, user *User) error

	// Delete removes a user from the database
	Delete(ctx context.Context, id string) error

	// List retrieves all users with optional pagination
	List(ctx context.Context, limit, offset int) ([]*User, error)
}

// repository is the concrete implementation of Repository.
// It uses the store.Store interface to remain database-agnostic.
type repository struct {
	store     store.Store
	tableName string
}

// NewRepository creates a new Repository instance.
// It follows the "Accept Interfaces, Return Structs" pattern by accepting
// the store.Store interface, making it work with any database implementation.
func NewRepository(st store.Store, config Config) *repository {
	tableName := config.TableName
	if tableName == "" {
		tableName = "users"
	}

	return &repository{
		store:     st,
		tableName: tableName,
	}
}

// Create inserts a new user into the database
func (r *repository) Create(ctx context.Context, user *User) error {
	// TODO: Implement user creation with SQL INSERT
	return nil
}

// GetByID retrieves a user by their unique identifier
func (r *repository) GetByID(ctx context.Context, id string) (*User, error) {
	// TODO: Implement user retrieval with SQL SELECT WHERE id = ?
	return nil, nil
}

// GetByEmail retrieves a user by their email address
func (r *repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	// TODO: Implement user retrieval with SQL SELECT WHERE email = ?
	return nil, nil
}

// Update modifies an existing user's data
func (r *repository) Update(ctx context.Context, user *User) error {
	// TODO: Implement user update with SQL UPDATE
	return nil
}

// Delete removes a user from the database
func (r *repository) Delete(ctx context.Context, id string) error {
	// TODO: Implement user deletion with SQL DELETE
	return nil
}

// List retrieves all users with optional pagination
func (r *repository) List(ctx context.Context, limit, offset int) ([]*User, error) {
	// TODO: Implement user listing with SQL SELECT with LIMIT and OFFSET
	return nil, nil
}
