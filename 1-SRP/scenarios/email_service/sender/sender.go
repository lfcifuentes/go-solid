package sender

type EmailSender struct{}

func NewEmailSender() *EmailSender {
	return &EmailSender{}
}
