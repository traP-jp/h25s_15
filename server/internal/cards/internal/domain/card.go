package domain

import "github.com/google/uuid"

type Card struct {
	ID            uuid.UUID
	Type          CardType
	Value         int
	OwnerPlayerID int          // PlayerID of the owner, 0 if not owned
	Location      CardLocation // Location of the card in the game
}

type CardType string

const (
	CardTypeOperator CardType = "operator"
	CardTypeOperand  CardType = "operand"
	CardTypeItem     CardType = "item"
)

type CardLocation string

const (
	CardLocationHand  CardLocation = "hand"
	CardLocationField CardLocation = "field"
	CardLocationUsed  CardLocation = "used"
)
