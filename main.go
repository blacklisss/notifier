package main

import (
	"github.com/sirupsen/logrus"
	"notificationService/configuration"
	"notificationService/services"
)

func main() {
	log := logrus.New()

	config, err := configuration.Load("configuration/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	logger := services.NewLogger(log)

	notificator := services.NewNotificator(config, logger)
	notificator.Run()
	//fmt.Printf("%#v", *config)
}
