package factory

import (
	"go-notifier/template-service/internal/adapter/repositories"
	"go-notifier/template-service/internal/adapter/repositories/impl"
)

type AdapterFactory struct {
	TemplateRepository repositories.ITemplateRepository
}

func NewAdapterFactory() *AdapterFactory {
	return &AdapterFactory{
		TemplateRepository: impl.NewTemplateRepository(),
	}
}
