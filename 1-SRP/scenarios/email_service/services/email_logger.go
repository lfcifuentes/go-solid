package services

import (
	"github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/interfaces"
)

type EmailLogger struct{}

func NewEmailLogger() interfaces.EmailLogger {
	return &EmailLogger{}
}

func (l *EmailLogger) Log(message, level string) error {
	return nil
}
