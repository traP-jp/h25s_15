package games

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var GameIDSessionKey = "gameID"

func (h *Handler) GameWS(c echo.Context) error {
	gameIDStr := c.Param("gameID")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid game ID format")
	}

	h.events.HandleRequestWithKeys(c.Response().Writer, c.Request(), map[string]any{
		GameIDSessionKey: gameID,
	})
	return nil
}
