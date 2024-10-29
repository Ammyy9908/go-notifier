package factory

import (
	"go-notifier/template-service/internal/adapter/factory"
	"go-notifier/template-service/internal/domain/services"
	"go-notifier/template-service/internal/domain/services/impl"
)

type DomainFactory struct {
	TemplateService services.ITemplateService
}

func NewDomainFactory() *DomainFactory {
	adapterFactory := factory.NewAdapterFactory()
	return &DomainFactory{
		TemplateService: impl.NewTemplateService(adapterFactory.TemplateRepository),
	}
}
