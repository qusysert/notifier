package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"notifier/internal/app/service/queueservice"
	"notifier/internal/pkg/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	connectRabbitMQ, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.QueueUser, cfg.QueuePassword, cfg.QueueHost, cfg.QueuePort))
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	queueSrv := queueservice.New(connectRabbitMQ, channelRabbitMQ)

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
