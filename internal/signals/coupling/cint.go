package coupling

import (
	"context"

	"bot-go/internal/signals"
)

// CINTSignal computes Coupling Intensity (distinct methods called)
type CINTSignal struct{}

// NewCINTSignal creates a new CINT signal
func NewCINTSignal() *CINTSignal {
	return &CINTSignal{}
}

// Metadata returns information about this signal
func (s *CINTSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "CINT",
		FullName:    "Coupling Intensity",
		Category:    signals.CategoryCoupling,
		Scope:       signals.ScopeMethod,
		Description: "Number of distinct methods called by a method",
		Unit:        "count",
		LowerBetter: true, // Lower CINT means less coupling intensity
	}
}

// Dependencies returns names of signals this signal depends on
func (s *CINTSignal) Dependencies() []string {
	return nil
}

// ComputeMethod computes CINT for a method
func (s *CINTSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
