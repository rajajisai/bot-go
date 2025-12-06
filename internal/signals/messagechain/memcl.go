package messagechain

import (
	"context"

	"bot-go/internal/signals"
)

// MeMCLSignal computes Mean Message Chain Length
type MeMCLSignal struct{}

// NewMeMCLSignal creates a new MeMCL signal
func NewMeMCLSignal() *MeMCLSignal {
	return &MeMCLSignal{}
}

// Metadata returns information about this signal
func (s *MeMCLSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "MeMCL",
		FullName:    "Mean Message Chain Length",
		Category:    signals.CategoryMessageChain,
		Scope:       signals.ScopeMethod,
		Description: "Average chain length across all method calls in a method",
		Unit:        "ratio",
		LowerBetter: true, // Shorter average chains are better
	}
}

// Dependencies returns names of signals this signal depends on
func (s *MeMCLSignal) Dependencies() []string {
	return nil
}

// ComputeMethod computes MeMCL for a method
func (s *MeMCLSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
