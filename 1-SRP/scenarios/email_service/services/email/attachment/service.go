package attachmentemailservice

import (
	"github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/interfaces"
)

type EmailAttachment struct{}

func NewEmailAttachment() interfaces.EmailAttachment {
	return &EmailAttachment{}
}

func (a *EmailAttachment) AddAttachment(filepath string) error {
	return nil
}
