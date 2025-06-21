<script setup lang="ts">
import CommonButton from '@/components/CommonButton.vue'
import ExpressionResult from '@/components/ExpressionResult.vue'
import PointResult from '@/components/PointResult.vue'

// import { ref, onMounted } from 'vue'
// const props = defineProps<{
//   gameId: string;
// }>()

// const meInfo = ref()
// const resultInfo = ref()
// const isWin = ref()

// onMounted(async () => {

//   //自分のnameを取得
//   const meRes = await fetch('/users/me')
//   if (meRes.ok) {
//     meInfo.value = await meRes.json()
//   }

//   //ゲームの結果を取得
//   const resultRes = await fetch('/games/:gameId/results')
//   if (resultRes.ok) {
//     resultInfo.value = await resultRes.json()
//   }

//   //勝ち負けを判定し、isWinに代入
//   if(resultInfo.value.player0Name == meInfo.value.name){
//     if(resultInfo.value.player0Score > resultInfo.value.player1Score){
//       isWin.value = "VICTORY";
//     }else if(resultInfo.value.player0Score < resultInfo.value.player1Score){
//       isWin.value = "LOSE";
//     }else{
//       isWin.value = "DRAW"
//     }
//   }else{
//     if(resultInfo.value.player0Score > resultInfo.value.player1Score){
//       isWin.value = "LOSE";
//     }else if(resultInfo.value.player0Score < resultInfo.value.player1Score){
//       isWin.value = "VICTORY";
//     }else{
//       isWin.value = "DRAW"
//     }
//   }
// })
</script>

<template>
  <div class="isJugded">VICTORY</div>

  <PointResult :user1="{name: 'Ponjuice', score: 20}" :user2="{name: 'ikura-hamu', score: 30}" />

  <div class="button-container">
    <CommonButton size="large" theme="primary" variant="filled">traQでシェア</CommonButton>
    <CommonButton size="large" theme="secondary" variant="outline">ホームに戻る</CommonButton>
  </div>

  <ExpressionResult
    :myExpressions="['3+7=10', '2×4+2=10']"
    :opponentExpressions="['10=4+6', '10=2×4+2']"
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
