package echo

import (
	"github/mrbryside/pocket-service/internal/domain/eventgen"
	"github/mrbryside/pocket-service/internal/entity"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h Handler) GetById(c echo.Context) error {
	id, _ := uuid.Parse("ee62a160-c0a0-495c-9358-4e494fd355ce")
	result, err := h.eventGenService.GetPocketById(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}
	return c.JSON(http.StatusOK, toGetPocketResponse(result))
}

type getPocketResponse struct {
	PocketId     uuid.UUID            `json:"pocket_id"`
	PocketName   string               `json:"pocket_name"`
	Icon         string               `json:"icon"`
	Transactions []entity.Transaction `json:"transactions"`
}

func toGetPocketResponse(eg *eventgen.EventGen) getPocketResponse {
	return getPocketResponse{
		PocketId:     eg.Pocket.Id,
		PocketName:   eg.Pocket.Name,
		Icon:         eg.Pocket.Icon,
		Transactions: eg.Transactions,
	}
}
