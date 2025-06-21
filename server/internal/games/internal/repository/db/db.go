package db

import "github.com/traP-jp/h25s_15/internal/core/coredb"

type Repo struct {
	db *coredb.DB
}

func New(db *coredb.DB) *Repo {
	return &Repo{
		db: db,
	}
}
