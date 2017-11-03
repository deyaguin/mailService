package mail

import (
	"net/smtp"

	"gitlab/nefco/mail-service/src/models"

	"log"

	"fmt"

	"github.com/domodwyer/mailyak"
	"github.com/matcornic/hermes"
)

type Mail interface {
	Send(message *models.Message) error
}

type Client struct {
	mail   *mailyak.MailYak
	hermes hermes.Hermes
}

func NewMail() Mail {
	hermes := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Hermes",
			Link: "https://example-hermes.com/",
			// Optional product logo
			Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
	}
	mail := mailyak.New("smtp.gmail.com:465", smtp.PlainAuth("", "leagre2010@gmail.com", "GolangReact2017", "smtp.gmail.com"))
	return &Client{mail, hermes}
}

func (c *Client) Send(msg *models.Message) error {
	email := hermes.Email{
		Body: hermes.Body{
			FreeMarkdown: `
			> _Hermes_ service will shutdown the **1st August 2017** for maintenance operations.

			Services will be unavailable based on the following schedule:

			| Services | Downtime |
			| :------:| :-----------: |
			| Service A | 2AM to 3AM |
			| Service B | 4AM to 5AM |
			| Service C | 5AM to 6AM |

			---

			Feel free to contact us for any question regarding this matter at [support@hermes-example.com](mailto:support@hermes-example.com) or in our [Gitter](https://gitter.im/)

			`,
		},
	}

	emailBody, err := c.hermes.GenerateHTML(email)
	if err != nil {
		log.Println(err)
	}

	c.mail.To(msg.To)
	c.mail.From(msg.From)
	c.mail.Plain().Set(emailBody)
	c.mail.HTML().Set("Don't panic")
	fmt.Println(c.mail)

	if err := c.mail.Send(); err != nil {
		log.Println(err)
	}

	return nil
}
