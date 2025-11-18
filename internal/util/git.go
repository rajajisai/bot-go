package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GitInfo contains git repository information
type GitInfo struct {
	HeadCommitSHA  string
	HeadCommitMsg  string
	ModifiedFiles  map[string]bool // Set of files modified compared to HEAD
	IsGitRepo      bool
}

// GetGitInfo retrieves git information for a repository path
func GetGitInfo(repoPath string) (*GitInfo, error) {
	info := &GitInfo{
		ModifiedFiles: make(map[string]bool),
	}

	// Check if this is a git repository
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	cmd.Dir = repoPath
	if err := cmd.Run(); err != nil {
		info.IsGitRepo = false
		return info, nil
	}
	info.IsGitRepo = true

	// Get HEAD commit SHA
	cmd = exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = repoPath
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD commit SHA: %w", err)
	}
	info.HeadCommitSHA = strings.TrimSpace(string(output))

	// Get HEAD commit message (first line)
	cmd = exec.Command("git", "log", "-1", "--pretty=%s")
	cmd.Dir = repoPath
	output, err = cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD commit message: %w", err)
	}
	info.HeadCommitMsg = strings.TrimSpace(string(output))

	// Get modified files (compared to HEAD)
	// This includes: modified, added, deleted files in working directory and index
	cmd = exec.Command("git", "diff", "--name-only", "HEAD")
	cmd.Dir = repoPath
	output, err = cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get modified files: %w", err)
	}

	modifiedFiles := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, file := range modifiedFiles {
		if file != "" {
			// Convert to absolute path
			absPath := filepath.Join(repoPath, file)
			info.ModifiedFiles[absPath] = true
		}
	}

	return info, nil
}

// GetFileContentFromGit retrieves file content from git HEAD
// Returns error if file is not tracked by git
func GetFileContentFromGit(repoPath, filePath string) ([]byte, error) {
	// Get relative path from repo root
	relPath, err := filepath.Rel(repoPath, filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get relative path: %w", err)
	}

	// Use git show to get file content from HEAD
	cmd := exec.Command("git", "show", fmt.Sprintf("HEAD:%s", relPath))
	cmd.Dir = repoPath
	output, err := cmd.Output()
	if err != nil {
		// Check if it's because the file doesn't exist in git
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 128 {
			return nil, fmt.Errorf("file not tracked by git: %s", relPath)
		}
		return nil, fmt.Errorf("failed to get file content from git: %w", err)
	}

	return output, nil
}

// IsFileModified checks if a file is modified compared to HEAD
func IsFileModified(gitInfo *GitInfo, filePath string) bool {
	if gitInfo == nil || !gitInfo.IsGitRepo {
		return false
	}
	return gitInfo.ModifiedFiles[filePath]
}

// ReadFileOptimized reads file content, using git HEAD if useHead is true and file is unmodified
// In HEAD mode, untracked files are skipped (returns nil content with error)
func ReadFileOptimized(repoPath, filePath string, useHead bool, gitInfo *GitInfo) ([]byte, error) {
	// If not using HEAD mode, read from disk
	if !useHead || gitInfo == nil || !gitInfo.IsGitRepo {
		return os.ReadFile(filePath)
	}

	// If file is modified compared to HEAD, read from disk
	if IsFileModified(gitInfo, filePath) {
		return os.ReadFile(filePath)
	}

	// File is unmodified according to git diff, try to read from git HEAD
	content, err := GetFileContentFromGit(repoPath, filePath)
	if err != nil {
		// If file is not tracked by git (e.g., in .gitignore or new untracked file),
		// return error to skip processing
		if strings.Contains(err.Error(), "file not tracked by git") {
			return nil, err
		}
		// For other git errors, fall back to reading from disk
		return os.ReadFile(filePath)
	}

	return content, nil
}

// GetLastCommitForFile gets the commit SHA of the last commit that modified a file
func GetLastCommitForFile(repoPath, filePath string) (string, error) {
	// Get relative path from repo root
	relPath, err := filepath.Rel(repoPath, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to get relative path: %w", err)
	}

	// Get the last commit SHA for this file
	cmd := exec.Command("git", "log", "-1", "--pretty=%H", "--", relPath)
	cmd.Dir = repoPath
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get last commit for file: %w", err)
	}

	commitSHA := strings.TrimSpace(string(output))
	if commitSHA == "" {
		return "", fmt.Errorf("no commits found for file: %s", relPath)
	}

	return commitSHA, nil
}

// CalculateFileSHA256 calculates the SHA256 hash of file content
func CalculateFileSHA256(content []byte) string {
	hash := sha256.Sum256(content)
	return hex.EncodeToString(hash[:])
}

// GetRelativePath returns the relative path of a file from the repository root
func GetRelativePath(repoPath, filePath string) (string, error) {
	relPath, err := filepath.Rel(repoPath, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to get relative path: %w", err)
	}
	return relPath, nil
}
