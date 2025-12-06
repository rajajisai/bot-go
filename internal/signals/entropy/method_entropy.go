package entropy

import (
	"context"

	"bot-go/internal/signals"
)

// MethodEntropySignal computes entropy for a method's code
type MethodEntropySignal struct{}

// NewMethodEntropySignal creates a new MethodEntropy signal
func NewMethodEntropySignal() *MethodEntropySignal {
	return &MethodEntropySignal{}
}

// Metadata returns information about this signal
func (s *MethodEntropySignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "MethodEntropy",
		FullName:    "Method Entropy",
		Category:    signals.CategoryEntropy,
		Scope:       signals.ScopeMethod,
		Description: "Cross-entropy of a method's code against the corpus model",
		Unit:        "bits",
		LowerBetter: true, // Lower entropy means more predictable/natural code
	}
}

// Dependencies returns names of signals this signal depends on
func (s *MethodEntropySignal) Dependencies() []string {
	return nil
}

// ComputeMethod computes entropy for a method
func (s *MethodEntropySignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}

// ClassEntropySignal computes entropy for a class's code
type ClassEntropySignal struct{}

// NewClassEntropySignal creates a new ClassEntropy signal
func NewClassEntropySignal() *ClassEntropySignal {
	return &ClassEntropySignal{}
}

// Metadata returns information about this signal
func (s *ClassEntropySignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "ClassEntropy",
		FullName:    "Class Entropy",
		Category:    signals.CategoryEntropy,
		Scope:       signals.ScopeClass,
		Description: "Cross-entropy of a class's code against the corpus model",
		Unit:        "bits",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *ClassEntropySignal) Dependencies() []string {
	return nil
}

// ComputeClass computes entropy for a class
func (s *ClassEntropySignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
