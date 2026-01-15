package usermgmt

import "context"

// AuthProvider defines the interface for authentication mechanisms.
// This allows applications to choose their preferred auth strategy
// (JWT, sessions, etc.) by providing different implementations.
type AuthProvider interface {
	// GenerateToken creates an authentication token for the user
	GenerateToken(ctx context.Context, user *User) (string, error)

	// ValidateToken verifies a token and returns the associated user
	ValidateToken(ctx context.Context, token string) (*User, error)

	// RevokeToken invalidates a token (for logout)
	RevokeToken(ctx context.Context, token string) error
}

// PasswordHasher defines the interface for password hashing operations.
// This allows different hashing algorithms to be used.
type PasswordHasher interface {
	// Hash converts a plain text password into a secure hash
	Hash(password string) (string, error)

	// Compare checks if a plain text password matches a hash
	Compare(password, hash string) error
}
