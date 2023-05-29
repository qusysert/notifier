package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"notifier/internal/app/model"
	"notifier/internal/app/service/queueservice"
	"notifier/internal/app/service/tgservice"
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

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	queueSrv := queueservice.New(connectRabbitMQ, channelRabbitMQ)
	tgSrv := tgservice.New()

	messages, err := queueSrv.NewConsumer("Message")
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for message := range messages {
			var msg model.Message
			err := json.Unmarshal(message.Body, &msg)
			if err != nil {
				log.Fatal(err)
			}
			err = tgSrv.SendMessage(msg)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf(" > Sending message: %s\n", message.Body)
		}
	}()

	<-forever

}
