<script setup lang="ts">
import { useRoute } from 'vue-router'
import { GameInfo } from '../lib/gameLogic'
import { ref } from 'vue'

const routes = useRoute()
const gameId = routes.params.gameId as string
const gameState = ref(new GameInfo(gameId))

const httpBaseUrl = import.meta.env.VUE_APP_HTTP_BASEURL || 'http://localhost:8080'
const wsBaseUrl = import.meta.env.VUE_APP_WS_BASEURL || 'ws://localhost:8080'

// const ws = new WebSocket(`${wsBaseUrl}/${gameId}`);
// ws.onmessage = (event) => {
//   const data = JSON.parse(event.data);
//   gameinfo.value.onEvent(data);
// };

function pickCard(cardId: string) {
  fetch(`${httpBaseUrl}game/${gameId}/field/${cardId}`, {
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
    fetch(`${httpBaseUrl}game/${gameId}/items`, {
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
  fetch(`${httpBaseUrl}game/${gameId}/clear`, {
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

  fetch(`${httpBaseUrl}game/${gameId}/submissions`, {
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

// 使ってる風
pickCard('')
useCard(wsBaseUrl)
clearHandCards()
deleteExpression()
addOperator('(')
submitExpression()
</script>

<template>
  <h1>this is GamePage</h1>
</template>
