package queue_client

import (
	"fmt"
	"go-notifier/commons/utils/logger"
	"log"

	"github.com/streadway/amqp"
)

type IQueueInterface interface {
	Connect(url string) error
	Enqueue(queueName string, message []byte) error
	Dequeue(queueName string) (<-chan amqp.Delivery, error)
	Close() error
}

type QueueClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewQueueClient() *QueueClient {
	return &QueueClient{}
}

func (qc *QueueClient) Connect(url string) error {
	var err error
	qc.conn, err = amqp.Dial(url)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	qc.channel, err = qc.conn.Channel()
	if err != nil {
		qc.conn.Close() // Ensure connection is closed if channel creation fails
		return fmt.Errorf("failed to open a channel: %w", err)
	}

	log.Printf("Connected to RabbitMQ at %s", url)
	return nil
}

func (qc *QueueClient) Enqueue(queueName string, message []byte) error {

	log := logger.GetLogger()
	log.Info("Enqueuing message to queue", "queueName", queueName)
	_, err := qc.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	err = qc.channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Info("Message enqueued to %s: %s", queueName, string(message))
	return nil
}

func (qc *QueueClient) Dequeue(queueName string) (<-chan amqp.Delivery, error) {
	_, err := qc.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	messages, err := qc.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume messages: %w", err)
	}

	return messages, nil
}

func (qc *QueueClient) Close() error {
	if qc.channel != nil {
		if err := qc.channel.Close(); err != nil {
			return fmt.Errorf("failed to close channel: %w", err)
		}
	}
	if qc.conn != nil {
		if err := qc.conn.Close(); err != nil {
			return fmt.Errorf("failed to close connection: %w", err)
		}
	}
	return nil
}
