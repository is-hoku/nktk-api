package controller

import (
	"net/http"

	"github.com/is-hoku/nktk-api/server/model"
	"github.com/labstack/echo/v4"
)

func GetRandom(c echo.Context) error {
	item := model.Item{}
	err := item.GetRandom()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}
