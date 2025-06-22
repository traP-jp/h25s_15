package items

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/cards"
)

func (h Handler) IncreaseFieldCards(c echo.Context, gameID uuid.UUID) error {
	cardID := uuid.New()
	cardType, cardValue, err := cards.DecideMakingCard(c.Request().Context())
	if err != nil {
		return fmt.Errorf("couldn't decide making card")
	}
	err = h.repo.CreateCard(c.Request().Context(), cardID, gameID, cardType, cardValue)
	if err != nil {
		c.Logger().Errorf("failed to create cards: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = h.repo.IncreaseFieldCardsMaxNumber(c.Request().Context(), gameID)
	if err != nil {
		c.Logger().Errorf("failed to increase field cards max number: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return nil
}
