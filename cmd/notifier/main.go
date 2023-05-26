package main

import (
	"log"
	"notifier/internal/app/model"
	"notifier/internal/app/service/queueservice"
	"notifier/internal/pkg/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	queueSrv := queueservice.New(model.Conn{
		User:     cfg.QueueUser,
		Password: cfg.QueuePassword,
		Host:     cfg.QueueHost,
		Port:     cfg.QueuePort},
	)

	messages, err := queueSrv.NewConsumer("Email")
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for message := range messages {
			// For example, show received message in a console.
			log.Printf(" > Received message: %s\n", message.Body)
		}
	}()

	<-forever

}
