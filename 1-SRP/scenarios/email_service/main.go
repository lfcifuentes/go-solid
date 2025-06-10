package main

import (
	"log"

	"github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/config"
	emailservices "github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/services/email"
	attachmentemailservice "github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/services/email/attachment"
	composeremailservice "github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/services/email/composer"
	loggeremailservice "github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/services/email/logger"
	senderemailservice "github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/services/email/sender"
	validatoremailservice "github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/services/email/validator"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	service := emailservices.NewEmailService(
		validatoremailservice.NewEmailValidator(),
		composeremailservice.NewEmailComposer(),
		senderemailservice.NewEmailSender(),
		loggeremailservice.NewEmailLogger(),
		attachmentemailservice.NewEmailAttachment(),
	)

	service.Send(
		"Test Subject",
		"Test Body",
		[]string{"recipient@example.com"},
		[]string{"attachment1.txt", "attachment2.txt"},
	)
}
