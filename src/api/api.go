package api

import (
	"log"

	"gitlab/nefco/mail-service/src/services"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type API struct {
	address  string
	services *services.Services
}

func NewApi(
	services *services.Services,
	address string,
) {
	api := &API{
		address,
		services,
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

	e.POST("/send", api.SendMail)

	err := e.Start(api.address)
	if err != nil {
		log.Fatal(
			err,
		)
	}

	log.Println("API started successfully")
}
