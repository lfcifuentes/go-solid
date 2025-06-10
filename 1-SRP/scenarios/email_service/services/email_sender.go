package services

import "github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/interfaces"

type EmailSender struct{}

func NewEmailSender() interfaces.EmailSender {
	return &EmailSender{}
}
func (s *EmailSender) Send(email string, to []string) error {

	return nil
}
