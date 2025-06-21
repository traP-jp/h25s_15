package db

type HandLimit struct {
	GameID   int64 `db:"game_id"`
	PlayerID int   `db:"player_id"`
	Limit    int   `db:"limit"`
}
