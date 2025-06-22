package cards

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/users"
)

type RequestCard struct {
	CardID uuid.UUID `db:"id"`
}

func (h Handler) PickFieldCards(c echo.Context) error {
	gameIDStr := c.Param("gameID")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid uuid error")
	}

	userName, err := users.GetUserName(c)
	if err != nil {
		c.Logger().Errorf("failed to get user name: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user name")
	}

	player, err := h.repo.GetPlayer(c.Request().Context(), gameID, userName)
	if errors.Is(err, coredb.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusBadRequest, "player not found")
	}
	if err != nil {
		c.Logger().Errorf("failed to get player: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	limits, err := h.repo.GetGameHandLimit(c.Request().Context(), gameID)
	if err != nil {
		c.Logger().Errorf("failed to get game hand limit: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	cards, err := h.repo.GetPlayerHandCards(c.Request().Context(), gameID, player.PlayerID)
	if err != nil {
		c.Logger().Errorf("failed to get player hand cards: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if len(cards) >= limits[player.PlayerID] {
		return echo.NewHTTPError(http.StatusBadRequest, "too many cards in hand")
	}

	requestCard := &RequestCard{}
	err = c.Bind(requestCard)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "no request card")
	}
	err = h.repo.PickFieldCards(c.Request().Context(), gameID, player.PlayerID, requestCard.CardID)
	if err != nil {
		c.Logger().Errorf("failed to pick a field card: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cardID := uuid.New()
	cardType, cardValue, err := DecideMakingCard(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "couldn't decide making card")
	}
	err = h.repo.CreateCard(c.Request().Context(), cardID, gameID, cardType, cardValue)
	if err != nil {
		c.Logger().Errorf("failed to create cards: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
