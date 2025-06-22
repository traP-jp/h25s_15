<script setup lang="ts">
import CommonButton from '@/components/CommonButton.vue'
import ExpressionResult from '@/components/ExpressionResult.vue'
import PointResult from '@/components/PointResult.vue'

import { ref, onMounted } from 'vue'
const props = defineProps<{
  gameId: string
}>()

type MeInfo = {
  name: string
  iconUrl: string
}

type ResultInfo = {
  gameId: string
  player0Name: string
  player1Name: string
  player0Score: number
  player1Score: number
  player0SuccessExpressions: string[]
  player1SuccessExpressions: string[]
}

const meInfo = ref<MeInfo | null>(null)
const resultInfo = ref<ResultInfo | null>(null)
const isWin = ref<'VICTORY' | 'LOSE' | 'DRAW' | null>(null)

const judgeResult = () => {
  if (!meInfo.value || !resultInfo.value) return
  const isPlayer0 = resultInfo.value.player0Name === meInfo.value.name
  const myScore = isPlayer0 ? resultInfo.value.player0Score : resultInfo.value.player1Score
  const oppScore = isPlayer0 ? resultInfo.value.player1Score : resultInfo.value.player0Score

  if (myScore > oppScore) isWin.value = 'VICTORY'
  else if (myScore < oppScore) isWin.value = 'LOSE'
  else isWin.value = 'DRAW'
}

onMounted(async () => {
  try {
    //自分のnameを取得
    const meRes = await fetch('/users/me')
    if (!meRes.ok) throw new Error('ユーザー情報の取得に失敗しました')
    meInfo.value = await meRes.json()

    //ゲームの結果を取得
    const resultRes = await fetch('/games/' + props.gameId + '/results')
    if (!resultRes.ok) throw new Error('ゲーム結果の取得に失敗しました')
    resultInfo.value = await resultRes.json()

    //勝敗判定
    judgeResult()
  } catch (error) {
    console.error(error)
  }
})

const shareText = 'いい感じの文字列'
function share_traq() {
  const url = `https://q.trap.jp/share-target?text=${shareText}`
  window.location.href = url
  return
}

function to_home() {
  router.push('/')
}
</script>

<template>
  <div v-if="isWin" class="isJugded">{{ isWin }}</div>

  <PointResult
    v-if="resultInfo"
    :user1="{ name: resultInfo.player0Name, score: resultInfo.player0Score }"
    :user2="{ name: resultInfo.player1Name, score: resultInfo.player1Score }"
  />

  <div class="button-container">
    <CommonButton size="large" theme="primary" variant="filled">traQでシェア</CommonButton>
    <CommonButton size="large" theme="secondary" variant="outline">ホームに戻る</CommonButton>
  </div>

  <ExpressionResult
    v-if="resultInfo"
    :myExpressions="resultInfo.player0SuccessExpressions"
    :opponentExpressions="resultInfo.player1SuccessExpressions"
    class="resultField"
  ></ExpressionResult>
</template>

<style scoped>
.isJugded {
  font-size: 6rem;
  padding-top: 8.375rem;
  padding-bottom: 1.25rem;
  display: block;
  text-align: center;
  color: var(--theme-primary);
}
.button-container {
  width: 45.125rem;
  height: 6.125rem;
  display: flex;
  gap: 2.5rem;
  margin: 5.2rem auto;
}
.resultField {
  margin: 0 auto;
}
</style>
