package composeremailservice

import (
	"strings"

	"github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/compose"
	emailerrors "github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/errors"
	"github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/interfaces"
)

type EmailComposer struct {
	config compose.EmailConfig
}

func NewEmailComposer() interfaces.EmailComposer {
	return &EmailComposer{
		config: compose.NewComposeConfig(),
	}
}

func (ec *EmailComposer) Compose(subject, body string, to []string) (string, error) {
	if subject == "" || body == "" || len(to) == 0 {
		return "", emailerrors.ErrInvalidEmailParameters
	}

	if len(to) > ec.config.MaxToRecipients {
		return "", emailerrors.ErrTooManyRecipients
	}
	emailMessage := strings.Builder{}
	emailMessage.WriteString("From: " + ec.config.From + "\n")
	emailMessage.WriteString("To: " + strings.Join(to, ", ") + "\n")
	emailMessage.WriteString("Subject: " + subject + "\n")
	emailMessage.WriteString("Content-Type: text/plain; charset=" + ec.config.CharsetEncoding + "\n\n")
	emailMessage.WriteString(body)

	return emailMessage.String(), nil

}
