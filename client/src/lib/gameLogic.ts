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

export class GameInfo {
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
      { id: 0, score: 0, handsLimit: 0, cards: [], expression: '', expressionCards: [] },
      { id: 1, score: 0, handsLimit: 0, cards: [], expression: '', expressionCards: [] },
    ]
    this.currentPlayerId = -1
    this.turn = 0
    this.turnTimeRemaining = 0
  }

  useCard(cardId: string): void {
    const player = this.players[this.myPlayerId]
    const cardIndex = player.cards.findIndex((card) => card.id === cardId)
    if (cardIndex === -1) return
    const card = player.cards[cardIndex]
    player.expressionCards.push(card)
    player.expression += card.value
  }
  addOperator(operator: '(' | ')'): void {
    const player = this.players[this.myPlayerId]
    player.expression += operator
  }
  deleteExpr(): void {
    const player = this.players[this.myPlayerId]
    player.expression = ''
    player.expressionCards = []
  }

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

  private handleGameReady(event: GameReady): void {
    this.fieldCards = event.fieldCards

    this.players[0].id = 0
    this.players[0].cards = event.player0
    this.players[0].handsLimit = event.player0HandsLimit
    this.players[0].score = event.player0Score

    this.players[1].id = 1
    this.players[1].cards = event.player1
    this.players[1].handsLimit = event.player1HandsLimit
    this.players[1].score = event.player1Score

    this.currentPlayerId = event.currentPlayerId
    this.turn = 0
  }

  private handleGameStarted(event: GameStarted): void {
    this.currentPlayerId = event.currentPlayerId
    this.turn = event.turn
  }

  private handleTurnStarted(event: TurnStarted): void {
    this.currentPlayerId = event.currentPlayerId
    this.turn = event.turn
    this.turnTimeRemaining = event.turnTimeRemaining
  }

  private handleCardsUpdated(event: CardsUpdated): void {
    this.fieldCards = event.fieldCards

    this.players[0].cards = event.player0
    this.players[0].handsLimit = event.player0HandsLimit

    this.players[1].cards = event.player1
    this.players[1].handsLimit = event.player1HandsLimit
  }

  private handleTurnTimeRemainingChanged(event: TurnTimeRemainingChanged): void {
    this.turnTimeRemaining = event.remainingSeconds
    this.currentPlayerId = event.currentPlayerId
  }

  private handleSubmissionSucceeded(event: SubmissionSucceeded): void {
    const player = this.players[event.playerId]
    player.expressionCards = []
    if (event.playerId === this.myPlayerId) {
      player.expression = ''
    } else {
      player.expression = event.expression
    }
  }

  private handleScoreUpdated(event: ScoreUpdated): void {
    this.players[0].score = event.player0
    this.players[1].score = event.player1
  }

  private handleTurnEnded(event: TurnEnded): void {
    this.currentPlayerId = event.nextPlayerId
    this.turn = event.nextTurn ?? this.turn + 1
  }

  private handleGameEnded(event: GameEnded): void {
    // やることないよ
    console.log('Game ended:', event)
  }
}
