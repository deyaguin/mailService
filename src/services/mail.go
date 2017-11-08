package services

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
	mail *mailyak.MailYak
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
	const mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	var err error
	var emailBody string

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

	if msg.Type == "html" {
		emailBody, err = h.GenerateHTML(email)
	} else {
		emailBody, err = h.GeneratePlainText(email)
	}

	if err != nil {
		log.Panic(err)
	}

	c.mail.To(msg.To)
	c.mail.From(msg.From)
	c.mail.Subject("test")
	c.mail.Plain().Set(mime)
	c.mail.HTML().Set(emailBody)

	if err := c.mail.Send(); err != nil {
		log.Panic(err)
	}

	return nil
}
