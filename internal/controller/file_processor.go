package controller

import (
	"bot-go/internal/config"
	"context"
)

// FileProcessor defines the interface for processing individual files
// and performing repository-level post-processing operations
type FileProcessor interface {
	// ProcessFile processes a single file in the repository
	// filePath: absolute path to the file
	// content: file content read from disk (avoids multiple reads)
	// Returns an error if processing fails
	ProcessFile(ctx context.Context, repo *config.Repository, filePath string, content []byte) error

	// PostProcess performs repository-level operations after all files are processed
	// This is called once after all files have been processed
	PostProcess(ctx context.Context, repo *config.Repository) error

	// Name returns the name of this processor (for logging purposes)
	Name() string
}
