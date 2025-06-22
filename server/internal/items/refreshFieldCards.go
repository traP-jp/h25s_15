package items

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/cards"
)

func (h Handler) RefreshFieldCards(c echo.Context, gameID uuid.UUID) error {
	fieldCardMaxNumber, err := h.repo.GetFieldCardsMaxNumber(c.Request().Context(), gameID)
	if err != nil {
		return fmt.Errorf("failed to get field cards max number: %w", err)
	}

	//　フィールドカードを全て破棄
	clearedNumber, err := h.repo.ClearAllCards(c.Request().Context(), gameID, nil, "field")
	if err != nil {
		return fmt.Errorf("failed to clear cards: %w", err)
	}
	if clearedNumber != fieldCardMaxNumber {
		return fmt.Errorf("not match cleared cards number: %w", err)
	}

	// フィールドカードを補充する
	for range fieldCardMaxNumber {
		cardType, value, err := cards.DecideMakingCard(c.Request().Context())
		if err != nil {
			return fmt.Errorf("failed to decide making card: %w", err)
		}
		cardID := uuid.New()
		err = h.repo.CreateCard(c.Request().Context(), cardID, gameID, cardType, value)
		if err != nil {
			return fmt.Errorf("failed to create card: %w", err)
		}
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
