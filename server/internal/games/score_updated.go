package games

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/games/internal/events"
)

func (h *Handler) ScoreUpdatedEvent(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			return err
		}

		gameIDStr := c.Param("gameID")
		gameID, err := uuid.Parse(gameIDStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid game ID")
		}

		ctx := c.Request().Context()
		if _, err := h.repo.GetGame(ctx, gameID); err != nil {
			if errors.Is(err, coredb.ErrRecordNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, "game not found")
			}
			c.Logger().Errorf("failed to get game: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		players, err := h.repo.GetPlayers(ctx, gameID)
		if err != nil {
			c.Logger().Errorf("failed to get players: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		var player0Score, player1Score int
		for _, player := range players {
			switch player.PlayerID {
			case 0:
				player0Score = player.Score
			case 1:
				player1Score = player.Score
			default:
				c.Logger().Errorf("unexpected player ID: %d", player.PlayerID)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
		}

		event := events.ScoreUpdatedEvent{
			Type:    "scoreUpdated",
			Player0: player0Score,
			Player1: player1Score,
		}
		if err := h.events.ScoreUpdated(ctx, gameID, event); err != nil {
			c.Logger().Errorf("failed to send score updated event: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return nil
	}
}
