package factory

import (
	"go-notifier/profile-service/internal/adapter/factory"
	"go-notifier/profile-service/internal/domain/services"
	"go-notifier/profile-service/internal/domain/services/impl"
)

type DomainFactory struct {
	PreferenceService services.IPreferenceService
}

func NewDomainFactory() *DomainFactory {
	adapterFactory := factory.NewAdapterFactory()
	return &DomainFactory{
		PreferenceService: impl.NewPreferenceService(adapterFactory.PreferenceRepository),
	}
}
