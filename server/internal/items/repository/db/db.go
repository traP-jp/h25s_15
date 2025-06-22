package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
	"github.com/traP-jp/h25s_15/internal/items/domain"
)

type Repo struct {
	db *coredb.DB
}

func New(db *coredb.DB) *Repo {
	return &Repo{
		db: db,
	}
}

// func (r *Repo) CreateCard(ctx context.Context, gameID, cardID, cardType, value) error

func (r *Repo) GetCard(ctx context.Context, cardID uuid.UUID, gameID uuid.UUID) (domain.Card, error) {
	var cardDB Card
	err := r.db.DB(ctx).GetContext(ctx, &cardDB, "SELECT * FROM cards WHERE id = ? and game_id = ?", cardID, gameID)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Card{}, coredb.ErrRecordNotFound
	}
	if err != nil {
		return domain.Card{}, fmt.Errorf("get card infomation: %w", err)
	}
	var value *int
	if cardDB.OwnerPlayerID.Valid {
		tmp := int(cardDB.OwnerPlayerID.Int16)
		value = &tmp
	}
	card := domain.Card{ID: cardDB.ID, Type: domain.CardType(cardDB.Type), Value: cardDB.Value, OwnerPlayerID: value, Location: domain.CardLocation(cardDB.Location)}
	return card, nil
}

func (r *Repo) CreateCard(ctx context.Context, cardID uuid.UUID, gameID uuid.UUID, cardType string, cardValue string) error {
	_, err := r.db.DB(ctx).ExecContext(ctx, "INSERT INTO cards (id, game_id, type, value, location) VALUES (?, ?, ?, ?, 'field')",
		cardID, gameID, cardType, cardValue)
	if err != nil {
		return fmt.Errorf("failed to replenish field card: %w", err)
	}
	return nil
}

func (r *Repo) GetFieldCardsMaxNumber(ctx context.Context, gameID uuid.UUID) (int, error) {
	var fieldCardMaxNumber int
	err := r.db.DB(ctx).GetContext(ctx, &fieldCardMaxNumber, "SELECT field_cards FROM field_cards_limits WHERE game_id = ?", gameID)
	if err != nil {
		return 0, fmt.Errorf("failed to get field cards max number: %w", err)
	}
	return fieldCardMaxNumber, nil
}

func (r *Repo) IncreaseFieldCardsMaxNumber(ctx context.Context, gameID uuid.UUID) error {
	_, err := r.db.DB(ctx).ExecContext(ctx, "UPDATE field_cards_limits SET field_cards = (field_cards + 1) WHERE game_id = ?", gameID)
	if err != nil {
		return fmt.Errorf("failed to get field cards max number: %w", err)
	}
	return nil
}

func (r *Repo) ClearAllCards(ctx context.Context, gameID uuid.UUID, ownerPlayerID *int, location string) (clearedNumber int, err error) {
	if ownerPlayerID == nil {
		result, err := r.db.DB(ctx).ExecContext(ctx, "UPDATE cards SET location = 'used' WHERE game_id = ? and location = ?",
			gameID, location)
		if err != nil {
			return 0, fmt.Errorf("failed to clear field cards: %w", err)
		}
		ok, _ := result.RowsAffected()
		clearedNumber = int(ok)
		return clearedNumber, nil
	} else {
		result, err := r.db.DB(ctx).ExecContext(ctx, "UPDATE cards SET location = 'used' WHERE game_id = ? and owner_player_id = ? and location = ?",
			gameID, ownerPlayerID, location)
		if err != nil {
			return 0, fmt.Errorf("failed to clear field cards: %w", err)
		}
		ok, _ := result.RowsAffected()
		clearedNumber = int(ok)
		return clearedNumber, nil
	}
}

func (r *Repo) GetPlayer(ctx context.Context, gameID uuid.UUID, userName string) (domain.GamePlayer, error) {
	var player Player
	err := r.db.DB(ctx).GetContext(ctx, &player, "SELECT * FROM game_players WHERE game_id = ? and user_name = ?", gameID, userName)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.GamePlayer{}, coredb.ErrRecordNotFound
	}
	if err != nil {
		return domain.GamePlayer{}, fmt.Errorf("get player: %w", err)
	}
	return domain.GamePlayer(player), nil
}
