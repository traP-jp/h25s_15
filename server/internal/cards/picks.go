package cards

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/users"
)

type RequestCard struct {
	CardID uuid.UUID `db:"id"`
}

func (h Handler) PickFieldCards(c echo.Context) error {
	gameIDStr := c.Param("gameID")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid uuid error")
	}

	userName, err := users.GetUserName(c)
	if err != nil {
		c.Logger().Errorf("failed to get user name: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user name")
	}

	player, err := h.repo.GetPlayer(c.Request().Context(), gameID, userName)
	if errors.Is(err, coredb.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusBadRequest, "player not found")
	}
	if err != nil {
		c.Logger().Errorf("failed to get player: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	requestCard := &RequestCard{}
	err = c.Bind(requestCard)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "no request card")
	}
	err = h.repo.PickFieldCards(c.Request().Context(), gameID, player.PlayerID, requestCard.CardID)
	if err != nil {
		c.Logger().Errorf("failed to pick a field card: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
