package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello :
func (hdl Handler) Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
