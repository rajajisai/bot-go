package size

import (
	"context"

	"bot-go/internal/signals"
)

// LOCSignal computes Lines of Code
type LOCSignal struct{}

// NewLOCSignal creates a new LOC signal
func NewLOCSignal() *LOCSignal {
	return &LOCSignal{}
}

// Metadata returns information about this signal
func (s *LOCSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "LOC",
		FullName:    "Lines of Code",
		Category:    signals.CategorySize,
		Scope:       signals.ScopeClass,
		Description: "Total lines of code in a class or method",
		Unit:        "lines",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *LOCSignal) Dependencies() []string {
	return nil
}

// ComputeClass computes LOC for a class
func (s *LOCSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}

// ComputeMethod computes LOC for a method
func (s *LOCSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}

// LOCNAMMSignal computes Lines of Code without Accessors/Mutators
type LOCNAMMSignal struct{}

// NewLOCNAMMSignal creates a new LOCNAMM signal
func NewLOCNAMMSignal() *LOCNAMMSignal {
	return &LOCNAMMSignal{}
}

// Metadata returns information about this signal
func (s *LOCNAMMSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "LOCNAMM",
		FullName:    "Lines of Code without Accessors/Mutators",
		Category:    signals.CategorySize,
		Scope:       signals.ScopeClass,
		Description: "Lines of code excluding simple getter and setter methods",
		Unit:        "lines",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *LOCNAMMSignal) Dependencies() []string {
	return nil
}

// ComputeClass computes LOCNAMM for a class
func (s *LOCNAMMSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
