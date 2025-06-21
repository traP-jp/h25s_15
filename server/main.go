package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/olahol/melody"
	"github.com/traP-jp/h25s_15/internal/cards"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/games"
	"github.com/traP-jp/h25s_15/internal/users"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger(), middleware.Recover())

	db, err := coredb.New()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	m := melody.New()

	_ = games.New(db, m) // handler

	card := cards.New(db, m)
	user := users.New()
	e.Use(user.AuthMiddleware())

	userApi := e.Group("/users")
	userApi.GET("/me", user.GetMe)

	e.POST("/games/:gameID/clear", card.ClearHandCards)

	e.Logger.Fatal(e.Start(":8080"))
}
