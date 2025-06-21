package events

import (
	"context"

	"github.com/google/uuid"
)

type Event interface {
	CardUpdated(ctx context.Context, gameID uuid.UUID) error
}
