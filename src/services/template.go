package services

import "gitlab/nefco/mail-service/src/models"

type Template interface {
	Create(int, []interface{}) (string, error)
}

type template struct {
	storage interface{}
}

func NewTemplate() Template { return nil }

func (service *template) Create(template *models.Template) error {
	return nil
}

func (service *template) GetTemplateById(templateId int) (*models.Template, error) {
	return nil, nil
}

func (service *template) GetTemplates() (*[]models.Template, error) {
	return nil, nil
}
