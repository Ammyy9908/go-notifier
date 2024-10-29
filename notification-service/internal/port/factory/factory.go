package factory

import (
	"go-notifier/commons/utils/validator"
	"go-notifier/notification-service/internal/domain/factory"
	"go-notifier/notification-service/internal/port/controller"
)

type PortFactory struct {
	HealthController       *controller.HealthController
	NotificationController *controller.NotificationController
}

func NewPortFactory() *PortFactory {
	domains := factory.NewDomainFactory()
	rv := validator.NewValidator()
	return &PortFactory{
		HealthController:       controller.NewHealthController(),
		NotificationController: controller.NewNotificationController(rv, domains.NotificationService),
	}
}
