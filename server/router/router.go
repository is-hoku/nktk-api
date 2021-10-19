package router

import (
	"github.com/is-hoku/nktk-api/server/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	item := new(controller.ItemController)
	e.GET("/random", item.GetRandom)

	return e
}
