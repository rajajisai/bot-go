package coupling

import (
	"context"

	"bot-go/internal/signals"
)

// FDPSignal computes Foreign Data Providers
type FDPSignal struct{}

// NewFDPSignal creates a new FDP signal
func NewFDPSignal() *FDPSignal {
	return &FDPSignal{}
}

// Metadata returns information about this signal
func (s *FDPSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "FDP",
		FullName:    "Foreign Data Providers",
		Category:    signals.CategoryCoupling,
		Scope:       signals.ScopeMethod,
		Description: "Number of distinct classes whose attributes are accessed by a method",
		Unit:        "count",
		LowerBetter: true, // Lower FDP means less dependency on foreign data
	}
}

// Dependencies returns names of signals this signal depends on
func (s *FDPSignal) Dependencies() []string {
	return nil
}

// ComputeMethod computes FDP for a method
func (s *FDPSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
