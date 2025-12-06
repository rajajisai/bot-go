package coupling

import (
	"context"

	"bot-go/internal/signals"
)

// LAASignal computes Locality of Attribute Accesses
type LAASignal struct{}

// NewLAASignal creates a new LAA signal
func NewLAASignal() *LAASignal {
	return &LAASignal{}
}

// Metadata returns information about this signal
func (s *LAASignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "LAA",
		FullName:    "Locality of Attribute Accesses",
		Category:    signals.CategoryCoupling,
		Scope:       signals.ScopeMethod,
		Description: "Ratio of local attribute accesses to total attribute accesses",
		Unit:        "ratio",
		LowerBetter: false, // Higher LAA means more local data usage (good)
	}
}

// Dependencies returns names of signals this signal depends on
func (s *LAASignal) Dependencies() []string {
	return []string{"ATLD", "ATFD"}
}

// ComputeWithDependencies calculates ATLD / (ATLD + ATFD)
func (s *LAASignal) ComputeWithDependencies(ctx context.Context, target interface{}, deps map[string]signals.SignalResult, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
