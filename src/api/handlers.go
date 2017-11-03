package api

import (
	"net/http"

	"github.com/labstack/echo"
	"gitlab/nefco/mail-service/src/models"
)

func (api *API) SendMail(c echo.Context) error {
	msg := new(models.Message)

	if err := c.Bind(msg); err != nil {
		return err
	}

	api.mail.Send(msg)

	return c.JSON(http.StatusOK, "successfully")
}
