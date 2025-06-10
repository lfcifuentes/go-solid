package validator

import (
	"errors"
	"regexp"
)

// Validator provides methods to validate email addresses and subjects.
type Validator struct {
	config ValidationConfig
}

// NewValidator creates a new instance of Validator.
func NewValidator() *Validator {
	return &Validator{
		config: ValidationConfig{
			MaxEmailLength:   DefaultMaxEmailLength,
			MaxSubjectLength: DefaultMaxSubjectLength,
			EmailPattern:     DefaultEmailPattern,
		},
	}
}

// errorResult creates a ValidationResult indicating a validation error.
func (v *Validator) errorResult(field, message string) ValidationResult {
	return ValidationResult{
		IsValid: false,
		Errors: []ValidationError{
			{
				Field:   field,
				Message: message,
			},
		},
	}
}

// successResult creates a ValidationResult indicating successful validation.
func (v *Validator) successResult() ValidationResult {
	return ValidationResult{
		IsValid: true,
		Errors:  nil,
	}
}

// validateStringLength checks if the given string exceeds the specified maximum length.
func (v *Validator) validateStringLength(s string, maxLength int) error {
	// Check if the string exceeds the maximum length
	if len(s) > maxLength {
		return errors.New("string exceeds maximum length")
	}

	return nil
}

// ValidateEmail checks if the provided email address is valid according to the configured pattern and length.
func (v *Validator) ValidateEmail(email string) ValidationResult {
	// A simple regex for email validation
	matched, err := regexp.MatchString(
		v.config.EmailPattern,
		email,
	)

	if err != nil {
		return v.errorResult("email", "Error occurred while validating email")
	}

	if !matched {
		return v.errorResult("email", "Invalid email format")
	}

	// Check if the email exceeds the maximum length
	if err := v.validateStringLength(email, v.config.MaxEmailLength); err != nil {
		return v.errorResult("email", err.Error())
	}

	return v.successResult()
}

// ValidateStringNotEmpty checks if the provided string is not empty.
func (v *Validator) ValidateStringNotEmpty(s string) ValidationResult {
	// Subject should not be empty
	if s == "" {
		return v.errorResult("string", "String should not be empty")
	}

	return v.successResult()
}

// ValidateSubjectLength checks if the provided subject exceeds the maximum length.
func (v *Validator) ValidateSubjectLength(subject string) ValidationResult {
	// Check if the subject exceeds the maximum length
	if err := v.validateStringLength(subject, v.config.MaxSubjectLength); err != nil {
		return v.errorResult("subject", err.Error())
	}

	return v.successResult()
}
