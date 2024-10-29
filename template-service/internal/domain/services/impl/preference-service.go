package impl

import (
	"go-notifier/commons/utils/logger"
	"go-notifier/template-service/internal/adapter/repositories"
)

type TemplateService struct {
	TemplateRepository repositories.ITemplateRepository
}

func NewTemplateService(templateRepository repositories.ITemplateRepository) *TemplateService {
	return &TemplateService{
		TemplateRepository: templateRepository,
	}
}

func (p *TemplateService) GetTemplate(templateId string) (interface{}, error) {
	log := logger.GetLogger()
	log.Info("Getting template", "templateId", templateId)
	return p.TemplateRepository.GetTemplate(templateId)
}
