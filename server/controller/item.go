package controller

import (
	"net/http"
	"strconv"

	"github.com/is-hoku/nktk-api/server/model"
	"github.com/labstack/echo/v4"
)

func GetRandom(c echo.Context) error {
	item := model.Item{}
	class := c.QueryParams().Get("class")
	err := item.GetRandom(class)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

func GetByID(c echo.Context) error {
	item := model.Item{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = item.GetByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}
