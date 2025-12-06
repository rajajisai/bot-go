package util

// NormalizationRange defines min/max for normalization
type NormalizationRange struct {
	Min float64
	Max float64
}

// Normalizer normalizes signal values to 0-1 range
type Normalizer struct {
	// Configured ranges per signal
	ranges map[string]NormalizationRange
}

// NewNormalizer creates a new normalizer with default ranges
func NewNormalizer() *Normalizer {
	return &Normalizer{
		ranges: DefaultRanges(),
	}
}

// Configure sets the normalization range for a signal
func (n *Normalizer) Configure(signalName string, min, max float64) {
}

// GetRange returns the configured range for a signal
func (n *Normalizer) GetRange(signalName string) (NormalizationRange, bool) {
	return NormalizationRange{}, false
}

// Normalize normalizes a value to 0-1 range
func (n *Normalizer) Normalize(signalName string, value float64) float64 {
	return 0
}

// NormalizeInverse normalizes where lower is better (inverts the scale)
func (n *Normalizer) NormalizeInverse(signalName string, value float64) float64 {
	return 0
}

// NormalizeWithRange normalizes using explicit min/max values
func (n *Normalizer) NormalizeWithRange(value, min, max float64) float64 {
	return 0
}

// NormalizeInverseWithRange normalizes inversely using explicit min/max
func (n *Normalizer) NormalizeInverseWithRange(value, min, max float64) float64 {
	return 0
}

// DefaultRanges returns default normalization ranges for common signals
func DefaultRanges() map[string]NormalizationRange {
	return map[string]NormalizationRange{
		// Size signals
		"LOC":     {Min: 0, Max: 500},
		"LOCNAMM": {Min: 0, Max: 400},
		"NOM":     {Min: 0, Max: 50},
		"NOMNAMM": {Min: 0, Max: 40},
		"NOF":     {Min: 0, Max: 30},
		"NOPA":    {Min: 0, Max: 20},
		"NOAM":    {Min: 0, Max: 20},

		// Complexity signals
		"CYCLO":      {Min: 1, Max: 50},
		"WMC":        {Min: 0, Max: 100},
		"WMCNAMM":    {Min: 0, Max: 80},
		"MAXNESTING": {Min: 0, Max: 10},
		"NOLV":       {Min: 0, Max: 30},
		"AMC":        {Min: 1, Max: 20},

		// Cohesion signals (ratios, already 0-1)
		"TCC":   {Min: 0, Max: 1},
		"LCC":   {Min: 0, Max: 1},
		"LCOM":  {Min: 0, Max: 100},
		"LCOM4": {Min: 1, Max: 10},
		"ATLD":  {Min: 0, Max: 20},

		// Coupling signals
		"ATFD":   {Min: 0, Max: 20},
		"FANOUT": {Min: 0, Max: 30},
		"CINT":   {Min: 0, Max: 50},
		"CDISP":  {Min: 0, Max: 1},
		"FDP":    {Min: 0, Max: 15},
		"LAA":    {Min: 0, Max: 1},
		"CBO":    {Min: 0, Max: 30},
		"RFC":    {Min: 0, Max: 100},

		// Message chain signals
		"MaMCL": {Min: 1, Max: 10},
		"MeMCL": {Min: 1, Max: 5},
		"NMCS":  {Min: 0, Max: 20},

		// Change signals
		"CC": {Min: 0, Max: 20},
		"CM": {Min: 0, Max: 30},

		// Entropy signals
		"FileEntropy":   {Min: 0, Max: 10},
		"MethodEntropy": {Min: 0, Max: 10},
		"ClassEntropy":  {Min: 0, Max: 10},
		"ZScore":        {Min: -3, Max: 3},

		// Composite signals
		"WOC": {Min: 0, Max: 1},
	}
}

// Thresholds contains default threshold values for code smell detection
type Thresholds struct {
	values map[string]float64
}

// NewThresholds creates a new thresholds instance with defaults
func NewThresholds() *Thresholds {
	return &Thresholds{
		values: DefaultThresholds(),
	}
}

// Get returns the threshold for a signal
func (t *Thresholds) Get(signalName string) (float64, bool) {
	return 0, false
}

// Set sets a threshold value
func (t *Thresholds) Set(signalName string, value float64) {
}

// DefaultThresholds returns default thresholds for code smell detection
func DefaultThresholds() map[string]float64 {
	return map[string]float64{
		// God Class thresholds (Lanza & Marinescu)
		"LOCNAMM": 176,
		"WMCNAMM": 22,
		"NOMNAMM": 18,
		"TCC":     0.33,  // Lower bound (less than this is bad)
		"ATFD":    6,

		// Data Class thresholds
		"WOC":  0.33,  // Lower bound
		"NOPA": 5,
		"NOAM": 5,

		// Long Method thresholds
		"LOC_METHOD":   50,
		"CYCLO_METHOD": 10,
		"NOLV_METHOD":  10,

		// Feature Envy thresholds
		"LAA": 0.33,  // Lower bound
		"FDP": 5,

		// Brain Method thresholds
		"MAXNESTING_BRAIN": 5,
		"CYCLO_BRAIN":      20,
		"LOC_BRAIN":        100,
	}
}
