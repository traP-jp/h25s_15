package corews

import (
	"github.com/google/uuid"
	"github.com/olahol/melody"
)

var (
	SessionKeyGameID   = "gameID"
	SessionKeyUserName = "userName"
	SessionKeyWaiting  = "waiting"
	SessionKeyPlayerID = "playerID"
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

func FilterUserName(userName string) MelodyFilterFun {
	return func(s *melody.Session) bool {
		if sessUserNameI, ok := s.Get(SessionKeyUserName); ok {
			if sessUserName, ok := sessUserNameI.(string); ok {
				return sessUserName == userName
			}
		}
		return false
	}
}
