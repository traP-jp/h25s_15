<script setup lang="ts">
import CommonButton from '@/components/CommonButton.vue'
import ExpressionResult from '@/components/ExpressionResult.vue'
import PointResult from '@/components/PointResult.vue'
import { useRouter } from 'vue-router'
import { ref, onMounted } from 'vue'

const httpBaseUrl = import.meta.env.VITE_HTTP_BASEURL || 'http://localhost:8080'

const router = useRouter()

const props = defineProps<{
  gameId: string
}>()

type MyInfo = {
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

const myInfo = ref<MyInfo | undefined>(undefined)
const resultInfo = ref<ResultInfo | undefined>(undefined)
const resultStatus = ref<'VICTORY' | 'DEFEAT' | 'DRAW' | undefined>(undefined)

const judgeResult = () => {
  if (!myInfo.value || !resultInfo.value) return
  const isPlayer0 = resultInfo.value.player0Name === myInfo.value.name
  const myScore = isPlayer0 ? resultInfo.value.player0Score : resultInfo.value.player1Score
  const oppScore = isPlayer0 ? resultInfo.value.player1Score : resultInfo.value.player0Score

  if (myScore > oppScore) resultStatus.value = 'VICTORY'
  else if (myScore < oppScore) resultStatus.value = 'DEFEAT'
  else resultStatus.value = 'DRAW'
}

let shareText = ''

onMounted(async () => {
  try {
    //自分のnameを取得
    const meRes = await fetch(`${httpBaseUrl}/users/me`)
    if (!meRes.ok) throw new Error('ユーザー情報の取得に失敗しました')
    myInfo.value = await meRes.json()

    //ゲームの結果を取得
    const resultRes = await fetch(`${httpBaseUrl}/games/${props.gameId}/results`)
    if (!resultRes.ok) throw new Error('ゲーム結果の取得に失敗しました')
    resultInfo.value = await resultRes.json()

    //シェア用のテキストを生成
    if (resultInfo.value) {
      shareText =
        `【${resultInfo.value.player0Name} vs ${resultInfo.value.player1Name}】\n` +
        `結果: ${resultStatus.value}\n` +
        `スコア: ` +
        `${resultInfo.value.player0Name}: ${resultInfo.value.player0Score}\n` +
        `${resultInfo.value.player1Name}: ${resultInfo.value.player1Score}\n` +
        `[詳細はこちら](https://h25s15.trap.show/result/${resultInfo.value.gameId})`
    }

    //勝敗判定
    judgeResult()
  } catch (error) {
    console.error(error)
  }
})

function share_traq() {
  const url = encodeURI(`https://q.trap.jp/share-target?text=${shareText}`)
  window.location.href = url
  return
}

function to_home() {
  router.push('/')
}
</script>

<template>
  <div v-if="resultStatus" class="isJugded">{{ resultStatus }}</div>

  <PointResult
    v-if="resultInfo"
    :user1="{ name: resultInfo.player0Name, score: resultInfo.player0Score }"
    :user2="{ name: resultInfo.player1Name, score: resultInfo.player1Score }"
  />

  <div class="button-container">
    <CommonButton size="large" theme="primary" variant="filled" @click="share_traq"
      >traQでシェア</CommonButton
    >
    <CommonButton size="large" theme="secondary" variant="outline" @click="to_home"
      >ホームに戻る</CommonButton
    >
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
