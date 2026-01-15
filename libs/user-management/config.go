package usermgmt

// Config holds configuration options for the user management system.
type Config struct {
	// TableName specifies the database table name for users.
	// Defaults to "users" if not specified.
	TableName string

	// PasswordMinLength sets the minimum password length requirement.
	// Defaults to 8 if not specified.
	PasswordMinLength int
}

// DefaultConfig returns a Config with sensible defaults.
func DefaultConfig() Config {
	return Config{
		TableName:         "users",
		PasswordMinLength: 8,
	}
}
