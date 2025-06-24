package games

import (
	"context"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type rankingItem struct {
	UserName   string `json:"name"`
	IconURL    string `json:"iconUrl"`
	Wins       int    `json:"wins"`
	Losses     int    `json:"losses"`
	Draws      int    `json:"draws"`
	TotalScore int    `json:"totalScore"`
}

type getRankingResponse struct {
	Count   int           `json:"count"`
	Ranking []rankingItem `json:"ranking"`
}

func (h *Handler) Ranking(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	if limitStr == "" {
		limitStr = "20"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		return echo.NewHTTPError(400, "invalid limit parameter")
	}

	ctx := c.Request().Context()
	var count int
	var rankingRes []rankingItem
	err = h.db.Transaction(ctx, func(ctx context.Context) error {
		ranking, err := h.repo.GetRanking(ctx, limit)
		if err != nil {
			c.Logger().Errorf("failed to get ranking: %v", err)
			return echo.NewHTTPError(500, "internal server error")
		}

		count, err = h.repo.GetUsersCount(ctx)
		if err != nil {
			c.Logger().Errorf("failed to get user count: %v", err)
			return echo.NewHTTPError(500, "internal server error")
		}

		rankingRes = make([]rankingItem, 0, len(ranking))
		for _, item := range ranking {
			rankingRes = append(rankingRes, rankingItem{
				UserName:   item.UserName,
				IconURL:    fmt.Sprintf("https://q.trap.jp/api/v3/public/icon/%s", item.UserName),
				Wins:       item.Wins,
				Losses:     item.Losses,
				Draws:      item.Draws,
				TotalScore: item.TotalScore,
			})
		}

		return nil
	})
	if err != nil {
		return err
	}

	return c.JSON(200, getRankingResponse{
		Count:   count,
		Ranking: rankingRes,
	})
}
