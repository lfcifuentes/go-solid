package validator

// ValidationResult represents the outcome of a validation operation
type ValidationResult struct {
	IsValid bool
	Errors  []ValidationError
}

// ValidationError represents a single validation error
type ValidationError struct {
	Field   string
	Message string
}

// ValidationConfig holds configuration for validation rules
type ValidationConfig struct {
	MaxEmailLength   int
	MaxSubjectLength int
	EmailPattern     string
}
