package services

type Template interface {
	Create(int, []interface{}) (string, error)
}

type template struct {
	storage interface{}
}

func NewTemplate() Template { return nil }

func (service *template) Create(templateId int, params []interface{}) (string, error) {
	return "", nil
}
