package featureflag

// Manager provides the main API for working with feature flags.
type Manager struct {
	provider Provider
}

// New creates a new Manager with the given provider.
func New(provider Provider) *Manager {
	return &Manager{
		provider: provider,
	}
}

// IsEnabled checks if a feature flag is enabled.
func (m *Manager) IsEnabled(flagName string) bool {
	return m.provider.IsEnabled(flagName)
}

// IsDisabled checks if a feature flag is disabled (convenience method).
func (m *Manager) IsDisabled(flagName string) bool {
	return !m.IsEnabled(flagName)
}

// Select returns the enabled implementation if the flag is on, otherwise the fallback.
// Both enabled and fallback are factory functions that create the implementation.
func (m *Manager) Select(flagName string, enabled, fallback func() any) any {
	if m.IsEnabled(flagName) {
		return enabled()
	}
	return fallback()
}

// When executes the enabled function if the flag is on, otherwise the fallback.
// This is useful for simple conditional execution without return values.
func (m *Manager) When(flagName string, enabled, fallback func()) {
	if m.IsEnabled(flagName) {
		enabled()
	} else {
		fallback()
	}
}
