package validator

const (
	// Maximum length for the email body
	DefaultMaxBodyLength = 10000
	// Maximum length for the email address
	DefaultMaxEmailLength = 254
	// Maximum length for the email subject
	DefaultMaxSubjectLength = 998
	// Regex pattern for validating email addresses
	DefaultEmailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)
