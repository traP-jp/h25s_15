<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import { GameInfo } from '../lib/gameLogic'
import { computed, ref } from 'vue'
import GameCard from '@/components/GameCard.vue'
import ParenthesisButtons from '@/components/ParenthesisButtons.vue'
import HandCards from '@/components/HandCards.vue'
import FieldArea from '@/components/FieldArea.vue'
import CommonButton from '@/components/CommonButton.vue'
import TurnTimer from '@/components/TurnTimer.vue'
import ExpressionCards from '@/components/ExpressionCards.vue'
import ScoreBoard from '@/components/ScoreBoard.vue'
import HandCardCounter from '@/components/HandCardCounter.vue'
import { useGameEvent } from '@/composables/useGameEvent'

const routes = useRoute()
const router = useRouter()
const gameId = routes.params.gameId as string
const gameState = ref(new GameInfo(gameId))

const httpBaseUrl = import.meta.env.VITE_HTTP_BASEURL || 'http://localhost:8080'
const wsBaseUrl = import.meta.env.VITE_WS_BASEURL || 'ws://localhost:8080'
const gameWsUrl = `${wsBaseUrl}/games/${gameId}/ws`

useGameEvent(gameWsUrl, (event) => {
  gameState.value.onEvent(event)
  if (event.type == 'gameEnded') {
    router.replace({
      name: 'result',
      params: { gameId },
    })
  }
})

function pickCard(cardId: string) {
  fetch(`${httpBaseUrl}/games/${gameId}/field/${cardId}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ cardId: cardId }),
  })
}

function useCard(cardId: string) {
  const card = gameState.value.players[gameState.value.myPlayerId].cards.find(
    (card) => card.id === cardId
  )
  if (card === undefined) {
    console.error(`Card with id ${cardId} not found`)
    return
  }
  if (card.type === 'item') {
    // アイテムを使用する
    fetch(`${httpBaseUrl}/games/${gameId}/items`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ cardId: cardId }),
    })
  } else {
    // 式に追加する
    gameState.value.useCard(cardId)
  }
}

function clearHandCards() {
  fetch(`${httpBaseUrl}/games/${gameId}/clear`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
  })
}

function addOperator(operation: '(' | ')') {
  gameState.value.addOperator(operation)
}

function deleteExpression() {
  gameState.value.deleteExpr()
}

function submitExpression() {
  const player = gameState.value.players[gameState.value.myPlayerId]
  if (player.expression.length === 0) {
    console.error('Expression is empty')
    return
  }

  fetch(`${httpBaseUrl}/games/${gameId}/submissions`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      expression: player.expression,
      cards: player.cards.map((card) => card.id),
    }),
  })
}

const myPlayer = computed(() => {
  const myPlayer = gameState.value.players.find(({ id }) => id == gameState.value.myPlayerId)
  if (!myPlayer) throw new Error('No my player found')

  return myPlayer
})
const opponentPlayer = computed(() => {
  const opponentPlayer = gameState.value.players.find(({ id }) => id != gameState.value.myPlayerId)
  if (!opponentPlayer) throw new Error('No opponent player found')

  return opponentPlayer
})
</script>

<template>
  <div class="game-container">
    <div class="opponent-container">
      <HandCards :cards="opponentPlayer.cards" card-size="small">
        <GameCard v-for="handCard in opponentPlayer.cards" size="small" :key="handCard.id">
          {{ handCard.value }}
        </GameCard>
      </HandCards>
      <div :style="{ flex: 1 }" />
      <div class="opponent-expression">{{ opponentPlayer.expression }}</div>
      <ScoreBoard opponent :score="opponentPlayer.score" />
    </div>

    <div class="field-container">
      <div :style="{ flex: 1 }" />
      <FieldArea>
        <GameCard
          v-for="fieldCard in gameState.fieldCards"
          size="large"
          :key="fieldCard.id"
          :onClick="() => pickCard(fieldCard.id)"
        >
          {{ fieldCard.value }}
        </GameCard>
      </FieldArea>
      <div class="turn-timer-container" :style="{ flex: 1 }">
        <TurnTimer
          :max_value="gameState.currentTurnTimeLimit"
          :now_value="gameState.turnTimeRemaining"
          :turn="gameState.turnTotal - gameState.turn + 1"
        />
      </div>
    </div>

    <div class="my-hand-container">
      <div :style="{ flex: 1 }">
        <div class="my-hand-info">
          <HandCardCounter :current-count="myPlayer.cards.length" :maxCount="myPlayer.handsLimit" />
          <CommonButton @click="clearHandCards" theme="danger">Clear ( -3pt )</CommonButton>
        </div>
      </div>
      <HandCards :cards="myPlayer.cards" card-size="medium">
        <GameCard
          v-for="handCard in myPlayer.cards"
          size="medium"
          :key="handCard.id"
          :onClick="() => useCard(handCard.id)"
          :selected="myPlayer.expressionCards.includes(handCard)"
        >
          {{ handCard.value }}
        </GameCard>
      </HandCards>
      <div :style="{ flex: 1 }" />
    </div>

    <div class="my-expression-container">
      <ParenthesisButtons @left="addOperator('(')" @right="addOperator(')')"></ParenthesisButtons>
      <div class="my-expression">
        <ExpressionCards @delete="deleteExpression" @submit="submitExpression">
          <GameCard v-for="card in myPlayer.expression" :key="card">
            {{ card }}
          </GameCard>
        </ExpressionCards>
        <ScoreBoard :score="myPlayer.score" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.game-container {
  padding: 1.375rem;
  display: flex;
  flex-direction: column;
  width: 100vw;
  align-items: center;
}

.opponent-container {
  width: 100%;
  height: 14.5rem;
  display: flex;
  gap: 1.125rem;
}

.opponent-expression {
  display: flex;
  height: 7.8125rem;
  width: 58.35rem;
  color: var(--theme-text-white);
  background: var(--theme-surface);
  font-size: 2.5rem;
  justify-content: center;
  align-items: center;
}

.field-container {
  width: 100%;
  display: flex;
}

.turn-timer-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.my-hand-container {
  padding: 2.8125rem 0 1.5rem 0;
  width: 100%;
  display: flex;
  gap: 5.125rem;
}

.my-hand-container > *:first-child {
  display: flex;
  justify-content: end;
}

.my-hand-info {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  text-align: right;
}

.my-expression-container {
  display: flex;
  flex-direction: column;
  gap: 1.3125rem;
}

.my-expression {
  display: flex;
  gap: 1.1875rem;
  height: 10.1875rem;
  align-items: center;
}
</style>
