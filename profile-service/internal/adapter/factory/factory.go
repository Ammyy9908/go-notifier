package factory

import (
	"go-notifier/profile-service/internal/adapter/repositories"
	"go-notifier/profile-service/internal/adapter/repositories/impl"
)

type AdapterFactory struct {
	PreferenceRepository repositories.IPreferenceRepository
}

func NewAdapterFactory() *AdapterFactory {
	return &AdapterFactory{
		PreferenceRepository: impl.NewPreferenceRepository(),
	}
}
