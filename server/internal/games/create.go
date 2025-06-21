package games

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/users"
)

func (h *Handler) CreateGame(c echo.Context) error {
	userName, err := users.GetUserName(c)
	if err != nil {
		c.Logger().Errorf("failed to get user name: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	err = h.repo.CreateWaitingPlayer(c.Request().Context(), userName)
	if errors.Is(err, coredb.ErrDuplicateKey) {
		return echo.NewHTTPError(http.StatusBadRequest, "player already waiting")
	}
	if err != nil {
		c.Logger().Errorf("failed to create waiting player: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
