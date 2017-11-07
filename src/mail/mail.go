package mail

import (
	"net/smtp"

	"gitlab/nefco/mail-service/src/models"

	"log"

	"github.com/domodwyer/mailyak"
	"github.com/matcornic/hermes"
)

type Mail interface {
	Send(message *models.Message) error
}

type mail struct {
	client *mailyak.MailYak
}

func NewMail() Mail {
	mail := mailyak.New(
		"smtp.gmail.com:587",
		smtp.PlainAuth(
			"",
			"leagre2010@gmail.com",
			"GolangReact2017",
			"smtp.gmail.com",
		),
	)
	return &mail{mail}
}

func (c *mail) Send(msg *models.Message) error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	h := hermes.Hermes{
		Product: hermes.Product{
			Name:      "Nefco",
			Link:      "nefco.ru",
			Copyright: "Nefco",
		},
	}
	email := hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Dictionary: []hermes.Entry{
				{Key: "Firstname", Value: "Jon"},
				{Key: "Lastname", Value: "Snow"},
				{Key: "Birthday", Value: "01/01/283"},
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Text: "Confirm your account",
						Link: "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		log.Panic(err)
	}

	c.client.To(msg.To)
	c.client.From(msg.From)
	c.client.Subject("test")
	c.client.Plain().Set(mime)
	c.client.HTML().Set(emailBody)

	if err := c.client.Send(); err != nil {
		log.Panic(err)
	}

	return nil
}
