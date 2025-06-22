package ws

import (
	"context"
	"errors"
	"fmt"

	"github.com/traP-jp/h25s_15/internal/core/corews"
)

func (e *Event) GetConnectedWaitingUsers(ctx context.Context) ([]string, error) {
	sessions, err := e.m.Sessions()
	if err != nil {
		return nil, fmt.Errorf("get sessions: %w", err)
	}

	userNames := make([]string, 0, len(sessions))
	for _, session := range sessions {
		if _, ok := session.Get(corews.SessionKeyWaiting); ok {
			userNameI, ok := session.Get(corews.SessionKeyUserName)
			if !ok {
				return nil, errors.New("session missing user name")
			}
			userName, ok := userNameI.(string)
			if !ok {
				return nil, fmt.Errorf("session user name is not a string: %T", userNameI)
			}

			userNames = append(userNames, userName)
		}
	}

	return userNames, nil
}
