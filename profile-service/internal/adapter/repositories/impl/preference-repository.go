package impl

import (
	"context"
	"go-notifier/commons/utils/logger"
	"go-notifier/profile-service/internal/adapter/constants"
	"go-notifier/profile-service/internal/adapter/db"

	"go.mongodb.org/mongo-driver/bson"
)

type PreferenceRepository struct{}

func NewPreferenceRepository() *PreferenceRepository {
	return &PreferenceRepository{}
}

func (p *PreferenceRepository) GetPreferences(userID string) (interface{}, error) {
	log := logger.GetLogger()
	log.Info("Getting preferences for user", "userID", userID)
	var result bson.M
	dbE := db.GetDatabase().Collection(constants.PREFERENCE_COLLECTION).FindOne(context.Background(), bson.M{"user_id": userID}).Decode(&result)
	if dbE != nil {
		log.Error("Error getting preferences for user", "userID", userID, "error", dbE)
		return nil, dbE
	}

	log.Info("Preferences for user", "userID", userID, "preferences", result)
	return result, nil
}
