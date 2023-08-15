package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func JanuaryTwentyFive(c echo.Context, dateMethod func() string) error {
	return c.String(http.StatusOK, dateMethod())
}
