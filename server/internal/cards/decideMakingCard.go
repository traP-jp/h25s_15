package cards

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
)

// 	どんなカードを作るか決める関数
// func F(ctx context.Context) (cardType string, value string) {}

const operandProbability = 5
const operatorProbability = 3
const itemProbability = 1

func DecideMakingCard(ctx context.Context) (cardType string, value string, err error) {
	randomIntForType := rand.Intn(operandProbability + operatorProbability + itemProbability)
	if randomIntForType < operandProbability {
		cardType = "operand"
		value = strconv.Itoa(rand.Intn(10))
	} else if randomIntForType < operandProbability+operatorProbability {
		cardType = "operator"
		operators := [7]string{"+", "+", "-", "-", "/", "*", "*"}
		randomIndex := rand.Intn(len(operators))
		value = operators[randomIndex]
	} else {
		cardType = "item"
		items := []string{
			"increaseFieldCards",
			"refreshFieldCards",
			"clearOpponentHandCards",
			// "increaseTurnTime", 未実装
			"increaseHandCardsLimit",
		}
		randomIndex := rand.Intn(len(items))
		value = items[randomIndex]
	}
	if cardType == "" || value == "" {
		return cardType, value, fmt.Errorf("couldn't decide making card")
	}
	return cardType, value, nil
}
