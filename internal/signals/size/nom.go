package size

import (
	"context"

	"bot-go/internal/signals"
)

// NOMSignal computes Number of Methods
type NOMSignal struct{}

// NewNOMSignal creates a new NOM signal
func NewNOMSignal() *NOMSignal {
	return &NOMSignal{}
}

// Metadata returns information about this signal
func (s *NOMSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "NOM",
		FullName:    "Number of Methods",
		Category:    signals.CategorySize,
		Scope:       signals.ScopeClass,
		Description: "Count of methods in a class",
		Unit:        "count",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *NOMSignal) Dependencies() []string {
	return nil
}

// ComputeClass computes NOM for a class
func (s *NOMSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}

// NOMNAMMSignal computes Number of Methods without Accessors/Mutators
type NOMNAMMSignal struct{}

// NewNOMNAMMSignal creates a new NOMNAMM signal
func NewNOMNAMMSignal() *NOMNAMMSignal {
	return &NOMNAMMSignal{}
}

// Metadata returns information about this signal
func (s *NOMNAMMSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "NOMNAMM",
		FullName:    "Number of Methods without Accessors/Mutators",
		Category:    signals.CategorySize,
		Scope:       signals.ScopeClass,
		Description: "Count of methods excluding simple getter and setter methods",
		Unit:        "count",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *NOMNAMMSignal) Dependencies() []string {
	return nil
}

// ComputeClass computes NOMNAMM for a class
func (s *NOMNAMMSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
