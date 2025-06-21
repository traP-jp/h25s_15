package corews

import (
	"github.com/google/uuid"
	"github.com/olahol/melody"
)

var (
	SessionKeyGameID   = "gameID"
	SessionKeyUserName = "userName"
)

type MelodyFilterFun func(*melody.Session) bool

func FilterGameID(gameID uuid.UUID) MelodyFilterFun {
	return func(s *melody.Session) bool {
		if sessGameIDI, ok := s.Get(SessionKeyGameID); ok {
			if sessGameID, ok := sessGameIDI.(uuid.UUID); ok {
				return sessGameID == gameID
			}
		}
		return false
	}
}
