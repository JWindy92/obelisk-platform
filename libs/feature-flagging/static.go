package featureflag

// StaticProvider is a simple in-memory provider backed by a map.
// Useful for configuration files or simple use cases.
type StaticProvider struct {
	flags map[string]bool
}

// NewStaticProvider creates a provider with the given flag states.
func NewStaticProvider(flags map[string]bool) *StaticProvider {
	if flags == nil {
		flags = make(map[string]bool)
	}
	return &StaticProvider{
		flags: flags,
	}
}

// IsEnabled checks if a feature flag is enabled.
// Returns false if the flag doesn't exist (fail-safe default).
func (s *StaticProvider) IsEnabled(flagName string) bool {
	enabled, exists := s.flags[flagName]
	if !exists {
		return false // Fail-safe: unknown flags are disabled
	}
	return enabled
}

// Set updates a flag's state (useful for testing or runtime changes).
func (s *StaticProvider) Set(flagName string, enabled bool) {
	s.flags[flagName] = enabled
}
