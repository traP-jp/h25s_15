package items

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/cards"
)

func (h Handler) IncreaseFieldCards(c echo.Context, gameID uuid.UUID, numCards int) error {
	for range numCards {
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
	}
	fieldCardMaxNumber, err := h.repo.GetFieldCardsMaxNumber(c.Request().Context(), gameID)
	if err != nil {
		return fmt.Errorf("failed to get field cards max number: %w", err)
	}
	var increasedCardNumber int
	err = h.db.DB(c.Request().Context()).GetContext(c.Request().Context(), &increasedCardNumber, "SELECT count(*) FROM cards WHERE location = 'field'")
	if err != nil {
		return fmt.Errorf("failed to get field cards number: %w", err)
	}
	if increasedCardNumber != fieldCardMaxNumber {
		return fmt.Errorf("number of cards does not match: %w", err)
	}
	return nil
}
