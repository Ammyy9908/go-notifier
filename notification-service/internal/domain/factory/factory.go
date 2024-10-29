package factory

import (
	"go-notifier/commons/utils/api_client"
	"go-notifier/commons/utils/queue_client"
	"go-notifier/notification-service/internal/domain/services"
	"go-notifier/notification-service/internal/domain/services/impl"
	"log"
	"time"
)

type DomainFactory struct {
	NotificationService services.INotificationService
}

func NewDomainFactory() *DomainFactory {
	api_client := api_client.NewClient("http://localhost:8081", api_client.WithTimeout(10*time.Second))
	queueClient := &queue_client.QueueClient{}
	err := queueClient.Connect("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	return &DomainFactory{
		NotificationService: impl.NewNotificationService(api_client, queueClient),
	}
}
