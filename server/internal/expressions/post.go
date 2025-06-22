package expressions

import (
	"context"
	"errors"
	"math/big"
	"net/http"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/expressions/internal/domain"
	"github.com/traP-jp/h25s_15/internal/expressions/internal/events"
	"github.com/traP-jp/h25s_15/internal/users"
)

type submitExpressionRequest struct {
	Expression string      `json:"expression"`
	Cards      []uuid.UUID `json:"cards"`
}

type submitExpressionResponse struct {
	Success bool   `json:"success"`
	Value   string `json:"value"`
}

func (h *Handler) Post(c echo.Context) error {
	var req submitExpressionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if len(req.Cards) < 3 || len(req.Cards) > 9 {
		return echo.NewHTTPError(http.StatusBadRequest, "number of cards must be between 3 and 9")
	}

	gameIDStr := c.Param("gameID")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid game ID")
	}

	userName, err := users.GetUserName(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized user")
	}

	var val *big.Rat
	var success bool

	err = h.db.Transaction(c.Request().Context(), func(ctx context.Context) error {
		players, err := h.repo.GetPlayers(c.Request().Context(), gameID)
		if err != nil {
			c.Logger().Errorf("failed to get players for game %s: %v", gameID, err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if len(players) == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "no players found for the game")
		}

		playerIndex := slices.IndexFunc(players, func(p domain.GamePlayer) bool {
			return p.UserName == userName
		})
		if playerIndex == -1 {
			return echo.NewHTTPError(http.StatusForbidden, "user is not a player in this game")
		}
		player := players[playerIndex]

		cards, err := h.repo.GetCards(c.Request().Context(), gameID, player.PlayerID, req.Cards)
		if err != nil {
			c.Logger().Errorf("failed to get cards: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		for _, card := range cards {
			if card.Type == domain.CardTypeItem {
				return echo.NewHTTPError(http.StatusBadRequest, "cannot use item cards in expressions")
			}
		}

		cardsValuesCountMap := make(map[string]int, len(cards))
		for _, card := range cards {
			if _, ok := cardsValuesCountMap[card.Value]; !ok {
				cardsValuesCountMap[card.Value] = 0
			}
			cardsValuesCountMap[card.Value]++
		}

		expressionValuesCountMap := make(map[string]int, len(req.Expression))
		for i := range req.Expression {
			char := string([]byte{req.Expression[i]})
			if len(strings.TrimSpace(char)) == 0 {
				continue // Skip whitespace
			}
			expressionValuesCountMap[char]++
		}

		for value, count := range cardsValuesCountMap {
			expCount, ok := expressionValuesCountMap[value]
			if !ok || expCount != count {
				return echo.NewHTTPError(http.StatusBadRequest, "expression does not match the cards provided")
			}
		}

		expr, err := h.parser.ParseString("", req.Expression)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid expression syntax")
		}
		result := expr.Eval()
		val = result
		success = result.Cmp(ten) == 0
		if !success {
			return nil
		}

		score, err := calculateScore(cards)
		if err != nil {
			c.Logger().Errorf("failed to calculate score: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to calculate score")
		}

		err = h.repo.UseCards(c.Request().Context(), gameID, player.PlayerID, req.Cards)
		if err != nil {
			c.Logger().Errorf("failed to use cards: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		err = h.repo.CreateExpression(c.Request().Context(),
			uuid.New(), gameID, player.PlayerID, req.Expression, val.RatString(), score, success)
		if err != nil {
			c.Logger().Errorf("failed to create expression: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		err = h.repo.UpdatePlayerScore(c.Request().Context(), gameID, player.PlayerID, score)
		if err != nil {
			c.Logger().Errorf("failed to update player score: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		err = h.events.SubmissionSucceeded(ctx, gameID, events.SubmissionSucceededEvent{
			Type:       "submissionSucceeded",
			PlayerID:   player.PlayerID,
			Expression: req.Expression,
			Score:      score,
		})
		if err != nil {
			c.Logger().Errorf("failed to publish submission succeeded event: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, submitExpressionResponse{
		Success: success,
		Value:   val.RatString(),
	})
}

var ten = big.NewRat(10, 1)

func calculateScore(cards []domain.Card) (int, error) {
	operandCount := 0
	for _, card := range cards {
		if card.Type == domain.CardTypeOperand {
			operandCount++
		}
	}

	switch operandCount {
	case 2:
		return 1, nil
	case 3:
		return 5, nil
	case 4:
		return 10, nil
	case 5:
		return 20, nil
	default:
		return 0, errors.New("invalid number of operands")
	}
}
