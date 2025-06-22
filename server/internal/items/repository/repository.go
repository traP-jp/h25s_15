package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/items/domain"
)

type Repo interface {
	GetCard(c context.Context, cardID uuid.UUID, gameID uuid.UUID) (domain.Card, error)
	CreateCard(c context.Context, cardID uuid.UUID, gameID uuid.UUID, cardType string, cardValue string) error
	GetFieldCardsMaxNumber(c context.Context, gameID uuid.UUID) (fieldCardMaxNumber int, err error)
	IncreaseFieldCardsMaxNumber(c context.Context, gameID uuid.UUID) error
	ClearAllCards(c context.Context, gameID uuid.UUID, ownerPlayerID *int, location string) (clearedNumber int, err error)
	GetPlayer(c context.Context, gameID uuid.UUID, userName string) (domain.GamePlayer, error)
	IncreaseHandCardsLimit(c context.Context, gameID uuid.UUID, playerID int) error
	UseCard(ctx context.Context, gameID uuid.UUID, cardID uuid.UUID, playerID int) error
}
