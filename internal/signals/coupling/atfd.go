package coupling

import (
	"context"

	"bot-go/internal/signals"
)

// ATFDSignal computes Access To Foreign Data
type ATFDSignal struct{}

// NewATFDSignal creates a new ATFD signal
func NewATFDSignal() *ATFDSignal {
	return &ATFDSignal{}
}

// Metadata returns information about this signal
func (s *ATFDSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "ATFD",
		FullName:    "Access To Foreign Data",
		Category:    signals.CategoryCoupling,
		Scope:       signals.ScopeClass,
		Description: "Number of external class attributes accessed directly or via accessor methods",
		Unit:        "count",
		LowerBetter: true, // Lower ATFD means less coupling to foreign data
	}
}

// Dependencies returns names of signals this signal depends on
func (s *ATFDSignal) Dependencies() []string {
	return nil
}

// ComputeClass computes ATFD for a class
func (s *ATFDSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}

// ComputeMethod computes ATFD for a method
func (s *ATFDSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
