package services

import (
	"gitlab/nefco/mail-service/src/models"
	"log"

	"gitlab/nefco/mail-service/src/mail"

	"github.com/matcornic/hermes"
)

type MailService interface {
	Send(message *models.Message) error
}

type mailService struct {
	mail mail.Mail
}

func NewMailService(config *mail.ConnectionConfig) MailService {
	auth := mail.NewAuthentication(config.Username, config.Password)
	mail := mail.New(config.Addr, auth)
	return &mailService{mail}
}

func (mS *mailService) Send(msg *models.Message) error {
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

	mS.mail.To(msg.To)
	mS.mail.From(msg.From)
	mS.mail.Msg([]byte(emailBody))

	if err := mS.mail.Send(); err != nil {
		panic(err)
	}

	return nil
}
