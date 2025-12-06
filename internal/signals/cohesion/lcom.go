package cohesion

import (
	"context"

	"bot-go/internal/signals"
)

// LCOMSignal computes Lack of Cohesion in Methods (LCOM1)
type LCOMSignal struct{}

// NewLCOMSignal creates a new LCOM signal
func NewLCOMSignal() *LCOMSignal {
	return &LCOMSignal{}
}

// Metadata returns information about this signal
func (s *LCOMSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "LCOM",
		FullName:    "Lack of Cohesion in Methods",
		Category:    signals.CategoryCohesion,
		Scope:       signals.ScopeClass,
		Description: "Number of method pairs that share no fields minus pairs that share fields",
		Unit:        "count",
		LowerBetter: true, // Lower LCOM means better cohesion
	}
}

// Dependencies returns names of signals this signal depends on
func (s *LCOMSignal) Dependencies() []string {
	return nil
}

// ComputeClass computes LCOM for a class
func (s *LCOMSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}

// LCOM4Signal computes LCOM4 (connected components variant)
type LCOM4Signal struct{}

// NewLCOM4Signal creates a new LCOM4 signal
func NewLCOM4Signal() *LCOM4Signal {
	return &LCOM4Signal{}
}

// Metadata returns information about this signal
func (s *LCOM4Signal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "LCOM4",
		FullName:    "Lack of Cohesion in Methods 4",
		Category:    signals.CategoryCohesion,
		Scope:       signals.ScopeClass,
		Description: "Number of connected components in the method-field graph",
		Unit:        "count",
		LowerBetter: true, // LCOM4 = 1 is ideal (all methods connected)
	}
}

// Dependencies returns names of signals this signal depends on
func (s *LCOM4Signal) Dependencies() []string {
	return nil
}

// ComputeClass computes LCOM4 for a class
func (s *LCOM4Signal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
