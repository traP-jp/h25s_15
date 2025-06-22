package games

import (
	"slices"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/games/internal/domain"
	"github.com/traP-jp/h25s_15/internal/users"
)

func (h *Handler) GameTurnAuth(next echo.HandlerFunc) echo.HandlerFunc {
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
			c.Logger().Errorf("Failed to get player for game %s: %v", gameID, err)
			return echo.NewHTTPError(500)
		}

		var player domain.Player
		idx := slices.IndexFunc(players, func(p domain.Player) bool {
			if p.UserName == userName {
				player = p
				return true
			}
			return false
		})
		if idx == -1 {
			return echo.NewHTTPError(403, "Forbidden: User is not a player in this game")
		}
		player = players[idx]

		turn, err := h.repo.GetTurn(c.Request().Context(), gameID)
		if err != nil {
			c.Logger().Errorf("Failed to get turn for game %s: %v", gameID, err)
			return echo.NewHTTPError(500)
		}
		if turn.PlayerID != player.PlayerID {
			return echo.NewHTTPError(403, "Forbidden: It's not your turn")
		}

		return next(c)
	}
}
