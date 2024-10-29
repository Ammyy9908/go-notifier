package impl

import (
	"go-notifier/commons/utils/logger"
	"go-notifier/profile-service/internal/adapter/repositories"
)

type PreferenceService struct {
	PreferenceRepository repositories.IPreferenceRepository
}

func NewPreferenceService(preferenceRepository repositories.IPreferenceRepository) *PreferenceService {
	return &PreferenceService{
		PreferenceRepository: preferenceRepository,
	}
}

func (p *PreferenceService) GetPreferences(userID string) (interface{}, error) {
	log := logger.GetLogger()
	log.Info("Getting preferences for user", "userID", userID)
	return p.PreferenceRepository.GetPreferences(userID)
}
