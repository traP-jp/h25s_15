package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/cards/internal/domain"
	"github.com/traP-jp/h25s_15/internal/core/coredb"
)

type Repo struct {
	db *coredb.DB
}

func New(db *coredb.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) ClearHandCards(ctx context.Context, gameID uuid.UUID, playerID int) error {
	_, err := r.db.DB(ctx).ExecContext(ctx, "UPDATE cards SET location = 'used' WHERE game_id = ? and owner_player_id = ?", gameID, playerID)
	if err != nil {
		return fmt.Errorf("update location: %w", err)
	}
	return nil
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

func (r *Repo) UpdateScore(ctx context.Context, gameID uuid.UUID, playerID int, diff int) error {
	_, err := r.db.DB(ctx).ExecContext(ctx, "UPDATE game_players SET score = score + (?) WHERE game_id = ? and player_id = ?",
		diff, gameID, playerID)
	if err != nil {
		return fmt.Errorf("update score: %w", err)
	}
	return nil
}

func (r *Repo) PickFieldCards(ctx context.Context, gameID uuid.UUID, playerID int, requestCard uuid.UUID) error {
	result, err := r.db.DB(ctx).ExecContext(ctx, "UPDATE cards SET location = 'hand', owner_player_id = ? WHERE game_id = ? and id = ?",
		playerID, gameID, requestCard)
	if err != nil {
		return fmt.Errorf("update location: %w", err)
	}
	ok, _ := result.RowsAffected()
	if ok != 1 {
		return coredb.ErrNoRecordUpdated
	}
	return nil
}

// func (r *Repo) CreateCard(ctx context.Context, gameID, cardID, cardType, value) error
func (r *Repo) CreateCard(ctx context.Context, cardID uuid.UUID, gameID uuid.UUID, cardType string, cardValue string) error {
	_, err := r.db.DB(ctx).ExecContext(ctx, "INSERT INTO cards (id, game_id, type, value, location) VALUES (?, ?, ?, ?, 'field')",
		cardID, gameID, cardType, cardValue)
	if err != nil {
		return fmt.Errorf("failed to replenish field card: %w", err)
	}
	return nil
}

func (r *Repo) GetActiveCards(ctx context.Context, gameID uuid.UUID) ([]domain.Card, error) {
	var cards []Card
	err := r.db.DB(ctx).SelectContext(ctx, &cards, "SELECT * FROM cards WHERE game_id = ? AND location IN (?, ?)",
		gameID, domain.CardLocationHand, domain.CardLocationField)
	if err != nil {
		return nil, fmt.Errorf("get active cards: %w", err)
	}

	domainCards := make([]domain.Card, 0, len(cards))
	for _, card := range cards {
		var playerID *int
		if card.OwnerPlayerID.Valid {
			id := int(card.OwnerPlayerID.Int16)
			playerID = &id
		}
		domainCards = append(domainCards, domain.Card{
			ID:            card.ID,
			Type:          domain.CardType(card.Type),
			Value:         card.Value,
			Location:      domain.CardLocation(card.Location),
			OwnerPlayerID: playerID,
		})
	}
	return domainCards, nil
}

func (r *Repo) GetGameHandLimit(ctx context.Context, gameID uuid.UUID) ([2]int, error) {
	var handLimits []HandLimit
	err := r.db.DB(ctx).SelectContext(ctx, &handLimits,
		"SELECT * FROM hand_cards_limits WHERE game_id = ? ORDER BY player_id ASC", gameID)
	if err != nil {
		return [2]int{}, fmt.Errorf("get game hand limit: %w", err)
	}
	if len(handLimits) < 2 {
		return [2]int{}, fmt.Errorf("get game hand limit: not enough data")
	}

	return [2]int{handLimits[0].HandCardsLimit, handLimits[1].HandCardsLimit}, nil
}

func (r *Repo) GetPlayerHandCards(ctx context.Context, gameID uuid.UUID, playerID int) ([]domain.Card, error) {
	var cards []Card
	err := r.db.DB(ctx).SelectContext(ctx, &cards, "SELECT * FROM cards WHERE game_id = ? AND owner_player_id = ? AND location = ?",
		gameID, playerID, domain.CardLocationHand)
	if err != nil {
		return nil, fmt.Errorf("get player hand cards: %w", err)
	}

	domainCards := make([]domain.Card, 0, len(cards))
	for _, card := range cards {
		var ownerPlayerID *int
		if card.OwnerPlayerID.Valid {
			id := int(card.OwnerPlayerID.Int16)
			ownerPlayerID = &id
		}
		domainCards = append(domainCards, domain.Card{
			ID:            card.ID,
			Type:          domain.CardType(card.Type),
			Value:         card.Value,
			Location:      domain.CardLocation(card.Location),
			OwnerPlayerID: ownerPlayerID,
		})
	}
	return domainCards, nil
}
