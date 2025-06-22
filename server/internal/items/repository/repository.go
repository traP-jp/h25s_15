package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/items/domain"
)

type Repo interface {
	UsingItem(c context.Context) error
	GetCard(c context.Context, cardID uuid.UUID, gameID uuid.UUID) (domain.Card, error)

	IncreaseFieldCards(c context.Context, gameID uuid.UUID, numCards int) error
	CreateCard(c context.Context, cardID uuid.UUID, gameID uuid.UUID, cardType string, cardValue string) error
	GetFieldCardsMaxNumber(c context.Context, gameID uuid.UUID) (fieldCardMaxNumber int, err error)
	IncreaseFieldCardsMaxNumber(c context.Context, gameID uuid.UUID) error

	RefreshFieldCards(c context.Context, gameID uuid.UUID) error
	ClearAllCards(c context.Context, gameID uuid.UUID, ownerPlayerID *int, location string) (clearedNumber int, err error)

	ClearOpponentHandCards(c context.Context, gameID uuid.UUID) error
	GetPlayer(c context.Context, gameID uuid.UUID, userName string) (domain.GamePlayer, error)

	RincreaseTurnTime(c context.Context, gameID uuid.UUID) error
	IncreaseHandCardsLimit(c context.Context, gameID uuid.UUID) error
}
