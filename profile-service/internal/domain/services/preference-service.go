package services

type IPreferenceService interface {
	GetPreferences(userID string) (interface{}, error)
}
