package messagechain

import (
	"context"

	"bot-go/internal/signals"
)

// NMCSSignal computes Number of Message Chain Statements
type NMCSSignal struct{}

// NewNMCSSignal creates a new NMCS signal
func NewNMCSSignal() *NMCSSignal {
	return &NMCSSignal{}
}

// Metadata returns information about this signal
func (s *NMCSSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "NMCS",
		FullName:    "Number of Message Chain Statements",
		Category:    signals.CategoryMessageChain,
		Scope:       signals.ScopeMethod,
		Description: "Count of statements containing message chains (chain length > 1)",
		Unit:        "count",
		LowerBetter: true, // Fewer chain statements is better
	}
}

// Dependencies returns names of signals this signal depends on
func (s *NMCSSignal) Dependencies() []string {
	return nil
}

// ComputeMethod computes NMCS for a method
func (s *NMCSSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
