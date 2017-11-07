package main

import (
	"gitlab/nefco/mail-service/src/api"
	"gitlab/nefco/mail-service/src/services"
)

func main() {
	template := services.NewTemplate()
	mail := services.NewMail()

	services := services.NewServices(
		template,
		mail,
	)

	api.NewApi(
		services,
		":1536",
	)
}
