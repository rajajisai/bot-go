package messagechain

import (
	"context"

	"bot-go/internal/signals"
)

// MaMCLSignal computes Maximum Message Chain Length
type MaMCLSignal struct{}

// NewMaMCLSignal creates a new MaMCL signal
func NewMaMCLSignal() *MaMCLSignal {
	return &MaMCLSignal{}
}

// Metadata returns information about this signal
func (s *MaMCLSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "MaMCL",
		FullName:    "Maximum Message Chain Length",
		Category:    signals.CategoryMessageChain,
		Scope:       signals.ScopeMethod,
		Description: "Maximum length of method chains in a method (e.g., a.b().c().d() = 3)",
		Unit:        "count",
		LowerBetter: true, // Shorter chains are better
	}
}

// Dependencies returns names of signals this signal depends on
func (s *MaMCLSignal) Dependencies() []string {
	return nil
}

// ComputeMethod computes MaMCL for a method
func (s *MaMCLSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
