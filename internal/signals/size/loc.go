package size

import (
	"context"

	"bot-go/internal/signals"
	"bot-go/internal/signals/util"
)

// LOCSignal computes Lines of Code
type LOCSignal struct {
	accessorDetector *util.AccessorDetector
}

// NewLOCSignal creates a new LOC signal
func NewLOCSignal() *LOCSignal {
	return &LOCSignal{
		accessorDetector: util.NewAccessorDetector(),
	}
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

// ComputeClass computes LOC for a class using Range from the CodeGraph
func (s *LOCSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	if classInfo == nil {
		return signals.NewSignalResultError("LOC", signals.ErrNilInput), nil
	}

	// Calculate LOC from Range (end line - start line + 1)
	loc := calculateLOCFromRange(classInfo.Range)

	return signals.NewSignalResultWithMetadata("LOC", float64(loc), map[string]any{
		"start_line": classInfo.Range.Start.Line,
		"end_line":   classInfo.Range.End.Line,
	}), nil
}

// ComputeMethod computes LOC for a method using Range from the CodeGraph
func (s *LOCSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	if methodInfo == nil {
		return signals.NewSignalResultError("LOC", signals.ErrNilInput), nil
	}

	// Calculate LOC from Range (end line - start line + 1)
	loc := calculateLOCFromRange(methodInfo.Range)

	return signals.NewSignalResultWithMetadata("LOC", float64(loc), map[string]any{
		"start_line": methodInfo.Range.Start.Line,
		"end_line":   methodInfo.Range.End.Line,
	}), nil
}

// LOCNAMMSignal computes Lines of Code without Accessors/Mutators
type LOCNAMMSignal struct {
	accessorDetector *util.AccessorDetector
}

// NewLOCNAMMSignal creates a new LOCNAMM signal
func NewLOCNAMMSignal() *LOCNAMMSignal {
	return &LOCNAMMSignal{
		accessorDetector: util.NewAccessorDetector(),
	}
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
// Total class LOC minus the LOC of all accessor methods
func (s *LOCNAMMSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	if classInfo == nil {
		return signals.NewSignalResultError("LOCNAMM", signals.ErrNilInput), nil
	}

	// Calculate total class LOC from Range
	totalLOC := calculateLOCFromRange(classInfo.Range)

	// Subtract LOC of accessor methods
	accessorLOC := 0
	accessorCount := 0
	for _, method := range classInfo.Methods {
		if s.accessorDetector.IsAccessor(method) {
			accessorLOC += calculateLOCFromRange(method.Range)
			accessorCount++
		}
	}

	locNamm := totalLOC - accessorLOC
	if locNamm < 0 {
		locNamm = 0
	}

	return signals.NewSignalResultWithMetadata("LOCNAMM", float64(locNamm), map[string]any{
		"total_loc":      totalLOC,
		"accessor_loc":   accessorLOC,
		"accessor_count": accessorCount,
	}), nil
}
