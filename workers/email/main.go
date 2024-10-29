package main

import (
	"encoding/json"
	"go-notifier/commons/utils/logger"
	"go-notifier/commons/utils/queue_client"
	"go-notifier/workers/email/dto"
	"go-notifier/workers/providers/email"
	"log"
)

func EmailWorker(queueClient *queue_client.QueueClient, queueName string) {
	messages, err := queueClient.Dequeue(queueName)
	if err != nil {
		log.Fatalf("Failed to dequeue messages from %s: %v", queueName, err)
	}

	// Listen to messages in the specified queue
	for msg := range messages {
		var notification dto.NotificationDTO
		err := json.Unmarshal(msg.Body, &notification)
		if err != nil {
			log.Printf("Failed to unmarshal message from %s: %v", queueName, err)
			continue
		}

		// Process the email notification (for example, send an email)
		err = processEmailNotification(notification)
		if err != nil {
			log.Printf("Failed to process notification from %s: %v", queueName, err)
			// Implement retry logic or log the error for future handling
		} else {
			log.Printf("Successfully processed notification for user %s from %s", notification.Recipient.UserID, queueName)
		}
	}
}

func main() {
	queueClient := &queue_client.QueueClient{}
	err := queueClient.Connect("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	go EmailWorker(queueClient, "Email_high_priority")
	go EmailWorker(queueClient, "Email_standard")
	select {}
}

func processEmailNotification(notification dto.NotificationDTO) error {
	log := logger.GetLogger()
	log.Info("processEmailNotification", "notification", notification)
	emailConfig := email.EmailConfig{
		SenderEmail: "info@kodnest.com",
		Region:      "ap-south-1",
		APIKey:      "test",
		APISecret:   "test",
	}

	emailProvider, err := email.EmailProviderFactory("ses", emailConfig)
	if err != nil {
		log.Error("processEmailNotification failed to get email provider", "error", err)
		return err
	}

	log.Info("Provider", "provider", emailProvider)
	err = emailProvider.SendEmail(notification.Recipient.Email, notification.Title, notification.Body)
	if err != nil {
		log.Error("processEmailNotification failed to send email", "user_id", notification.Recipient.UserID, "error", err)
		return err
	}

	log.Info("processEmailNotification completed successfully", "user_id", notification.Recipient.UserID)
	return nil
}
