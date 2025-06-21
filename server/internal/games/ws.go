package games

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/core/corews"
	"github.com/traP-jp/h25s_15/internal/games/internal/domain"
	"github.com/traP-jp/h25s_15/internal/users"
	"golang.org/x/sync/errgroup"
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

	players, err := h.repo.GetPlayers(c.Request().Context(), gameID)
	if err != nil {
		c.Logger().Errorf("failed to get players for game %s: %v", gameID, err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	playerID := slices.IndexFunc(players, func(p domain.Player) bool {
		return p.UserName == userName
	})
	if playerID == -1 {
		c.Logger().Errorf("user %s is not a player in game %s", userName, gameID)
		return echo.NewHTTPError(http.StatusForbidden, "You are not a player in this game")
	}

	connectedUsers, err := h.events.GetGameConnectedUsers(c.Request().Context(), gameID)
	if err != nil {
		c.Logger().Errorf("failed to get connected users for game %s: %v", gameID, err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	eg, ctx := errgroup.WithContext(c.Request().Context())
	eg.Go(func() error {
		func(context.Context) {
			h.events.HandleRequestWithKeys(c.Response().Writer, c.Request(), map[string]any{
				corews.SessionKeyGameID:   gameID,
				corews.SessionKeyUserName: userName,
				corews.SessionKeyPlayerID: playerID,
			})
		}(ctx)
		return nil
	})

	eg.Go(func() error {
		if len(slices.DeleteFunc(connectedUsers, func(u string) bool {
			return u == userName
		})) == 0 {
			return nil
		}

		err := h.PrepareGame(ctx, gameID)
		if err != nil {
			c.Logger().Errorf("failed to start game %s: %v", gameID, err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		err = h.StartGame(ctx, gameID, time.Now().Add(time.Second*5))
		if err != nil {
			c.Logger().Errorf("failed to start game %s: %v", gameID, err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		c.Logger().Errorf("error handling game websocket request: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

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
