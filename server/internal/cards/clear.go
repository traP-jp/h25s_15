package cards

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
)

func (h Handler) ClearHandCards(c echo.Context) error {
	gameIDStr := c.Param("gameID")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid uuid error")
	}
	userName := "mizu"
	player, err := h.repo.GetPlayer(c.Request().Context(), gameID, userName)
	if errors.Is(err, coredb.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusBadRequest, "player not found")
	}
	if err != nil {
		log.Printf("failed to get player: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = h.repo.ClearHandCards(c.Request().Context(), gameID, player.PlayerID)
	if err != nil {
		log.Printf("failed to clear hand cards: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = h.repo.UpdateScore(c.Request().Context(), gameID, player.PlayerID, player.Score)
	if err != nil {
		log.Printf("failed to clear hand cards: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
