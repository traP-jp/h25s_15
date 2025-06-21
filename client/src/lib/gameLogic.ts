import type {
  Card,
  CardsUpdated,
  GameEnded,
  GameEvent,
  GameReady,
  GameStarted,
  PlayerInfo,
  ScoreUpdated,
  SubmissionSucceeded,
  TurnEnded,
  TurnStarted,
  TurnTimeRemainingChanged,
} from './type'

/* eslint-disable */
class GameInfo {
  gameId: string
  myPlayerId: number
  fieldCards: Card[]

  players: PlayerInfo[]

  currentPlayerId: number
  turn: number
  turnTimeRemaining: number

  constructor(gameId: string, playerId: number) {
    this.gameId = gameId
    this.myPlayerId = playerId
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

  private handleGameReady(event: GameReady): void {}
  private handleGameStarted(event: GameStarted): void {}
  private handleTurnStarted(event: TurnStarted): void {}
  private handleCardsUpdated(event: CardsUpdated): void {}
  private handleTurnTimeRemainingChanged(event: TurnTimeRemainingChanged): void {}
  private handleSubmissionSucceeded(event: SubmissionSucceeded): void {}
  private handleScoreUpdated(event: ScoreUpdated): void {}
  private handleTurnEnded(event: TurnEnded): void {}
  private handleGameEnded(event: GameEnded): void {}
}
