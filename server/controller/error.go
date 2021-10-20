package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (err *APIError) Error() string {
	return fmt.Sprintf("%s", err.Message)
}

func JSONErrorHandler(err error, c echo.Context) {
	errmsg := err.Error()
	fmt.Println(errmsg)
	if he, ok := err.(*echo.HTTPError); ok {
		if he.Code == 404 {
			c.JSON(he.Code, APIError{
				Status:  he.Code,
				Message: "Not Found",
			})
		} else if he.Code == 500 {
			c.JSON(he.Code, APIError{
				Status:  he.Code,
				Message: "Internal Server Error",
			})
		}
	} else { // 内部での処理のエラーは隠す
		c.JSON(500, APIError{
			Status:  500,
			Message: "Internal Server Error",
		})
	}
}
