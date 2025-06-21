package cards

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_15/internal/cards/internal/domain"
	"github.com/traP-jp/h25s_15/internal/cards/internal/events"
)

func (h Handler) CardsUpdatedEvent(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			return err
		}

		gameIDStr := c.Param("gameID")
		gameID, err := uuid.Parse(gameIDStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid game ID")
		}

		ctx := c.Request().Context()

		cards, err := h.repo.GetActiveCards(ctx, gameID)
		if err != nil {
			return fmt.Errorf("get active cards: %w", err)
		}

		handLimits, err := h.repo.GetGameHandLimit(ctx, gameID)
		if err != nil {
			return fmt.Errorf("get game hand limit: %w", err)
		}

		fieldCards := make([]events.EventCard, 0, len(cards))
		player0Cards := make([]events.EventCard, 0, len(cards))
		player1Cards := make([]events.EventCard, 0, len(cards))

		for _, card := range cards {
			eventCard := events.EventCard{
				ID:    card.ID,
				Type:  string(card.Type),
				Value: card.Value,
			}

			switch card.Location {
			case domain.CardLocationField:
				fieldCards = append(fieldCards, eventCard)
			case domain.CardLocationHand:
				if card.OwnerPlayerID == nil {
					return fmt.Errorf("card in hand without owner player ID: %v", card)
				}
				ownerPlayerID := *card.OwnerPlayerID
				switch ownerPlayerID {
				case 0:
					player0Cards = append(player0Cards, eventCard)
				case 1:
					player1Cards = append(player1Cards, eventCard)
				default:
					return fmt.Errorf("unexpected owner player ID: %d", ownerPlayerID)
				}
			default:
				return fmt.Errorf("unexpected card location: %s", card.Location)
			}
		}

		event := events.CardUpdatedEvent{
			Type:              "cardUpdated",
			FieldCards:        fieldCards,
			Player0:           player0Cards,
			Player0HandsLimit: handLimits[0],
			Player1:           player1Cards,
			Player1HandsLimit: handLimits[1],
		}

		if err := h.events.CardUpdated(ctx, gameID, event); err != nil {
			return fmt.Errorf("publish card updated event: %w", err)
		}

		return nil
	}
}
