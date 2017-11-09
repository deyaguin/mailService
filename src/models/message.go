package models

type Message struct {
	From       string   `json:"from"`
	To         []string `json:"to"`
	TemplateId int      `json:"templateId"`
	Template   Template `json:"template"`
	Type       string   `json:"type"`
}
