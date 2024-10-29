package services

type ITemplateService interface {
	GetTemplate(templateId string) (interface{}, error)
}
