package complexity

import (
	"context"

	"bot-go/internal/signals"
)

// WMCSignal computes Weighted Methods Count (sum of cyclomatic complexity)
type WMCSignal struct{}

// NewWMCSignal creates a new WMC signal
func NewWMCSignal() *WMCSignal {
	return &WMCSignal{}
}

// Metadata returns information about this signal
func (s *WMCSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "WMC",
		FullName:    "Weighted Methods Count",
		Category:    signals.CategoryComplexity,
		Scope:       signals.ScopeClass,
		Description: "Sum of cyclomatic complexity of all methods in a class",
		Unit:        "count",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *WMCSignal) Dependencies() []string {
	return []string{"CYCLO"}
}

// ComputeClass computes WMC for a class
func (s *WMCSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}

// Aggregate sums CYCLO values from all methods
func (s *WMCSignal) Aggregate(ctx context.Context, methodResults []signals.SignalResult) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}

// WMCNAMMSignal computes WMC without Accessor/Mutator methods
type WMCNAMMSignal struct{}

// NewWMCNAMMSignal creates a new WMCNAMM signal
func NewWMCNAMMSignal() *WMCNAMMSignal {
	return &WMCNAMMSignal{}
}

// Metadata returns information about this signal
func (s *WMCNAMMSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "WMCNAMM",
		FullName:    "Weighted Methods Count without Accessors/Mutators",
		Category:    signals.CategoryComplexity,
		Scope:       signals.ScopeClass,
		Description: "Sum of cyclomatic complexity excluding accessor methods",
		Unit:        "count",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *WMCNAMMSignal) Dependencies() []string {
	return []string{"CYCLO"}
}

// ComputeClass computes WMCNAMM for a class
func (s *WMCNAMMSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}

// AMCSignal computes Average Method Complexity
type AMCSignal struct{}

// NewAMCSignal creates a new AMC signal
func NewAMCSignal() *AMCSignal {
	return &AMCSignal{}
}

// Metadata returns information about this signal
func (s *AMCSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "AMC",
		FullName:    "Average Method Complexity",
		Category:    signals.CategoryComplexity,
		Scope:       signals.ScopeClass,
		Description: "Average cyclomatic complexity per method (WMC / NOM)",
		Unit:        "ratio",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *AMCSignal) Dependencies() []string {
	return []string{"WMC", "NOM"}
}

// ComputeWithDependencies computes AMC using WMC and NOM
func (s *AMCSignal) ComputeWithDependencies(ctx context.Context, target interface{}, deps map[string]signals.SignalResult, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
