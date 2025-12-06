package entropy

import (
	"context"

	"bot-go/internal/signals"
)

// FileEntropySignal computes entropy for a file's code
type FileEntropySignal struct{}

// NewFileEntropySignal creates a new FileEntropy signal
func NewFileEntropySignal() *FileEntropySignal {
	return &FileEntropySignal{}
}

// Metadata returns information about this signal
func (s *FileEntropySignal) Metadata() signals.SignalMetadata {
	return signals.SignalMetadata{
		Name:        "FileEntropy",
		FullName:    "File Entropy",
		Category:    signals.CategoryEntropy,
		Scope:       signals.ScopeFile,
		Description: "Cross-entropy of a file's tokens against the corpus model",
		Unit:        "bits",
		LowerBetter: true, // Lower entropy means more predictable/natural code
	}
}

// Dependencies returns names of signals this signal depends on
func (s *FileEntropySignal) Dependencies() []string {
	return nil
}

// ComputeFile computes entropy for a file
func (s *FileEntropySignal) ComputeFile(ctx context.Context, fileInfo *signals.FileInfo, sctx *signals.SignalContext) (signals.SignalResult, error) {
	return signals.SignalResult{}, nil
}
