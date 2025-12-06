package complexity

import (
	"context"

	"bot-go/internal/signals"
)

// CYCLOSignal computes Cyclomatic Complexity
type CYCLOSignal struct{}

// NewCYCLOSignal creates a new CYCLO signal
func NewCYCLOSignal() *CYCLOSignal {
	return &CYCLOSignal{}
}

// Metadata returns information about this signal
func (s *CYCLOSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "CYCLO",
		FullName:    "Cyclomatic Complexity",
		Category:    signals.CategoryComplexity,
		Scope:       signals.ScopeMethod,
		Description: "Number of linearly independent paths through code",
		Unit:        "count",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *CYCLOSignal) Dependencies() []string {
	return nil
}

// ComputeMethod computes CYCLO for a method
func (s *CYCLOSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
