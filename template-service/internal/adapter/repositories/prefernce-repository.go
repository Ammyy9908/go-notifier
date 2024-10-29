package repositories

type ITemplateRepository interface {
	GetTemplate(templateId string) (interface{}, error)
}
