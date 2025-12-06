package cohesion

import (
	"context"

	"bot-go/internal/model/ast"
	"bot-go/internal/signals"
)

// TCCSignal computes Tight Class Cohesion
type TCCSignal struct{}

// NewTCCSignal creates a new TCC signal
func NewTCCSignal() *TCCSignal {
	return &TCCSignal{}
}

// Metadata returns information about this signal
func (s *TCCSignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "TCC",
		FullName:    "Tight Class Cohesion",
		Category:    signals.CategoryCohesion,
		Scope:       signals.ScopeClass,
		Description: "Ratio of directly connected method pairs to maximum possible pairs",
		Unit:        "ratio",
		LowerBetter: false, // Higher TCC means better cohesion
	}
}

// Dependencies returns names of signals this signal depends on
func (s *TCCSignal) Dependencies() []string {
	return nil
}

// ComputeClass computes TCC for a class
func (s *TCCSignal) ComputeClass(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}

// buildFieldAccessMatrix builds method-field access matrix for a class
func (s *TCCSignal) buildFieldAccessMatrix(ctx context.Context, classInfo *signals.ClassInfo, sctx *signals.SignalContext) (map[ast.NodeID][]ast.NodeID, error) {
	return nil, nil
}

// countConnectedPairs counts method pairs that share field access
func (s *TCCSignal) countConnectedPairs(accessMatrix map[ast.NodeID][]ast.NodeID) int {
	return 0
}
