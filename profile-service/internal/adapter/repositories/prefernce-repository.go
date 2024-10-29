package repositories

type IPreferenceRepository interface {
	GetPreferences(userID string) (interface{}, error)
}
