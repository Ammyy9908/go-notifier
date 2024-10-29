package factory

import (
	"go-notifier/template-service/internal/domain/factory"
	"go-notifier/template-service/internal/port/controller"
)

type PortFactory struct {
	HealthController   *controller.HealthController
	TemplateController *controller.TemplateController
}

func NewPortFactory() *PortFactory {
	domains := factory.NewDomainFactory()
	return &PortFactory{
		HealthController:   controller.NewHealthController(),
		TemplateController: controller.NewTemplateController(domains.TemplateService),
	}
}
