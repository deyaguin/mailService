package api

import (
	"net/http"

	"gitlab/nefco/mail-service/src/models"

	"github.com/labstack/echo"
)

func (api *API) SendMail(c echo.Context) error {
	msg := new(models.Message)

	if err := c.Bind(msg); err != nil {
		return err
	}

	if err := api.services.MailService.Send(msg); err != nil {
		return c.JSON(http.StatusNoContent, "failed")
	}

	return c.JSON(http.StatusOK, "successfully")
}
