package queueservice

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func (s QueueService) NewConsumer(name string) (<-chan amqp.Delivery, error) {
	amqpServerURL := fmt.Sprintf("amqp://%s:%s@%s:%d/", s.Connection.User, s.Connection.Password, s.Connection.Host, s.Connection.Port)
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// Subscribing to QueueService1 for getting messages.
	messages, err := channelRabbitMQ.Consume(
		name,  // queue name
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // arguments
	)
	if err != nil {
		log.Println(err)
	}

	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	return messages, nil
}
