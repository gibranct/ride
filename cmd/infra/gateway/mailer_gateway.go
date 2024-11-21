package gateway

import "fmt"

type MailerGateway interface {
	Send(recipient, subject, message string)
}

type MailerGatewayMemory struct{}

func (m *MailerGatewayMemory) Send(recipient, subject, message string) {
	fmt.Printf("sending message: %s to %s with subject: %s", message, recipient, subject)
}

func NewMailerGatewayMemory() *MailerGatewayMemory {
	return &MailerGatewayMemory{}
}
