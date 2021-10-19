package controller

import (
	"net/http"

	"github.com/is-hoku/nktk-api/server/model"
	"github.com/labstack/echo/v4"
)

type ItemController struct{}

var itemModel = new(model.ItemModel)

func (i ItemController) GetRandom(c echo.Context) error {
	item, err := itemModel.GetRandom()
	if err != nil {
		c.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, item)
}
