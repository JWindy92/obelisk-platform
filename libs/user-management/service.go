package usermgmt

import "context"

// Service defines the business logic interface for user management.
// It orchestrates operations between the repository and auth provider.
type Service interface {
	// Signup creates a new user account
	Signup(ctx context.Context, req CreateUserRequest) (*User, error)

	// Login authenticates a user and returns a token
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)

	// Logout invalidates a user's authentication token
	Logout(ctx context.Context, token string) error

	// GetUser retrieves a user by ID
	GetUser(ctx context.Context, id string) (*User, error)

	// UpdateUser modifies user information
	UpdateUser(ctx context.Context, id string, req UpdateUserRequest) (*User, error)

	// DeleteUser removes a user account
	DeleteUser(ctx context.Context, id string) error

	// ValidateToken verifies an auth token and returns the user
	ValidateToken(ctx context.Context, token string) (*User, error)
}

// service is the concrete implementation of Service.
type service struct {
	repo           Repository
	authProvider   AuthProvider
	passwordHasher PasswordHasher
	config         Config
}

// NewService creates a new Service instance.
// It accepts interfaces for repository, auth provider, and password hasher,
// following the dependency injection pattern.
func NewService(repo Repository, authProvider AuthProvider, passwordHasher PasswordHasher, config Config) *service {
	return &service{
		repo:           repo,
		authProvider:   authProvider,
		passwordHasher: passwordHasher,
		config:         config,
	}
}

// Signup creates a new user account
func (s *service) Signup(ctx context.Context, req CreateUserRequest) (*User, error) {
	// TODO: Validate email format
	// TODO: Validate password meets minimum requirements
	// TODO: Check if email already exists
	// TODO: Hash password
	// TODO: Create user with repository
	return nil, nil
}

// Login authenticates a user and returns a token
func (s *service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	// TODO: Get user by email
	// TODO: Compare password hash
	// TODO: Generate auth token
	// TODO: Return user and token
	return nil, nil
}

// Logout invalidates a user's authentication token
func (s *service) Logout(ctx context.Context, token string) error {
	// TODO: Revoke token using auth provider
	return nil
}

// GetUser retrieves a user by ID
func (s *service) GetUser(ctx context.Context, id string) (*User, error) {
	// TODO: Retrieve user from repository
	return nil, nil
}

// UpdateUser modifies user information
func (s *service) UpdateUser(ctx context.Context, id string, req UpdateUserRequest) (*User, error) {
	// TODO: Get existing user
	// TODO: Apply updates (email, password)
	// TODO: Hash new password if provided
	// TODO: Update user in repository
	return nil, nil
}

// DeleteUser removes a user account
func (s *service) DeleteUser(ctx context.Context, id string) error {
	// TODO: Delete user from repository
	return nil
}

// ValidateToken verifies an auth token and returns the user
func (s *service) ValidateToken(ctx context.Context, token string) (*User, error) {
	// TODO: Validate token using auth provider
	return nil, nil
}
