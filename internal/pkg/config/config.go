package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	QueueUser     string
	QueuePassword string
	QueueHost     string
	QueuePort     int
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
		panic(err)
	}

	queuePort, err := strconv.Atoi(os.Getenv("QUEUE_PORT"))
	if err != nil {
		return nil, fmt.Errorf("cant get queuePort: %v", err)
	}
	config := Config{
		QueueUser:     os.Getenv("QUEUE_USER"),
		QueuePassword: os.Getenv("QUEUE_PASSWORD"),
		QueueHost:     os.Getenv("QUEUE_HOST"),
		QueuePort:     queuePort,
	}
	log.Printf("config: %#v\n", config)
	return &config, nil
}
