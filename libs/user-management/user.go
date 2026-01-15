package usermgmt

import "time"

// User represents the core user entity with minimal required fields.
// Applications can extend this by embedding it in their own structs.
type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Never expose in JSON
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CreateUserRequest contains the data needed to create a new user.
type CreateUserRequest struct {
	Email    string
	Password string // Plain text - will be hashed
}

// UpdateUserRequest contains the data that can be updated for a user.
type UpdateUserRequest struct {
	Email    *string // Pointer allows partial updates
	Password *string // Pointer allows partial updates
}

// LoginRequest contains credentials for user authentication.
type LoginRequest struct {
	Email    string
	Password string
}

// LoginResponse contains the result of a successful login.
type LoginResponse struct {
	User  *User
	Token string
}
