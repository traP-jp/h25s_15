package games

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/cards"
	"github.com/traP-jp/h25s_15/internal/games/internal/events"
)

const initialFieldCardsLimit int = 4

func (h *Handler) PrepareGame(ctx context.Context, gameID uuid.UUID) error {
	err := h.db.Transaction(ctx, func(ctx context.Context) error {
		for range initialFieldCardsLimit {
			cardType, value, err := cards.DecideMakingCard(ctx)
			if err != nil {
				return fmt.Errorf("decide making card: %w", err)
			}

			cardID := uuid.New()
			err = h.repo.CreateCard(ctx, cardID, gameID, cardType, value)
			if err != nil {
				return fmt.Errorf("create card: %w", err)
			}
		}

		err := h.repo.InitializeFieldCardsLimit(ctx, gameID, initialFieldCardsLimit)
		if err != nil {
			return fmt.Errorf("initialize field cards limit: %w", err)
		}

		err = h.repo.InitializeHandLimit(ctx, gameID)
		if err != nil {
			return fmt.Errorf("initialize hand limit: %w", err)
		}

		players, err := h.repo.GetPlayers(ctx, gameID)
		if err != nil {
			return fmt.Errorf("get players: %w", err)
		}

		cards, err := h.repo.GetActiveCards(ctx, gameID)
		if err != nil {
			return fmt.Errorf("get active cards: %w", err)
		}

		handLimit, err := h.repo.GetGameHandLimit(ctx, gameID)
		if err != nil {
			return fmt.Errorf("get game hand limit: %w", err)
		}

		eventCards := make([]events.EventCard, 0, len(cards))
		for _, card := range cards {
			eventCard := events.EventCard{
				ID:    card.ID,
				Type:  string(card.Type),
				Value: card.Value,
			}
			eventCards = append(eventCards, eventCard)
		}

		event := events.GameReadyEvent{
			Type:              "gameReady",
			FieldCards:        eventCards,
			Player0:           []events.EventCard{},
			Player0HandsLimit: handLimit[0],
			Player1:           []events.EventCard{},
			Player1HandsLimit: handLimit[1],
			CurrentPlayerID:   0,
			Player0Score:      0,
			Player1Score:      0,
			StartTime:         time.Now().Add(time.Second * 5),
		}

		playerNames := [2]string{}
		for _, player := range players {
			playerNames[player.PlayerID] = player.UserName
		}

		err = h.events.GameReady(ctx, gameID, event, playerNames)
		if err != nil {
			return fmt.Errorf("send game ready event: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("transaction: %w", err)
	}

	return nil
}
