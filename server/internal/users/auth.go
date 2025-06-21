package users

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AuthMiddleware() echo.MiddlewareFunc {
	mode := os.Getenv("MODE")
	if mode == "prod" {
		return h.prodModeAuth
	}
	log.Println("Running in development mode, using devModeAuth middleware")
	return h.devModeAuth
}

const userNameKey string = "user"

func (h *Handler) devModeAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		debugUserName := c.QueryParam("debugUserName")
		if debugUserName == "" {
			debugUserName = "traP"
		}

		c.Set(userNameKey, debugUserName)

		return next(c)
	}
}

func (h *Handler) prodModeAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Request().Header.Get("X-Forwarded-User")
		if userName == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		c.Set(userNameKey, userName)
		return next(c)
	}
}

func GetUserName(c echo.Context) (string, error) {
	userNameI := c.Get(userNameKey)
	userName, ok := userNameI.(string)
	if !ok {
		return "", errors.New("get user name from context")
	}

	return userName, nil
}
