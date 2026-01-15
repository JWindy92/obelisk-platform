package featureflag

// Provider defines the interface for retrieving feature flag states.
// This allows different backends (static config, database, remote service, etc.)
type Provider interface {
	// IsEnabled checks if a feature flag is enabled.
	IsEnabled(flagName string) bool
}
