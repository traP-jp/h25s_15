package items

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/items/domain"
	"github.com/traP-jp/h25s_15/internal/users"
)

func (h Handler) IncreaseTurnTime(c echo.Context, gameID uuid.UUID) error {
	userName, err := users.GetUserName(c)
	if err != nil {
		c.Logger().Errorf("failed to get user name: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user name")
	}
	var player domain.GamePlayer
	player, err = h.repo.GetPlayer(c.Request().Context(), gameID, userName)
	if err != nil {
		c.Logger().Errorf("failed to get player: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = h.repo.IncreaseTurnTime(c.Request().Context(), gameID, player.PlayerID)
	if err != nil {
		c.Logger().Errorf("failed to increase turn time: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return nil
}
