package coupling

import (
	"context"

	"bot-go/internal/signals"
)

// RFCSignal computes Response For a Class
type RFCSignal struct{}

// NewRFCSignal creates a new RFC signal
func NewRFCSignal() *RFCSignal {
	return &RFCSignal{}
}

// Metadata returns information about this signal
func (s *RFCSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "RFC",
		FullName:    "Response For a Class",
		Category:    signals.CategoryCoupling,
		Scope:       signals.ScopeClass,
		Description: "Number of methods that can be invoked in response to a message (class methods + called methods)",
		Unit:        "count",
		LowerBetter: true, // Lower RFC means simpler response set
	}
}

// Dependencies returns names of signals this signal depends on
func (s *RFCSignal) Dependencies() []string {
	return []string{"NOM", "CINT"}
}

// ComputeClass computes RFC for a class
func (s *RFCSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
