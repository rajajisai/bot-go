package util

import (
	"net/url"
	"path/filepath"
	"strings"
)

func ToUri(path, rootPath string) (string, error) {
	u, err := url.Parse(path)
	if err == nil && u.Scheme != "" {
		return path, nil
	}
	if filepath.IsAbs(path) {
		return "file://" + filepath.ToSlash(path), nil
	}
	absPath := filepath.Join(rootPath, path)
	return "file://" + absPath, nil
}

func ToRelativePath(rootPath, fullPath string) string {
	relPath, err := filepath.Rel(rootPath, fullPath)
	if err != nil {
		return fullPath
	}
	return relPath
}

func ExtractPathFromURI(uri string) string {
	if len(uri) > 7 && uri[:7] == "file://" {
		return uri[7:]
	}
	return uri
}

func Ptr[T any](v T) *T { return &v }

// ShouldSkipDirectory checks if a directory should be skipped during traversal
func ShouldSkipDirectory(path string) bool {
	skipDirs := []string{
		".git", "node_modules", ".vscode", ".idea", "vendor", "target",
		"build", "dist", "__pycache__", ".pytest_cache", "coverage",
		"site-packages", ".next", ".nuxt", ".cache", "tmp", "temp",
	}

	baseName := filepath.Base(path)
	for _, skipDir := range skipDirs {
		if baseName == skipDir {
			return true
		}
	}

	// Skip hidden directories (starting with .)
	if len(baseName) > 0 && baseName[0] == '.' && baseName != "." && baseName != ".." {
		return true
	}

	return false
}

// ShouldSkipFile checks if a file should be skipped during indexing
// This includes special files like Dockerfiles, lock files, build artifacts, etc.
func ShouldSkipFile(filePath string) bool {
	baseName := filepath.Base(filePath)

	// Skip specific file names (case-insensitive)
	skipFileNames := []string{
		"dockerfile",
		"docker-compose.yml",
		"docker-compose.yaml",
		".dockerignore",
		"makefile",
		".gitignore",
		".gitattributes",
		".editorconfig",
		".prettierrc",
		".eslintrc",
		".pylintrc",
		".flake8",
		"license",
		"readme.md",
		"readme.txt",
		"changelog.md",
		"contributing.md",
		"code_of_conduct.md",
		"security.md",
	}

	lowerBaseName := strings.ToLower(baseName)
	for _, skipName := range skipFileNames {
		if lowerBaseName == skipName {
			return true
		}
	}

	// Skip lock files
	lockFilePatterns := []string{
		"package-lock.json",
		"yarn.lock",
		"pnpm-lock.yaml",
		"composer.lock",
		"gemfile.lock",
		"pipfile.lock",
		"poetry.lock",
		"cargo.lock",
		"go.sum",
	}

	for _, pattern := range lockFilePatterns {
		if baseName == pattern {
			return true
		}
	}

	// Skip binary and compiled files
	binaryExtensions := []string{
		".exe", ".dll", ".so", ".dylib", ".a", ".o", ".obj",
		".pyc", ".pyo", ".class", ".jar", ".war",
		".wasm", ".bin",
	}

	ext := filepath.Ext(baseName)
	for _, binExt := range binaryExtensions {
		if ext == binExt {
			return true
		}
	}

	// Skip image, video, audio files
	mediaExtensions := []string{
		".png", ".jpg", ".jpeg", ".gif", ".bmp", ".ico", ".svg", ".webp",
		".mp4", ".avi", ".mov", ".wmv", ".flv", ".webm",
		".mp3", ".wav", ".ogg", ".flac", ".aac",
	}

	for _, mediaExt := range mediaExtensions {
		if ext == mediaExt {
			return true
		}
	}

	// Skip document and archive files
	docArchiveExtensions := []string{
		".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx",
		".zip", ".tar", ".gz", ".bz2", ".7z", ".rar", ".tgz",
	}

	for _, docExt := range docArchiveExtensions {
		if ext == docExt {
			return true
		}
	}

	// Skip log files and temporary files
	if ext == ".log" || ext == ".tmp" || ext == ".temp" || ext == ".swp" || ext == ".swo" {
		return true
	}

	// Skip files in build/output directories (check full path)
	skipPathPatterns := []string{
		"/node_modules/",
		"/vendor/",
		"/target/",
		"/build/",
		"/dist/",
		"/__pycache__/",
		"/.pytest_cache/",
		"/coverage/",
		"/site-packages/",
		"/.next/",
		"/.nuxt/",
		"/bin/",
		"/obj/",
		"/.git/",
	}

	normalizedPath := filepath.ToSlash(filepath.Clean(filePath))
	for _, pattern := range skipPathPatterns {
		if containsPath(normalizedPath, pattern) {
			return true
		}
	}

	return false
}

// containsPath checks if a normalized path contains a pattern
func containsPath(path, pattern string) bool {
	// Simple substring check for path patterns
	return len(path) > 0 && len(pattern) > 0 &&
		(path == pattern ||
		 path[:min(len(path), len(pattern))] == pattern ||
		 containsSubstring(path, pattern))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
