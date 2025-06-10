package emailservices

import (
	"github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/interfaces"
)

// EmailService is a service that handles email operations.
type EmailService struct {
	Validation  interfaces.EmailValidator
	Composer    interfaces.EmailComposer
	Sender      interfaces.EmailSender
	Logger      interfaces.EmailLogger
	Attachments interfaces.EmailAttachment
}

func NewEmailService(
	validation interfaces.EmailValidator,
	composer interfaces.EmailComposer,
	sender interfaces.EmailSender,
	logger interfaces.EmailLogger,
	attachments interfaces.EmailAttachment,
) *EmailService {
	return &EmailService{
		Validation:  validation,
		Composer:    composer,
		Sender:      sender,
		Logger:      logger,
		Attachments: attachments,
	}
}

// composeEmail composes an email message.
func (s *EmailService) composeEmail(subject string, body string, to []string) (string, error) {
	email, err := s.Composer.Compose(subject, body, to)
	if err != nil {
		if logErr := s.Logger.Log("Failed to compose email: "+err.Error(), "error"); logErr != nil {
			return "", logErr
		}
		return "", err
	}
	return email, nil
}

// validateEmail validates an email address.
func (s *EmailService) validate(
	subject, body string,
	to []string,
) error {
	// Validate email addresses
	if !s.Validation.ValidateRecipient(to) {
		return s.Logger.Log("Invalid email recipient(s)", "error")
	}
	if !s.Validation.ValidateSubject(subject) {
		return s.Logger.Log("Invalid email subject", "error")
	}
	if !s.Validation.ValidateBody(body) {
		return s.Logger.Log("Invalid email body", "error")
	}
	return nil
}

// sendEmail sends an email message.
func (s *EmailService) sendEmail(email string, to []string) error {
	if err := s.Sender.Send(email, to); err != nil {
		if logErr := s.Logger.Log("Failed to send email: "+err.Error(), "error"); logErr != nil {
			return logErr
		}
		return err
	}
	if err := s.Logger.Log("Email sent successfully", "info"); err != nil {
		return err
	}
	return nil
}

// addAttachment adds an attachment to the email.
func (s *EmailService) addAttachment(filePath string) error {
	return s.Attachments.AddAttachment(filePath)
}

func (s *EmailService) Send(
	subject, body string,
	to []string,
	attachments []string,
) error {
	if err := s.validate(subject, body, to); err != nil {
		return err
	}

	email, err := s.composeEmail(subject, body, to)
	if err != nil {
		return err
	}

	if err := s.sendEmail(email, to); err != nil {
		return err
	}

	for _, attachment := range attachments {
		if err := s.addAttachment(attachment); err != nil {
			return err
		}
	}

	return nil
}
