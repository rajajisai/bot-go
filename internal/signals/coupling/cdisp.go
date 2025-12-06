package coupling

import (
	"context"

	"bot-go/internal/signals"
)

// CDISPSignal computes Coupling Dispersion (FANOUT / CINT)
type CDISPSignal struct{}

// NewCDISPSignal creates a new CDISP signal
func NewCDISPSignal() *CDISPSignal {
	return &CDISPSignal{}
}

// Metadata returns information about this signal
func (s *CDISPSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "CDISP",
		FullName:    "Coupling Dispersion",
		Category:    signals.CategoryCoupling,
		Scope:       signals.ScopeMethod,
		Description: "Ratio of called classes to called methods (FANOUT / CINT)",
		Unit:        "ratio",
		LowerBetter: false, // Low CDISP with high CINT indicates intensive coupling
	}
}

// Dependencies returns names of signals this signal depends on
func (s *CDISPSignal) Dependencies() []string {
	return []string{"CINT", "FANOUT"}
}

// ComputeWithDependencies computes CDISP using CINT and FANOUT
func (s *CDISPSignal) ComputeWithDependencies(ctx context.Context, target interface{}, deps map[string]signals.SignalResult, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
