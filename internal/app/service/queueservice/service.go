package queueservice

import "notifier/internal/app/model"

type QueueService struct {
	Connection model.Conn
}

type IService interface {
}

func New(conn model.Conn) *QueueService {
	return &QueueService{Connection: conn}
}
