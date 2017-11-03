package api

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gitlab/nefco/mail-service/src/mail"
)

type API struct {
	address string
	mail mail.Mail
}

func NewApi(address string) {
	api := &API{
		address,
		mail.NewMail(),
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

	e.POST("/send", api.SendMail)

	err := e.Start(api.address)
	if err != nil {
		log.Fatal(
			"API start failed",
		)
	}

	log.Println("API started successfully")
}
