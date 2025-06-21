package ws

import (
	"context"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/core/corews"
)

func (e *Event) GetGameConnectedUsers(ctx context.Context, gameID uuid.UUID) ([]string, error) {
	sessions, err := e.m.Sessions()
	if err != nil {
		return nil, err
	}

	var connectedUsers []string
	for _, sess := range sessions {
		if corews.FilterGameID(gameID)(sess) {
			userNameI, ok := sess.Get(corews.SessionKeyUserName)
			if !ok {
				continue
			}
			userName, ok := userNameI.(string)
			if !ok {
				continue
			}
			connectedUsers = append(connectedUsers, userName)
		}
	}

	return connectedUsers, nil
}
