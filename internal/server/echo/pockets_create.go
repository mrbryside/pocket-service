package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) CreatePockets(c echo.Context) error {
	err := h.saverService.InsertPocket("name", "test")
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}
	return c.String(http.StatusOK, "ok")
}
