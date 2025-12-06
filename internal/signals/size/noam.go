package size

import (
	"context"

	"bot-go/internal/signals"
)

// NOAMSignal computes Number of Accessor Methods
type NOAMSignal struct{}

// NewNOAMSignal creates a new NOAM signal
func NewNOAMSignal() *NOAMSignal {
	return &NOAMSignal{}
}

// Metadata returns information about this signal
func (s *NOAMSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "NOAM",
		FullName:    "Number of Accessor Methods",
		Category:    signals.CategorySize,
		Scope:       signals.ScopeClass,
		Description: "Count of getter and setter methods in a class",
		Unit:        "count",
		LowerBetter: false, // More accessors isn't necessarily bad
	}
}

// Dependencies returns names of signals this signal depends on
func (s *NOAMSignal) Dependencies() []string {
	return nil
}

// ComputeClass computes NOAM for a class
func (s *NOAMSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
