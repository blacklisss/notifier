package services

import "notificationService/sender"

type Notificator struct {
	sender sender.Sender
}

func NewNotificator(sender sender.Sender) {
	//..
}
