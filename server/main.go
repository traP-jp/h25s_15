package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/olahol/melody"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/games"
)

func main() {
	e := echo.New()

	db, err := coredb.New()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	m := melody.New()

	_ = games.New(db, m) // handler

	e.Logger.Fatal(e.Start(":8080"))
}
