package games

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/games/internal/domain"
)

type getGameResultResponse struct {
	GameID                    uuid.UUID `json:"gameId"`
	Player0Name               string    `json:"player0Name"`
	Player1Name               string    `json:"player1Name"`
	Player0Score              int       `json:"player0Score"`
	Player1Score              int       `json:"player1Score"`
	Player0SuccessExpressions []string  `json:"player0SuccessExpressions"`
	Player1SuccessExpressions []string  `json:"player1SuccessExpressions"`
}

func (h *Handler) GetResult(c echo.Context) error {
	gameIDStr := c.Param("gameID")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid game ID")
	}

	players, err := h.repo.GetPlayers(c.Request().Context(), gameID)
	if err != nil {
		c.Logger().Errorf("failed to get players for game %s: %v",
			gameID, err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	expressions, err := h.repo.GetSuccessExpressions(c.Request().Context(), gameID)
	if err != nil {
		c.Logger().Errorf("failed to get success expressions for game %s: %v",
			gameID, err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	var player0, player1 domain.Player
	for _, p := range players {
		switch p.PlayerID {
		case 0:
			player0 = p
		case 1:
			player1 = p
		default:
			c.Logger().Errorf("unexpected player ID %d in game %s", p.PlayerID, gameID)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	player0Expressions := make([]string, 0, len(expressions))
	player1Expressions := make([]string, 0, len(expressions))
	for _, expr := range expressions {
		switch expr.PlayerID {
		case player0.PlayerID:
			player0Expressions = append(player0Expressions, expr.Expression)
		case player1.PlayerID:
			player1Expressions = append(player1Expressions, expr.Expression)
		default:
			c.Logger().Errorf("unexpected player ID %d in success expressions for game %s", expr.PlayerID, gameID)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	return c.JSON(http.StatusOK, getGameResultResponse{
		GameID:                    gameID,
		Player0Name:               player0.UserName,
		Player1Name:               player1.UserName,
		Player0Score:              player0.Score,
		Player1Score:              player1.Score,
		Player0SuccessExpressions: player0Expressions,
		Player1SuccessExpressions: player1Expressions,
	})
}
