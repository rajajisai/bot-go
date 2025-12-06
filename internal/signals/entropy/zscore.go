package entropy

import (
	"context"

	"bot-go/internal/signals"
)

// ZScoreSignal computes z-score relative to corpus entropy
type ZScoreSignal struct{}

// NewZScoreSignal creates a new ZScore signal
func NewZScoreSignal() *ZScoreSignal {
	return &ZScoreSignal{}
}

// Metadata returns information about this signal
func (s *ZScoreSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "ZScore",
		FullName:    "Entropy Z-Score",
		Category:    signals.CategoryEntropy,
		Scope:       signals.ScopeMethod,
		Description: "Standard deviations from corpus mean entropy (high z-score indicates unusual code)",
		Unit:        "std_dev",
		LowerBetter: true, // Lower z-score means more typical code
	}
}

// Dependencies returns names of signals this signal depends on
func (s *ZScoreSignal) Dependencies() []string {
	return []string{"MethodEntropy"}
}

// ComputeMethod computes z-score for a method
func (s *ZScoreSignal) ComputeMethod(ctx context.Context, methodInfo *signals.MethodInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}

// ComputeWithDependencies computes z-score using pre-computed entropy
func (s *ZScoreSignal) ComputeWithDependencies(ctx context.Context, target interface{}, deps map[string]signals.SignalResult, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}

// HighEntropyMethodCountSignal counts methods with entropy above threshold
type HighEntropyMethodCountSignal struct {
	threshold float64
}

// NewHighEntropyMethodCountSignal creates a new HighEntropyMethodCount signal
func NewHighEntropyMethodCountSignal(threshold float64) *HighEntropyMethodCountSignal {
	return &HighEntropyMethodCountSignal{
		threshold: threshold,
	}
}

// Metadata returns information about this signal
func (s *HighEntropyMethodCountSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "HighEntropyMethodCount",
		FullName:    "High Entropy Method Count",
		Category:    signals.CategoryEntropy,
		Scope:       signals.ScopeClass,
		Description: "Number of methods with entropy above threshold",
		Unit:        "count",
		LowerBetter: true,
	}
}

// Dependencies returns names of signals this signal depends on
func (s *HighEntropyMethodCountSignal) Dependencies() []string {
	return []string{"MethodEntropy"}
}

// ComputeClass computes high entropy method count for a class
func (s *HighEntropyMethodCountSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
