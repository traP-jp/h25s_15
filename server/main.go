package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/olahol/melody"
	"github.com/traP-jp/h25s_15/internal/cards"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/games"
	"github.com/traP-jp/h25s_15/internal/items"
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

	game := games.New(db, m) // handler
	card := cards.New(db, m)
	user := users.New()
	item := items.New(db, m)

	e.Use(user.AuthMiddleware())

	userApi := e.Group("/users")
	userApi.GET("/me", user.GetMe)

	gameApi := e.Group("/games")
	gameApi.GET("/:gameID/ws", game.GameWS)

	e.POST("/games/:gameID/clear", card.ClearHandCards, card.CardsUpdatedEvent)

	e.POST("/games/:gameID/picks", card.PickFieldCards)

	e.POST("/games/:gameID/items", item.UsingItem)

	e.Logger.Fatal(e.Start(":8080"))
}
