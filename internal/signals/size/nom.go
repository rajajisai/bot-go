package size

import (
	"context"

	"bot-go/internal/signals"
	"bot-go/internal/signals/util"
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
// Counts all methods contained in the class from the CodeGraph
func (s *NOMSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	if classInfo == nil {
		return signals.NewSignalResultError("NOM", signals.ErrNilInput), nil
	}

	// Count methods from ClassInfo (populated from CodeGraph)
	methodCount := len(classInfo.Methods)

	return signals.NewSignalResultWithMetadata("NOM", float64(methodCount), map[string]any{
		"class_id":   classInfo.NodeID,
		"class_name": classInfo.Name,
	}), nil
}

// NOMNAMMSignal computes Number of Methods without Accessors/Mutators
type NOMNAMMSignal struct {
	accessorDetector *util.AccessorDetector
}

// NewNOMNAMMSignal creates a new NOMNAMM signal
func NewNOMNAMMSignal() *NOMNAMMSignal {
	return &NOMNAMMSignal{
		accessorDetector: util.NewAccessorDetector(),
	}
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
// Total method count minus accessor methods
func (s *NOMNAMMSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	if classInfo == nil {
		return signals.NewSignalResultError("NOMNAMM", signals.ErrNilInput), nil
	}

	// Count all methods
	totalMethods := len(classInfo.Methods)

	// Count accessor methods
	accessorCount := 0
	for _, method := range classInfo.Methods {
		if s.accessorDetector.IsAccessor(method) {
			accessorCount++
		}
	}

	nomNamm := totalMethods - accessorCount

	return signals.NewSignalResultWithMetadata("NOMNAMM", float64(nomNamm), map[string]any{
		"total_methods":  totalMethods,
		"accessor_count": accessorCount,
	}), nil
}
