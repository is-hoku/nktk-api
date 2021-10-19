package router

import (
	"github.com/is-hoku/nktk-api/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	nktk := new(controller.NktkController)
	e.GET("/", nktk.RandomOne)

	return e
}