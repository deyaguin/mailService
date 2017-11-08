package services

import (
	"net/smtp"

	"gitlab/nefco/mail-service/src/models"

	"log"

	"errors"

	"fmt"

	"crypto/tls"
	"encoding/base64"
	"net"

	"github.com/domodwyer/mailyak"
	"github.com/matcornic/hermes"
)

type Authentication struct {
	username, password string
}

func NewAuthentication(username, password string) smtp.Auth {
	return &Authentication{username, password}
}

func (a *Authentication) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *Authentication) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown fromServer")
		}
	}
	return nil, nil
}

type Mail interface {
	Send(message *models.Message) error
}

func SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	c, err := smtp.Dial(addr)
	host, _, _ := net.SplitHostPort(addr)
	if err != nil {
		fmt.Println("call dial")
		return err
	}
	defer c.Close()

	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: host, InsecureSkipVerify: true}
		if err = c.StartTLS(config); err != nil {
			fmt.Println("call start tls")
			return err
		}
	}

	if a != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(a); err != nil {
				fmt.Println("check auth with err:", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}

	header := make(map[string]string)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString(msg)
	_, err = w.Write([]byte(message))

	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

type mail struct {
	mail *mailyak.MailYak
}

func NewMail() Mail {
	m := mailyak.New(
		"mail.nefco.ru:2525",
		smtp.PlainAuth(
			"",
			"notice@nefis.local",
			"N76Qvb9t",
			"mail.nefco.ru",
		),
	)
	return &mail{m}
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
	//if err := c.mail.Send(); err != nil {
	//	log.Panic(err)
	//}
	const servername, username, password = "mail.nefco.ru:2525", "notice@nefis.local", "N76Qvb9t"
	//hostName, _, _ := net.SplitHostPort(servername)
	auth := NewAuthentication(username, password)

	if err := SendMail(servername, auth, msg.From, []string{msg.To}, []byte(emailBody)); err != nil {
		panic(err)
	}

	return nil
}
