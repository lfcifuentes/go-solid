package interfaces

// EmailValidator is an interface that validates email addresses.
type EmailValidator interface {
	Validate(email string) bool
	ValidateSubject(subject string) bool
	ValidateBody(body string) bool
	ValidateRecipient(recipients []string) bool
}

// EmailComposer is an interface that composes email messages.
type EmailComposer interface {
	Compose(subject string, body string, to []string) (string, error)
}

// EmailSender is an interface that sends email messages.
type EmailSender interface {
	Send(email string, to []string) error
}

// EmailLogger is an interface that logs email operations.
type EmailLogger interface {
	Log(message string, level string) error
}

// EmailAttachment is an interface that handles email attachments.
type EmailAttachment interface {
	AddAttachment(filePath string) error
}
