
Notifier is a backend service for sending messages to multiple Telegram users using RabbitMQ as a message broker. 

## Technologies

-   **Go (Golang)**: Main programming language
-   **RabbitMQ**: Message broker for handling message queues
-   **Telegram Bot API**: Library for interacting with the Telegram API
-   **JSON**: Data interchange format for message serialization

## Features

-   Connects to RabbitMQ and listens for incoming messages
-   Deserializes messages from JSON format into a  Message  struct
-   Sends messages to specified Telegram chat IDs using the Telegram Bot API
-   Logs messages for monitoring purposes (TODO)

## Functional

1.  Initialize configuration and connect to RabbitMQ
2.  Create a new queue service and listen for messages on the "Message" queue
3.  Deserialize received messages and process them with the Telegram service
4.  Send messages to specified chat IDs and log the results
