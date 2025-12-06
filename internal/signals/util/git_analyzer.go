package util

import (
	"context"
)

// GitAnalyzer provides git history analysis
type GitAnalyzer struct {
	repoPath string
}

// NewGitAnalyzer creates a new GitAnalyzer
func NewGitAnalyzer(repoPath string) *GitAnalyzer {
	return &GitAnalyzer{
		repoPath: repoPath,
	}
}

// GetRepoPath returns the repository path
func (g *GitAnalyzer) GetRepoPath() string {
	return g.repoPath
}

// GetCoChangedClasses returns classes that frequently change together
func (g *GitAnalyzer) GetCoChangedClasses(ctx context.Context, classPath string, lookbackCommits int) ([]CoChangeInfo, error) {
	return nil, nil
}

// GetCoChangedMethods returns methods that frequently change together
func (g *GitAnalyzer) GetCoChangedMethods(ctx context.Context, methodPath string, lookbackCommits int) ([]CoChangeInfo, error) {
	return nil, nil
}

// GetFileChangeHistory returns the change history for a file
func (g *GitAnalyzer) GetFileChangeHistory(ctx context.Context, filePath string, lookbackCommits int) ([]ChangeInfo, error) {
	return nil, nil
}

// GetCoChangedFiles returns files that frequently change together
func (g *GitAnalyzer) GetCoChangedFiles(ctx context.Context, filePath string, lookbackCommits int) ([]CoChangeInfo, error) {
	return nil, nil
}

// CoChangeInfo represents co-change information
type CoChangeInfo struct {
	EntityPath string   // Path to the co-changed entity
	Frequency  int      // Number of times changed together
	Commits    []string // Commit hashes where they changed together
}

// ChangeInfo represents a single change
type ChangeInfo struct {
	CommitHash   string
	Author       string
	Date         string
	Message      string
	LinesAdded   int
	LinesRemoved int
}
