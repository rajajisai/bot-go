package coupling

import (
	"context"

	"bot-go/internal/signals"
)

// FANOUTSignal computes Number of Called Classes
type FANOUTSignal struct{}

// NewFANOUTSignal creates a new FANOUT signal
func NewFANOUTSignal() *FANOUTSignal {
	return &FANOUTSignal{}
}

// Metadata returns information about this signal
func (s *FANOUTSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "FANOUT",
		FullName:    "Number of Called Classes",
		Category:    signals.CategoryCoupling,
		Scope:       signals.ScopeMethod,
		Description: "Number of distinct classes that a method calls methods on",
		Unit:        "count",
		LowerBetter: true, // Lower FANOUT means less coupling
	}
}

// Dependencies returns names of signals this signal depends on
func (s *FANOUTSignal) Dependencies() []string {
	return nil
}

// ComputeMethod computes FANOUT for a method
func (s *FANOUTSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
