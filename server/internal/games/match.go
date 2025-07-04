package games

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_15/internal/games/internal/domain"
	"github.com/traP-jp/h25s_15/internal/games/internal/repository"
)

func (h *Handler) StartGameMatchLoop(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("game match loop stopped")
				return
			case <-time.After(2 * time.Second):
				h.gameMatch(ctx)
			}
		}
	}()
	log.Println("game match loop started")
}

func (h *Handler) gameMatch(ctx context.Context) {
	defer func() {
		if cause := recover(); cause != nil {
			log.Printf("game match loop panicked: %v\n", cause)
			return
		}
	}()

	err := h.db.Transaction(ctx, func(ctx context.Context) error {
		waitingPlayers, err := h.repo.GetWaitingPlayers(ctx)
		if err != nil {
			return fmt.Errorf("get waiting players: %w", err)
		}

		connectedUsers, err := h.events.GetConnectedWaitingUsers(ctx)
		if err != nil {
			return fmt.Errorf("get connected users: %w", err)
		}

		connectedMap := make(map[string]struct{}, len(connectedUsers))
		for _, user := range connectedUsers {
			connectedMap[user] = struct{}{}
		}

		connectedWaitingPlayers := make([]domain.WaitingPlayer, 0, len(waitingPlayers))
		for _, player := range waitingPlayers {
			if _, ok := connectedMap[player.UserName]; ok {
				connectedWaitingPlayers = append(connectedWaitingPlayers, player)
			}
		}

		matchCount := len(connectedWaitingPlayers) / 2
		gameIDs := make([]uuid.UUID, 0, matchCount)
		matches := make([]repository.CreatePlayersArg, 0, matchCount)
		matchedUserNames := make([]string, 0, len(connectedWaitingPlayers))

		if matchCount == 0 {
			return nil
		}

		for i := range matchCount {
			player0 := connectedWaitingPlayers[i*2]
			player1 := connectedWaitingPlayers[i*2+1]

			gameID := uuid.New()
			gameIDs = append(gameIDs, gameID)

			matches = append(matches, repository.CreatePlayersArg{
				GameID:    gameID,
				UserName0: player0.UserName,
				UserName1: player1.UserName,
			})
			matchedUserNames = append(matchedUserNames, player0.UserName, player1.UserName)
		}

		if err := h.repo.CreateGames(ctx, gameIDs); err != nil {
			return fmt.Errorf("create games: %w", err)
		}
		if err := h.repo.CreatePlayers(ctx, matches); err != nil {
			return fmt.Errorf("create players: %w", err)
		}
		if err := h.repo.DeleteWaitingPlayers(ctx, matchedUserNames); err != nil {
			return fmt.Errorf("delete waiting players: %w", err)
		}
		log.Printf("matched %d games with %d players", matchCount, len(matchedUserNames))

		errorList := []error{}
		for _, match := range matches {
			playerName0 := match.UserName0
			playerName1 := match.UserName1
			if err := h.events.GameMatched(ctx, [2]string{playerName0, playerName1}, match.GameID); err != nil {
				errorList = append(errorList, fmt.Errorf("send game matched event for game %v: %w", match.GameID, err))
				continue // 他のゲームのマッチングの通知を続ける
			}
		}
		if len(errorList) > 0 {
			return fmt.Errorf("failed to send game matched events: %v", errorList)
		}

		return nil
	})
	if err != nil {
		log.Printf("game match failed: %v\n", err)
		return
	}
}
