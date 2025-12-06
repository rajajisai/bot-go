package coupling

import (
	"context"

	"bot-go/internal/signals"
)

// CBOSignal computes Coupling Between Objects
type CBOSignal struct{}

// NewCBOSignal creates a new CBO signal
func NewCBOSignal() *CBOSignal {
	return &CBOSignal{}
}

// Metadata returns information about this signal
func (s *CBOSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "CBO",
		FullName:    "Coupling Between Objects",
		Category:    signals.CategoryCoupling,
		Scope:       signals.ScopeClass,
		Description: "Number of classes coupled to a given class (both directions)",
		Unit:        "count",
		LowerBetter: true, // Lower CBO means less coupling
	}
}

// Dependencies returns names of signals this signal depends on
func (s *CBOSignal) Dependencies() []string {
	return nil
}

// ComputeClass computes CBO for a class
func (s *CBOSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
