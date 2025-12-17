package util

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"bot-go/internal/config"
)

// GitAnalyzer defines the interface for git history analysis
type GitAnalyzer interface {
	// GetRepoPath returns the repository path
	GetRepoPath() string

	// GetCoChangedClasses returns classes that frequently change together
	GetCoChangedClasses(ctx context.Context, classPath string, lookbackCommits int) ([]CoChangeInfo, error)

	// GetCoChangedMethods returns methods that frequently change together
	GetCoChangedMethods(ctx context.Context, methodPath string, lookbackCommits int) ([]CoChangeInfo, error)

	// GetFileChangeHistory returns the change history for a file
	GetFileChangeHistory(ctx context.Context, filePath string, lookbackCommits int) ([]ChangeInfo, error)

	// GetCoChangedFiles returns files that frequently change together
	GetCoChangedFiles(ctx context.Context, filePath string, lookbackCommits int) ([]CoChangeInfo, error)
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

// NewGitAnalyzer creates a new GitAnalyzer based on configuration
// Currently only supports "ondemand" mode; "precompute" mode is not yet implemented
// Returns an error if:
//   - cfg is nil (git_analysis config section is missing)
//   - cfg.Enabled is false
//   - cfg.Mode is invalid or unsupported
func NewGitAnalyzer(repoPath string, cfg *config.GitAnalysisConfig) (GitAnalyzer, error) {
	if cfg == nil {
		return nil, fmt.Errorf("git analysis configuration is required: add 'git_analysis' section to app.yaml")
	}

	if !cfg.Enabled {
		return nil, fmt.Errorf("git analysis is disabled in configuration: set 'git_analysis.enabled: true' to enable")
	}

	if repoPath == "" {
		return nil, fmt.Errorf("repository path is required for git analysis")
	}

	lookback := cfg.LookbackCommits
	if lookback <= 0 {
		lookback = 1000 // default
	}

	switch cfg.Mode {
	case config.GitAnalysisModeOnDemand:
		return NewOnDemandGitAnalyzer(repoPath, lookback), nil
	case "":
		return nil, fmt.Errorf("git analysis mode is required: set 'git_analysis.mode' to 'ondemand' or 'precompute'")
	case config.GitAnalysisModePrecompute:
		return nil, fmt.Errorf("precompute mode for git analysis is not yet implemented")
	default:
		return nil, fmt.Errorf("unknown git analysis mode: %s (valid modes: 'ondemand', 'precompute')", cfg.Mode)
	}
}

// OnDemandGitAnalyzer executes git commands on-demand when methods are called
type OnDemandGitAnalyzer struct {
	repoPath        string
	lookbackCommits int
}

// NewOnDemandGitAnalyzer creates a new on-demand git analyzer
func NewOnDemandGitAnalyzer(repoPath string, lookbackCommits int) *OnDemandGitAnalyzer {
	if lookbackCommits <= 0 {
		lookbackCommits = 1000
	}
	return &OnDemandGitAnalyzer{
		repoPath:        repoPath,
		lookbackCommits: lookbackCommits,
	}
}

// GetRepoPath returns the repository path
func (g *OnDemandGitAnalyzer) GetRepoPath() string {
	return g.repoPath
}

// GetCoChangedClasses returns classes that frequently change together
// For now, this delegates to GetCoChangedFiles since class-level tracking
// requires AST analysis of diffs
func (g *OnDemandGitAnalyzer) GetCoChangedClasses(ctx context.Context, classPath string, lookbackCommits int) ([]CoChangeInfo, error) {
	// TODO: Implement class-level co-change analysis
	// This would require:
	// 1. Getting commits that modified the file containing the class
	// 2. For each commit, parsing the diff to identify which classes changed
	// 3. Building co-change frequency matrix
	return nil, nil
}

// GetCoChangedMethods returns methods that frequently change together
func (g *OnDemandGitAnalyzer) GetCoChangedMethods(ctx context.Context, methodPath string, lookbackCommits int) ([]CoChangeInfo, error) {
	// TODO: Implement method-level co-change analysis
	// This would require AST diffing to identify method-level changes
	return nil, nil
}

// GetFileChangeHistory returns the change history for a file
func (g *OnDemandGitAnalyzer) GetFileChangeHistory(ctx context.Context, filePath string, lookbackCommits int) ([]ChangeInfo, error) {
	// TODO: Implement using git log
	// git log --follow -n {lookbackCommits} --pretty=format:"%H|%an|%ad|%s" --numstat -- {filePath}
	return nil, nil
}

// GetCoChangedFiles returns files that frequently change together with the given file
func (g *OnDemandGitAnalyzer) GetCoChangedFiles(ctx context.Context, filePath string, lookbackCommits int) ([]CoChangeInfo, error) {
	if lookbackCommits <= 0 {
		lookbackCommits = g.lookbackCommits
	}

	// Get the relative path from the repo root
	relPath, err := g.getRelativePath(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get relative path: %w", err)
	}

	// Step 1: Get commits that touched this file
	commits, err := g.getCommitsForFile(ctx, relPath, lookbackCommits)
	if err != nil {
		return nil, fmt.Errorf("failed to get commits for file: %w", err)
	}

	if len(commits) == 0 {
		return nil, nil
	}

	// Step 2: For each commit, get all files changed in that commit
	// Track: file -> list of commits where it changed together with target file
	coChanges := make(map[string][]string)

	for _, commit := range commits {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		files, err := g.getFilesInCommit(ctx, commit)
		if err != nil {
			// Log error but continue with other commits
			continue
		}

		for _, file := range files {
			// Skip the target file itself
			if file == relPath {
				continue
			}
			coChanges[file] = append(coChanges[file], commit)
		}
	}

	// Step 3: Convert to CoChangeInfo slice and sort by frequency
	results := make([]CoChangeInfo, 0, len(coChanges))
	for entityPath, commits := range coChanges {
		results = append(results, CoChangeInfo{
			EntityPath: entityPath,
			Frequency:  len(commits),
			Commits:    commits,
		})
	}

	// Sort by frequency (descending)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Frequency > results[j].Frequency
	})

	return results, nil
}

// getRelativePath converts an absolute or relative file path to a path relative to repo root
func (g *OnDemandGitAnalyzer) getRelativePath(filePath string) (string, error) {
	// If the path is already relative, use it as-is
	if !filepath.IsAbs(filePath) {
		return filePath, nil
	}

	// Get git root directory
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = g.repoPath
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get git root: %w", err)
	}
	gitRoot := strings.TrimSpace(string(output))

	// Calculate relative path
	relPath, err := filepath.Rel(gitRoot, filePath)
	if err != nil {
		return "", err
	}

	return relPath, nil
}

// getCommitsForFile returns commit hashes that modified the given file
func (g *OnDemandGitAnalyzer) getCommitsForFile(ctx context.Context, relPath string, limit int) ([]string, error) {
	cmd := exec.CommandContext(ctx, "git", "log", "--follow",
		fmt.Sprintf("-n%d", limit),
		"--pretty=format:%H",
		"--", relPath)
	cmd.Dir = g.repoPath

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	outputStr := strings.TrimSpace(string(output))
	if outputStr == "" {
		return nil, nil
	}

	return strings.Split(outputStr, "\n"), nil
}

// getFilesInCommit returns all files changed in a given commit
func (g *OnDemandGitAnalyzer) getFilesInCommit(ctx context.Context, commitHash string) ([]string, error) {
	cmd := exec.CommandContext(ctx, "git", "diff-tree", "--no-commit-id", "--name-only", "-r", commitHash)
	cmd.Dir = g.repoPath

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	outputStr := strings.TrimSpace(string(output))
	if outputStr == "" {
		return nil, nil
	}

	return strings.Split(outputStr, "\n"), nil
}

// Ensure OnDemandGitAnalyzer implements GitAnalyzer
var _ GitAnalyzer = (*OnDemandGitAnalyzer)(nil)
