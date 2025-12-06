package size

import (
	"context"

	"bot-go/internal/signals"
)

// NOFSignal computes Number of Fields
type NOFSignal struct{}

// NewNOFSignal creates a new NOF signal
func NewNOFSignal() *NOFSignal {
	return &NOFSignal{}
}

// Metadata returns information about this signal
func (s *NOFSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "NOF",
		FullName:    "Number of Fields",
		Category:    signals.CategorySize,
		Scope:       signals.ScopeClass,
		Description: "Count of fields/attributes in a class",
		Unit:        "count",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *NOFSignal) Dependencies() []string {
	return nil
}

// ComputeClass computes NOF for a class
func (s *NOFSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}

// NOPASignal computes Number of Public Attributes
type NOPASignal struct{}

// NewNOPASignal creates a new NOPA signal
func NewNOPASignal() *NOPASignal {
	return &NOPASignal{}
}

// Metadata returns information about this signal
func (s *NOPASignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "NOPA",
		FullName:    "Number of Public Attributes",
		Category:    signals.CategorySize,
		Scope:       signals.ScopeClass,
		Description: "Count of publicly accessible fields in a class",
		Unit:        "count",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *NOPASignal) Dependencies() []string {
	return nil
}

// ComputeClass computes NOPA for a class
func (s *NOPASignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
