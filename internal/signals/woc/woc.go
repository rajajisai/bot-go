package woc

import (
	"context"

	"bot-go/internal/signals"
)

// WOCSignal computes Weight of Class
type WOCSignal struct{}

// NewWOCSignal creates a new WOC signal
func NewWOCSignal() *WOCSignal {
	return &WOCSignal{}
}

// Metadata returns information about this signal
func (s *WOCSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "WOC",
		FullName:    "Weight of Class",
		Category:    signals.CategoryComposite,
		Scope:       signals.ScopeClass,
		Description: "Ratio of functional methods to total methods: (NOM - NOAM) / NOM",
		Unit:        "ratio",
		LowerBetter: false, // Higher WOC means more functional methods (less data-class-like)
	}
}

// Dependencies returns names of signals this signal depends on
func (s *WOCSignal) Dependencies() []string {
	return []string{"NOM", "NOAM"}
}

// ComputeWithDependencies computes WOC using NOM and NOAM
func (s *WOCSignal) ComputeWithDependencies(ctx context.Context, target interface{}, deps map[string]signals.SignalResult, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
