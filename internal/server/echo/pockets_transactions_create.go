package echo

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h Handler) CreateTransaction(c echo.Context) error {
	id, _ := uuid.Parse("ee62a160-c0a0-495c-9358-4e494fd355ce")
	err := h.saverService.InsertTransaction(id, 100.0, "hey")
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}
	return c.String(http.StatusOK, "ok")
}
