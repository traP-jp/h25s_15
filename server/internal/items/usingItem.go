package items

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/items/domain"
)

func (h Handler) UsingItem(c echo.Context) error {
	var item domain.Card
	err := c.Bind(&item.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "no request item")
	}
	gameIDStr := c.Param("gameID")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid uuid error")
	}
	item, err = h.repo.GetCard(c.Request().Context(), item.ID, gameID)
	if err != nil {
		c.Logger().Errorf("failed to get card infomation: %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get card infomation")
	}
	if item.Type != domain.CardTypeItem {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request card type")
	}
	switch item.Value {
	case "increaseFieldCards":
		err = h.repo.IncreaseFieldCardsMaxNumber(c.Request().Context(), gameID)
		if err != nil {
			c.Logger().Errorf("failed to increase field cards: %w", err)
		}
		err = h.IncreaseFieldCards(c, gameID, 1)
		if err != nil {
			c.Logger().Errorf("failed to increase field cards: %w", err)
		}
	case "refreshFieldCards":
		err = h.RefreshFieldCards(c, gameID)
		if err != nil {
			c.Logger().Errorf("failed to refresh field cards: %w", err)
		}
	case "clearOpponentHandCards":
		err = h.ClearOpponentHandCards(c, gameID)
		if err != nil {
			c.Logger().Errorf("failed to refresh field cards: %w", err)
		}
	case "rincreaseTurnTime":
		// err = h.RincreaseTurnTime(c, gameID)
		// if err != nil {
		// 	c.Logger().Errorf("failed to refresh field cards: %w", err)
		// }
	case "increaseHandCardsLimit":
		// err = h.repo.IncreaseHandCardsLimit(c, gameID)
		// if err != nil {
		// 	c.Logger().Errorf("failed to refresh field cards: %w", err)
		// }
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "not exist")
	}
	return nil
}
