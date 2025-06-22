package main

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/olahol/melody"
	"github.com/traP-jp/h25s_15/internal/cards"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/expressions"
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

	game := games.New(db, m) // handler
	card := cards.New(db, m)
	user := users.New()
	expr, err := expressions.New(db, m)
	if err != nil {
		log.Fatal("failed to create expressions handler:", err)
	}

	e.Use(user.AuthMiddleware())

	userApi := e.Group("/users")
	userApi.GET("/me", user.GetMe)

	gameApi := e.Group("/games")
	gameApi.POST("", game.CreateGame)
	gameApi.GET("/ws", game.WaitGameWS)
	gameApi.GET("/:gameID/ws", game.GameWS)
	gameApi.POST("/:gameID/submissions", expr.Post)

	e.POST("/games/:gameID/clear", card.ClearHandCards,
		card.CardsUpdatedEvent, game.ScoreUpdatedEvent)

	e.POST("/games/:gameID/picks", card.PickFieldCards,
		card.CardsUpdatedEvent)

	game.StartGameMatchLoop(context.Background())

	e.Logger.Fatal(e.Start(":8080"))
}
