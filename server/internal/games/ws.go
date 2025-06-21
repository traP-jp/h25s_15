package games

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/core/corews"
	"github.com/traP-jp/h25s_15/internal/users"
)

func (h *Handler) GameWS(c echo.Context) error {
	gameIDStr := c.Param("gameID")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid game ID format")
	}

	userName, err := users.GetUserName(c)
	if err != nil {
		c.Logger().Errorf("failed to get user name: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	h.events.HandleRequestWithKeys(c.Response().Writer, c.Request(), map[string]any{
		corews.SessionKeyGameID:   gameID,
		corews.SessionKeyUserName: userName,
	})
	return nil
}

func (h *Handler) WaitGameWS(c echo.Context) error {
	userName, err := users.GetUserName(c)
	if err != nil {
		c.Logger().Errorf("failed to get user name: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	h.events.HandleRequestWithKeys(c.Response().Writer, c.Request(), map[string]any{
		corews.SessionKeyUserName: userName,
		corews.SessionKeyWaiting:  struct{}{},
	})
	return nil
}
