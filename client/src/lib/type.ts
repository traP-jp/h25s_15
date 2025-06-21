export type Card = {
  id: string
  type: string
  value: string | number
}

export type GameReady = {
  type: 'gameReady'
  fieldCards: Card[]
  playerId: number
  player0: Card[]
  player0HandsLimit: number
  player1: Card[]
  player1HandsLimit: number
  currentPlayerId: number
  player0Score: number
  player1Score: number
  startTime: string
}

export type GameStarted = {
  type: 'gameStarted'
  currentPlayerId: number
  turn: number
}

export type TurnStarted = {
  type: 'turnStarted'
  currentPlayerId: number
  turn: number
  turnTimeRemaining: number
}

export type CardsUpdated = {
  type: 'cardsUpdated'
  fieldCards: Card[]
  player0: Card[]
  player0HandsLimit: number
  player1: Card[]
  player1HandsLimit: number
}

export type TurnTimeRemainingChanged = {
  type: 'turnTimeRemainingChanged'
  currentPlayerId: number
  remainingSeconds: number
}

export type SubmissionSucceeded = {
  type: 'submissionSucceeded'
  playerId: number
  expression: string
  score: number
}

export type ScoreUpdated = {
  type: 'scoreUpdated'
  player0: number
  player1: number
}

export type TurnEnded = {
  type: 'turnEnded'
  nextPlayerId: number
  nextTurn: number | null
}

export type GameEnded = {
  type: 'gameEnded'
  player0: number
  player1: number
}

export type GameEvent =
  | GameReady
  | GameStarted
  | TurnStarted
  | CardsUpdated
  | TurnTimeRemainingChanged
  | SubmissionSucceeded
  | ScoreUpdated
  | TurnEnded
  | GameEnded

export type PlayerInfo = {
  id: number
  score: number
  cards: Card[]
  handsLimit: number
  expression: string
  expressionCards: Card[]
}
