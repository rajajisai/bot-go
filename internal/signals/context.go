package signals

import (
	"bot-go/internal/model/ast"
	"bot-go/internal/service/codegraph"
	"bot-go/internal/service/ngram"
	"bot-go/internal/service/vector"

	"go.uber.org/zap"
)

// SignalContext provides shared resources for signal computation
type SignalContext struct {
	// Graph database access
	CodeGraph *codegraph.CodeGraph

	// N-gram service for entropy calculations
	NGramService *ngram.NGramService

	// Vector database for embeddings (optional)
	VectorDB vector.VectorDatabase

	// Repository information
	RepoName string
	RepoPath string

	// Caching
	Cache *SignalCache

	// Logger
	Logger *zap.Logger
}

// SignalCache caches computed signal values
type SignalCache struct {
	classResults  map[cacheKey]SignalResult
	methodResults map[cacheKey]SignalResult
	fileResults   map[cacheKey]SignalResult
}

// cacheKey is a composite key for cache lookups
type cacheKey struct {
	nodeID     ast.NodeID
	signalName string
}

// NewSignalContext creates a new signal computation context
func NewSignalContext(
	codeGraph *codegraph.CodeGraph,
	ngramService *ngram.NGramService,
	vectorDB vector.VectorDatabase,
	repoName string,
	repoPath string,
	logger *zap.Logger,
) *SignalContext {
	return &SignalContext{
		CodeGraph:    codeGraph,
		NGramService: ngramService,
		VectorDB:     vectorDB,
		RepoName:     repoName,
		RepoPath:     repoPath,
		Cache:        NewSignalCache(),
		Logger:       logger,
	}
}

// NewSignalCache creates a new signal cache
func NewSignalCache() *SignalCache {
	return &SignalCache{
		classResults:  make(map[cacheKey]SignalResult),
		methodResults: make(map[cacheKey]SignalResult),
		fileResults:   make(map[cacheKey]SignalResult),
	}
}

// GetClassResult retrieves cached class signal result
func (c *SignalCache) GetClassResult(classID ast.NodeID, signalName string) (SignalResult, bool) {
	return SignalResult{}, false
}

// SetClassResult caches a class signal result
func (c *SignalCache) SetClassResult(classID ast.NodeID, signalName string, result SignalResult) {
}

// GetMethodResult retrieves cached method signal result
func (c *SignalCache) GetMethodResult(methodID ast.NodeID, signalName string) (SignalResult, bool) {
	return SignalResult{}, false
}

// SetMethodResult caches a method signal result
func (c *SignalCache) SetMethodResult(methodID ast.NodeID, signalName string, result SignalResult) {
}

// GetFileResult retrieves cached file signal result
func (c *SignalCache) GetFileResult(fileID ast.NodeID, signalName string) (SignalResult, bool) {
	return SignalResult{}, false
}

// SetFileResult caches a file signal result
func (c *SignalCache) SetFileResult(fileID ast.NodeID, signalName string, result SignalResult) {
}

// Clear clears all cached values
func (c *SignalCache) Clear() {
}

// ClearClass clears cached values for a specific class
func (c *SignalCache) ClearClass(classID ast.NodeID) {
}

// ClearMethod clears cached values for a specific method
func (c *SignalCache) ClearMethod(methodID ast.NodeID) {
}

// ClearFile clears cached values for a specific file
func (c *SignalCache) ClearFile(fileID ast.NodeID) {
}

// Size returns the total number of cached entries
func (c *SignalCache) Size() int {
	return 0
}
