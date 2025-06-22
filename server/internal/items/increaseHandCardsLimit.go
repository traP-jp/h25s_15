package items

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/items/domain"
	"github.com/traP-jp/h25s_15/internal/users"
)

func (h Handler) IncreaseHandCardsLimit(c echo.Context, gameID uuid.UUID) error {
	userName, err := users.GetUserName(c)
	if err != nil {
		c.Logger().Errorf("failed to get user name: %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user name")
	}
	var player domain.GamePlayer
	player, err = h.repo.GetPlayer(c.Request().Context(), gameID, userName)
	if err != nil {
		c.Logger().Errorf("failed to get player: %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = h.repo.IncreaseHandCardsLimit(c.Request().Context(), gameID, player.PlayerID)
	if err != nil {
		c.Logger().Errorf("failed to increase hand cards limit: %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return nil
}
