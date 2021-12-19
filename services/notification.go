package services

import (
	"fmt"
	"notificationService/configuration"
	"notificationService/sender"
)

type Notificator struct {
	sender *sender.Sender
}

func (n *Notificator) Run() {
	fmt.Println("running notificator service...")
	fmt.Printf("%#v", n)
}

func NewNotificator(config *configuration.Config, logger *Logger) *Notificator {
	transport := NewTransport(logger)
	transports := transport.GetTransport(config)

	sndr := sender.NewSender(transports)

	return &Notificator{sender: sndr}
}
