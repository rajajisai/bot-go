package complexity

import (
	"context"

	"bot-go/internal/signals"
)

// MAXNESTINGSignal computes Maximum Nesting Level
type MAXNESTINGSignal struct{}

// NewMAXNESTINGSignal creates a new MAXNESTING signal
func NewMAXNESTINGSignal() *MAXNESTINGSignal {
	return &MAXNESTINGSignal{}
}

// Metadata returns information about this signal
func (s *MAXNESTINGSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "MAXNESTING",
		FullName:    "Maximum Nesting Level",
		Category:    signals.CategoryComplexity,
		Scope:       signals.ScopeMethod,
		Description: "Maximum depth of nested control structures in a method",
		Unit:        "levels",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *MAXNESTINGSignal) Dependencies() []string {
	return nil
}

// ComputeMethod computes MAXNESTING for a method
func (s *MAXNESTINGSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
