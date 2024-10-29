package factory

import (
	"go-notifier/profile-service/internal/domain/factory"
	"go-notifier/profile-service/internal/port/controller"
)

type PortFactory struct {
	HealthController     *controller.HealthController
	PreferenceController *controller.PreferenceController
}

func NewPortFactory() *PortFactory {
	domains := factory.NewDomainFactory()
	return &PortFactory{
		HealthController:     controller.NewHealthController(),
		PreferenceController: controller.NewPreferenceController(domains.PreferenceService),
	}
}
