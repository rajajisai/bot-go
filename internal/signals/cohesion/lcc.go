package cohesion

import (
	"context"

	"bot-go/internal/signals"
)

// LCCSignal computes Loose Class Cohesion (transitive TCC)
type LCCSignal struct{}

// NewLCCSignal creates a new LCC signal
func NewLCCSignal() *LCCSignal {
	return &LCCSignal{}
}

// Metadata returns information about this signal
func (s *LCCSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "LCC",
		FullName:    "Loose Class Cohesion",
		Category:    signals.CategoryCohesion,
		Scope:       signals.ScopeClass,
		Description: "Ratio of transitively connected method pairs to maximum possible pairs",
		Unit:        "ratio",
		LowerBetter: false, // Higher LCC means better cohesion
	}
}

// Dependencies returns names of signals this signal depends on
func (s *LCCSignal) Dependencies() []string {
	return nil
}

// ComputeClass computes LCC for a class
func (s *LCCSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
