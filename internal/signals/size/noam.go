package size

import (
	"context"

	"bot-go/internal/signals"
	"bot-go/internal/signals/util"
)

// NOAMSignal computes Number of Accessor Methods
type NOAMSignal struct {
	accessorDetector *util.AccessorDetector
}

// NewNOAMSignal creates a new NOAM signal
func NewNOAMSignal() *NOAMSignal {
	return &NOAMSignal{
		accessorDetector: util.NewAccessorDetector(),
	}
}

// Metadata returns information about this signal
func (s *NOAMSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "NOAM",
		FullName:    "Number of Accessor Methods",
		Category:    signals.CategorySize,
		Scope:       signals.ScopeClass,
		Description: "Count of getter and setter methods in a class",
		Unit:        "count",
		LowerBetter: false, // More accessors isn't necessarily bad
	}
}

// Dependencies returns names of signals this signal depends on
func (s *NOAMSignal) Dependencies() []string {
	return nil
}

// ComputeClass computes NOAM for a class
// Counts getter and setter methods using the accessor detector
func (s *NOAMSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	if classInfo == nil {
		return signals.NewSignalResultError("NOAM", signals.ErrNilInput), nil
	}

	// Count accessor methods
	getterCount := 0
	setterCount := 0
	for _, method := range classInfo.Methods {
		if s.accessorDetector.IsGetter(method) {
			getterCount++
		} else if s.accessorDetector.IsSetter(method) {
			setterCount++
		}
	}

	totalAccessors := getterCount + setterCount

	return signals.NewSignalResultWithMetadata("NOAM", float64(totalAccessors), map[string]any{
		"getter_count":   getterCount,
		"setter_count":   setterCount,
		"total_methods":  len(classInfo.Methods),
		"accessor_ratio": safeRatio(totalAccessors, len(classInfo.Methods)),
	}), nil
}

// safeRatio computes a ratio, returning 0 if denominator is 0
func safeRatio(numerator, denominator int) float64 {
	if denominator == 0 {
		return 0
	}
	return float64(numerator) / float64(denominator)
}
