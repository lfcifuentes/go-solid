package main

import (
	"log"

	"github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/config"
	"github.com/lfcifuentes/go-solid/1-SRP/scenarios/email_service/services"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	services := services.NewEmailService(
		services.NewEmailValidator(),
		services.NewEmailComposer(),
		services.NewEmailSender(),
		services.NewEmailLogger(),
		services.NewEmailAttachment(),
	)

	services.Send(
		"Test Subject",
		"Test Body",
		[]string{"recipient@example.com"},
		[]string{"attachment1.txt", "attachment2.txt"},
	)
}
