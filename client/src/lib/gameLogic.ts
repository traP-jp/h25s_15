import type { Card, GameEvent, PlayerInfo } from './type'

/* eslint-disable */
class GameInfo {
  gameId: string
  playerId: number
  fieldCards: Card[]

  players: PlayerInfo[]

  currentPlayerId: number
  turn: number
  turnTimeRemaining: number

  constructor(gameId: string, playerId: number) {
    this.gameId = gameId
    this.playerId = playerId
    this.fieldCards = []
    this.players = [
      { id: 0, score: 0, handsLimit: 0, cards: [], expression: '' },
      { id: 1, score: 0, handsLimit: 0, cards: [], expression: '' },
    ]
    this.currentPlayerId = -1
    this.turn = 0
    this.turnTimeRemaining = 0
  }

  useCard(cardId: string): void {}
  addOperator(operator: '(' | ')'): void {}
  deleteExpr(): void {}

  onEvent(event: GameEvent): void {
    switch (event.type) {
      case 'gameReady':
        this.handleGameReady(event)
        break
      case 'gameStarted':
        this.handleGameStarted(event)
        break
      case 'turnStarted':
        this.handleTurnStarted(event)
        break
      case 'cardsUpdated':
        this.handleCardsUpdated(event)
        break
      case 'turnTimeRemainingChanged':
        this.handleTurnTimeRemainingChanged(event)
        break
      case 'submissionSucceeded':
        this.handleSubmissionSucceeded(event)
        break
      case 'scoreUpdated':
        this.handleScoreUpdated(event)
        break
      case 'turnEnded':
        this.handleTurnEnded(event)
        break
      case 'gameEnded':
        this.handleGameEnded(event)
        break
      default:
        console.warn('Unhandled game event:', event)
    }
  }

  private handleGameReady(event: GameEvent): void {}
  private handleGameStarted(event: GameEvent): void {}
  private handleTurnStarted(event: GameEvent): void {}
  private handleCardsUpdated(event: GameEvent): void {}
  private handleTurnTimeRemainingChanged(event: GameEvent): void {}
  private handleSubmissionSucceeded(event: GameEvent): void {}
  private handleScoreUpdated(event: GameEvent): void {}
  private handleTurnEnded(event: GameEvent): void {}
  private handleGameEnded(event: GameEvent): void {}
}
