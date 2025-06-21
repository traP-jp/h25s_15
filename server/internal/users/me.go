package users

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

type GetMeResponse struct {
	Name    string `json:"name"`
	IconURL string `json:"iconUrl"`
}

func (h *Handler) GetMe(c echo.Context) error {
	userName, err := GetUserName(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user name")
	}

	iconUrl, err := url.JoinPath("https://q.trap.jp/api/v3/public/icon", userName)
	if err != nil {
		c.Logger().Errorf("failed to create icon URL: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create icon URL")
	}

	return c.JSON(http.StatusOK, GetMeResponse{
		Name:    userName,
		IconURL: iconUrl,
	})
}
