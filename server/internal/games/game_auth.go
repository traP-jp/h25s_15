package games

import (
	"slices"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/games/internal/domain"
	"github.com/traP-jp/h25s_15/internal/users"
)

func (h *Handler) GamePlayerAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		gameIDStr := c.Param("gameID")
		gameID, err := uuid.Parse(gameIDStr)
		if err != nil {
			return echo.NewHTTPError(400, "Invalid game ID")
		}

		userName, err := users.GetUserName(c)
		if err != nil {
			return echo.NewHTTPError(401, "Unauthorized")
		}

		players, err := h.repo.GetPlayers(c.Request().Context(), gameID)
		if err != nil {
			c.Logger().Errorf("Failed to get players for game %s: %v", gameID, err)
			return echo.NewHTTPError(500)
		}

		if !slices.ContainsFunc(players, func(p domain.Player) bool {
			return p.UserName == userName
		}) {
			return echo.NewHTTPError(403, "Forbidden: User is not a player in this game")
		}

		return next(c)
	}
}
