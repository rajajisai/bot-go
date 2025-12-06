package signals

import (
	"time"

	"bot-go/internal/model/ast"
)

// SignalResult holds the computed value of a signal
type SignalResult struct {
	SignalName string         // Name of the signal
	Value      float64        // Computed value
	RawValue   interface{}    // Original value before normalization
	Normalized float64        // Value normalized to 0-1 range
	Metadata   map[string]any // Additional context (e.g., breakdown)
	ComputedAt time.Time      // When the result was computed
	Error      error          // Non-nil if computation failed
}

// SignalResultSet holds multiple signal results for an entity
type SignalResultSet struct {
	EntityType string                  // "class", "method", "file"
	EntityID   ast.NodeID              // Node ID of the entity
	EntityName string                  // Name of the entity
	FilePath   string                  // File path containing the entity
	Results    map[string]SignalResult // signal name -> result
	ComputedAt time.Time               // When the result set was computed
}

// NewSignalResult creates a successful result
func NewSignalResult(name string, value float64) SignalResult {
	return SignalResult{
		SignalName: name,
		Value:      value,
		RawValue:   value,
		ComputedAt: time.Now(),
	}
}

// NewSignalResultWithMetadata creates a result with additional metadata
func NewSignalResultWithMetadata(name string, value float64, metadata map[string]any) SignalResult {
	return SignalResult{
		SignalName: name,
		Value:      value,
		RawValue:   value,
		Metadata:   metadata,
		ComputedAt: time.Now(),
	}
}

// NewSignalResultError creates a failed result
func NewSignalResultError(name string, err error) SignalResult {
	return SignalResult{
		SignalName: name,
		Error:      err,
		ComputedAt: time.Now(),
	}
}

// IsValid returns true if computation succeeded
func (r SignalResult) IsValid() bool {
	return r.Error == nil
}

// ExceedsThreshold checks if value exceeds the given threshold
// For lowerBetter=true, returns true if value > threshold (bad)
// For lowerBetter=false, returns true if value < threshold (bad)
func (r SignalResult) ExceedsThreshold(threshold float64, lowerBetter bool) bool {
	if !r.IsValid() {
		return false
	}
	if lowerBetter {
		return r.Value > threshold
	}
	return r.Value < threshold
}

// WithNormalized returns a copy of the result with normalized value set
func (r SignalResult) WithNormalized(normalized float64) SignalResult {
	r.Normalized = normalized
	return r
}

// WithMetadata returns a copy of the result with metadata added
func (r SignalResult) WithMetadata(key string, value any) SignalResult {
	if r.Metadata == nil {
		r.Metadata = make(map[string]any)
	}
	r.Metadata[key] = value
	return r
}

// NewSignalResultSet creates a new result set for an entity
func NewSignalResultSet(entityType string, entityID ast.NodeID, entityName string, filePath string) *SignalResultSet {
	return &SignalResultSet{
		EntityType: entityType,
		EntityID:   entityID,
		EntityName: entityName,
		FilePath:   filePath,
		Results:    make(map[string]SignalResult),
		ComputedAt: time.Now(),
	}
}

// AddResult adds a signal result to the set
func (s *SignalResultSet) AddResult(result SignalResult) {
	if s.Results == nil {
		s.Results = make(map[string]SignalResult)
	}
	s.Results[result.SignalName] = result
}

// GetResult retrieves a result by signal name
func (s *SignalResultSet) GetResult(signalName string) (SignalResult, bool) {
	if s.Results == nil {
		return SignalResult{}, false
	}
	result, ok := s.Results[signalName]
	return result, ok
}

// HasResult checks if a result exists for a signal
func (s *SignalResultSet) HasResult(signalName string) bool {
	if s.Results == nil {
		return false
	}
	_, ok := s.Results[signalName]
	return ok
}

// GetValidResults returns only results without errors
func (s *SignalResultSet) GetValidResults() map[string]SignalResult {
	valid := make(map[string]SignalResult)
	for name, result := range s.Results {
		if result.IsValid() {
			valid[name] = result
		}
	}
	return valid
}

// GetResultNames returns all signal names in the set
func (s *SignalResultSet) GetResultNames() []string {
	if s.Results == nil {
		return nil
	}
	names := make([]string, 0, len(s.Results))
	for name := range s.Results {
		names = append(names, name)
	}
	return names
}

// Merge combines another result set into this one
// Results from other will overwrite existing results with the same name
func (s *SignalResultSet) Merge(other *SignalResultSet) {
	if other == nil || other.Results == nil {
		return
	}
	if s.Results == nil {
		s.Results = make(map[string]SignalResult)
	}
	for name, result := range other.Results {
		s.Results[name] = result
	}
}

// Size returns the number of results in the set
func (s *SignalResultSet) Size() int {
	if s.Results == nil {
		return 0
	}
	return len(s.Results)
}

// GetErrorResults returns only results with errors
func (s *SignalResultSet) GetErrorResults() map[string]SignalResult {
	errors := make(map[string]SignalResult)
	for name, result := range s.Results {
		if !result.IsValid() {
			errors[name] = result
		}
	}
	return errors
}

// GetValue returns the value for a signal, or 0 if not found
func (s *SignalResultSet) GetValue(signalName string) float64 {
	if result, ok := s.GetResult(signalName); ok && result.IsValid() {
		return result.Value
	}
	return 0
}

// GetNormalizedValue returns the normalized value for a signal, or 0 if not found
func (s *SignalResultSet) GetNormalizedValue(signalName string) float64 {
	if result, ok := s.GetResult(signalName); ok && result.IsValid() {
		return result.Normalized
	}
	return 0
}
