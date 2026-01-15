package main

import (
	"fmt"

	featureflag "github.com/JWindy92/obelisk-platform/libs/feature-flagging"
)

func main() {
	// Initialize feature flags
	provider := featureflag.NewStaticProvider(map[string]bool{
		"new-algorithm":  true,
		"beta-feature":   false,
		"enhanced-logic": true,
	})

	ff := featureflag.New(provider)

	fmt.Println("=== Feature Flag Examples ===\n")

	// Example 1: Simple flag check
	example1_SimpleCheck(ff)

	// Example 2: Select implementation based on flag
	example2_SelectImplementation(ff)

	// Example 3: Conditional execution
	example3_ConditionalExecution(ff)
}

// Example 1: Simple flag check
func example1_SimpleCheck(ff *featureflag.Manager) {
	fmt.Println("1. Simple Flag Check")

	if ff.IsEnabled("new-algorithm") {
		fmt.Println("   ✓ Using new algorithm")
		processWithNewAlgorithm()
	} else {
		fmt.Println("   ✗ Using old algorithm")
		processWithOldAlgorithm()
	}

	if ff.IsDisabled("beta-feature") {
		fmt.Println("   ✗ Beta feature is disabled")
	}

	fmt.Println()
}

// Example 2: Select implementation (DI pattern)
func example2_SelectImplementation(ff *featureflag.Manager) {
	fmt.Println("2. Select Implementation")

	// Select which processor to use based on flag
	processor := ff.Select("enhanced-logic",
		func() any { return &EnhancedProcessor{} },
		func() any { return &StandardProcessor{} },
	).(Processor)

	result := processor.Process("test data")
	fmt.Printf("   Result: %s\n\n", result)
}

// Example 3: Conditional execution with When
func example3_ConditionalExecution(ff *featureflag.Manager) {
	fmt.Println("3. Conditional Execution")

	ff.When("new-algorithm",
		func() {
			fmt.Println("   ✓ Executing new code path")
		},
		func() {
			fmt.Println("   ✗ Executing fallback code path")
		},
	)

	fmt.Println()
}

// Processor interface for demonstration
type Processor interface {
	Process(data string) string
}

type StandardProcessor struct{}

func (p *StandardProcessor) Process(data string) string {
	return "Standard: " + data
}

type EnhancedProcessor struct{}

func (p *EnhancedProcessor) Process(data string) string {
	return "Enhanced: " + data + " (with improvements!)"
}

func processWithNewAlgorithm() {
	// New algorithm implementation
}

func processWithOldAlgorithm() {
	// Old algorithm implementation
}
