package cohesion

import (
	"context"

	"bot-go/internal/signals"
)

// ATLDSignal computes Access To Local Data
type ATLDSignal struct{}

// NewATLDSignal creates a new ATLD signal
func NewATLDSignal() *ATLDSignal {
	return &ATLDSignal{}
}

// Metadata returns information about this signal
func (s *ATLDSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "ATLD",
		FullName:    "Access To Local Data",
		Category:    signals.CategoryCohesion,
		Scope:       signals.ScopeMethod,
		Description: "Number of class attributes accessed by a method within the same class",
		Unit:        "count",
		LowerBetter: false, // Higher ATLD means method works with local data (good)
	}
}

// Dependencies returns names of signals this signal depends on
func (s *ATLDSignal) Dependencies() []string {
	return nil
}

// ComputeMethod computes ATLD for a method
func (s *ATLDSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
