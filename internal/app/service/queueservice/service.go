package queueservice

import (
	"github.com/streadway/amqp"
)

type QueueService struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func New(conn *amqp.Connection, channel *amqp.Channel) *QueueService {
	return &QueueService{Connection: conn, Channel: channel}
}
