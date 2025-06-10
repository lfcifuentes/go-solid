package services

import (
	"github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/interfaces"
	"github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/validator"
)

type EmailValidator struct {
	validator *validator.Validator
}

func NewEmailValidator() interfaces.EmailValidator {
	return &EmailValidator{
		validator: validator.NewValidator(),
	}
}

// EmailValidator implements the EmailValidator interface for validating email addresses.
func (v *EmailValidator) Validate(email string) bool {
	return v.validator.ValidateEmail(email).IsValid
}

// ValidateSubject checks if the email subject is valid.
func (v *EmailValidator) ValidateSubject(subject string) bool {
	return v.validator.ValidateStringNotEmpty(subject).IsValid
}

// ValidateBody checks if the email body is valid.
func (v *EmailValidator) ValidateBody(body string) bool {
	return v.validator.ValidateStringNotEmpty(body).IsValid
}

// ValidateRecipient checks if the recipient email addresses are valid.
func (v *EmailValidator) ValidateRecipient(recipients []string) bool {
	for _, recipient := range recipients {
		if !v.validator.ValidateEmail(recipient).IsValid {
			return false
		}
	}
	return true
}
