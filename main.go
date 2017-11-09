package main

import (
	"gitlab/nefco/mail-service/src/api"
	"gitlab/nefco/mail-service/src/mail"
	"gitlab/nefco/mail-service/src/services"
)

const addr, username, pwd = "mail.nefco.ru:465", "notice@nefis.local", "N76Qvb9t"

func main() {
	template := services.NewTemplate()
	mail := services.NewMailService(
		&mail.ConnectionConfig{
			addr,
			username,
			pwd,
		},
	)

	services := services.NewServices(
		template,
		mail,
	)

	api.NewApi(
		services,
		":1535",
	)
}
