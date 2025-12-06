package complexity

import (
	"context"

	"bot-go/internal/signals"
)

// NOLVSignal computes Number of Local Variables
type NOLVSignal struct{}

// NewNOLVSignal creates a new NOLV signal
func NewNOLVSignal() *NOLVSignal {
	return &NOLVSignal{}
}

// Metadata returns information about this signal
func (s *NOLVSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "NOLV",
		FullName:    "Number of Local Variables",
		Category:    signals.CategoryComplexity,
		Scope:       signals.ScopeMethod,
		Description: "Count of local variables declared in a method (excluding parameters)",
		Unit:        "count",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *NOLVSignal) Dependencies() []string {
	return nil
}

// ComputeMethod computes NOLV for a method
func (s *NOLVSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
