package models

type Message struct {
	From       string `json:"from"`
	To         string `json:"to"`
	TemplateId string `json:"templateId"`
}
