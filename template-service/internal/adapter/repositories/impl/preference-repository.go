package impl

import (
	"context"
	"go-notifier/commons/utils/logger"
	"go-notifier/template-service/internal/adapter/constants"
	"go-notifier/template-service/internal/adapter/db"

	"go.mongodb.org/mongo-driver/bson"
)

type TemplateRepository struct{}

func NewTemplateRepository() *TemplateRepository {
	return &TemplateRepository{}
}

func (p *TemplateRepository) GetTemplate(templateId string) (interface{}, error) {
	log := logger.GetLogger()
	log.Info("Getting template", "templateId", templateId)
	var result bson.M
	dbE := db.GetDatabase().Collection(constants.TEMPLATE_COLLECTION).FindOne(context.Background(), bson.M{"template_id": templateId}).Decode(&result)
	if dbE != nil {
		log.Error("Error getting template", "templateId", templateId, "error", dbE)
		return nil, dbE
	}

	log.Info("Template", "templateId", templateId, "template", result)
	return result, nil
}
