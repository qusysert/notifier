package queueservice

import (
	"github.com/streadway/amqp"
	"log"
)

func (s QueueService) NewConsumer(name string) (<-chan amqp.Delivery, error) {

	// Subscribing to QueueService1 for getting messages.
	messages, err := s.Channel.Consume(
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
