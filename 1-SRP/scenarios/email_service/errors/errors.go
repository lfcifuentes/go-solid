package emailerrors

import "errors"

var (
	ErrInvalidEmailParameters = errors.New("invalid email parameters")
	ErrTooManyRecipients      = errors.New("too many recipients")
)
