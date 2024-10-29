package services

import "go-notifier/notification-service/internal/domain/dtos/request"

type INotificationService interface {
	SendNotification(payload *request.NotificationRequest) (interface{}, error)
}
