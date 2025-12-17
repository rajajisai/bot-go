package signals

import "errors"

// Common errors for signal computation
var (
	// ErrNilInput indicates that a nil input was provided
	ErrNilInput = errors.New("nil input provided")

	// ErrNoData indicates that required data is not available
	ErrNoData = errors.New("required data not available")

	// ErrInvalidRange indicates an invalid source range
	ErrInvalidRange = errors.New("invalid source range")

	// ErrDependencyMissing indicates a required dependency signal is missing
	ErrDependencyMissing = errors.New("required dependency signal missing")
)
