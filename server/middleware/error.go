package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIError struct {
	StatusCode int
	Message    string
}

// TODO: FIX CUSTOM ERROR HANDLING
func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)

	if he, ok := err.(*APIError); ok {
		code = he.StatusCode
		msg = he.Message
	}
	if !c.Request().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			err := c.JSON(code, echo.NewHTTPError(code))
		}
	}
}
