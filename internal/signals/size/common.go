package size

import (
	"bot-go/pkg/lsp/base"
)

// calculateLOCFromRange calculates lines of code from a Range
// Returns end line - start line + 1 (inclusive count)
// Returns 0 if the range is invalid (end before start)
func calculateLOCFromRange(r base.Range) int {
	if r.End.Line < r.Start.Line {
		return 0
	}
	return r.End.Line - r.Start.Line + 1
}
