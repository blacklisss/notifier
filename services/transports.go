package services

import (
	"notificationService/configuration"
	"notificationService/sender"
	"notificationService/transports"
	"reflect"
)

type Transporter interface {
	GetTransport(config *configuration.Config) []sender.ISender
}

type Transport struct {
	*Logger
}

func (t *Transport) GetTransport(config *configuration.Config) []sender.ISender {
	var tr []sender.ISender

	val := reflect.ValueOf(config.Transports)
	for i := 0; i < val.NumField(); i++ {
		switch val.Field(i).Type().Name() {
		case "TbotConfig":
			if config.Transports.Tbot.Active == true {
				tbot, err := transports.NewTgBot(config.Transports.Tbot.ApiKey)
				if err != nil {
					t.logger.Errorf("can't create Tbot Transport: %v\n", err)
				}

				tr = append(tr, tbot)
			}
		case "EmailConfig":

		}
	}

	return tr
}

func NewTransport(logger *Logger) *Transport {
	return &Transport{
		logger,
	}
}
