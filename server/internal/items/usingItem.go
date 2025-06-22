package items

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/items/domain"
)

type useItemRequest struct {
	CardID uuid.UUID `json:"cardId"`
}

func (h Handler) UsingItem(c echo.Context) error {
	var req useItemRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "no request item")
	}
	gameIDStr := c.Param("gameID")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid uuid error")
	}
	var item domain.Card
	item, err = h.repo.GetCard(c.Request().Context(), req.CardID, gameID)
	if err != nil {
		c.Logger().Errorf("failed to get card infomation: %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get card infomation")
	}
	if item.Type != domain.CardTypeItem {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request card type")
	}

	err = h.repo.UseCard(c.Request().Context(), gameID, item.ID, *item.OwnerPlayerID)
	if err != nil {
		c.Logger().Errorf("failed to use card: %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to use card")
	}

	switch item.Value {
	case "increaseFieldCards":
		err = h.repo.IncreaseFieldCardsMaxNumber(c.Request().Context(), gameID)
		if err != nil {
			c.Logger().Errorf("failed to increase field cards: %w", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		err = h.IncreaseFieldCards(c, gameID)
		if err != nil {
			c.Logger().Errorf("failed to increase field cards: %w", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	case "refreshFieldCards":
		err = h.RefreshFieldCards(c, gameID)
		if err != nil {
			c.Logger().Errorf("failed to refresh field cards: %w", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	case "clearOpponentHandCards":
		err = h.ClearOpponentHandCards(c, gameID)
		if err != nil {
			c.Logger().Errorf("failed to refresh field cards: %w", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	case "increaseTurnTime":
	// err = h.IncreaseTurnTime(c, gameID)
	// if err != nil {
	// 	c.Logger().Errorf("failed to refresh field cards: %w", err)
	// return echo.NewHTTPError(http.StatusInternalServerError)
	// }
	case "increaseHandCardsLimit":
		err = h.IncreaseHandCardsLimit(c, gameID)
		if err != nil {
			c.Logger().Errorf("failed to refresh field cards: %w", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "not exist")
	}
	return nil
}
