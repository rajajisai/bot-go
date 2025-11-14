package util

import (
	"net/url"
	"path/filepath"
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
