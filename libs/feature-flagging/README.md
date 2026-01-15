# Feature Flagging Library

A simple, extensible feature flag library that integrates seamlessly with dependency injection.

## Design Philosophy

- **Simple by default**: Start with basic on/off flags
- **Extensible**: Add new providers (database, Redis, remote services) later
- **DI-friendly**: Works naturally with your existing dependency injection pattern
- **Fail-safe**: Unknown flags default to disabled

## Usage

### Basic Flag Checks

```go
package main

import "github.com/JWindy92/obelisk-platform/libs/feature-flagging"

func main() {
    // Initialize with static flags
    provider := featureflag.NewStaticProvider(map[string]bool{
        "new-feature": true,
        "beta-api":    false,
    })
    
    ff := featureflag.New(provider)
    
    // Simple check
    if ff.IsEnabled("new-feature") {
        // New code path
    } else {
        // Old code path
    }
}
```

### Selecting Implementations (DI Pattern)

```go
// Initialize with feature flag
func InitializeUserService(repo Repository, ff *featureflag.Manager) UserService {
    return ff.Select("user-service-v2",
        func() any { return NewUserServiceV2(repo) },  // enabled
        func() any { return NewUserServiceV1(repo) },  // fallback
    ).(UserService)
}

// Application code doesn't change
svc := InitializeUserService(repo, ffManager)
svc.CreateUser(ctx, user) // Uses V2 if flag enabled
```

### Conditional Execution

```go
ff.When("send-welcome-email",
    func() {
        sendEmail(user)
    },
    func() {
        // Do nothing or old behavior
    },
)
```

## Extending with Custom Providers

Implement the `Provider` interface:

```go
type Provider interface {
    IsEnabled(flagName string) bool
}
```

Examples of future providers:
- Database-backed (using store.Store)
- Environment variables
- Redis
- Remote services (LaunchDarkly, etc.)

## Current Status

✅ Core Manager API  
✅ Static provider (map-based)  
✅ `IsEnabled()` / `IsDisabled()` checks  
✅ `Select()` for DI integration  
✅ `When()` for conditional execution  
⏳ Database provider (coming later)  
⏳ Environment variable provider (coming later)  
⏳ User-specific flags (coming later)
