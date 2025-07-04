<script setup lang="ts">
import CommonButton from '@/components/CommonButton.vue'
import RankingRow from '@/components/RankingRow.vue'
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const httpBaseUrl = import.meta.env.VITE_HTTP_BASEURL || 'http://localhost:8080'
const wsBaseUrl = import.meta.env.VITE_WS_BASEURL || 'ws://localhost:8080'
const router = useRouter()

const wating_matching = ref(false)

function gameMatching() {
  // POST /games でゲームのリクエスト
  fetch(`${httpBaseUrl}/games`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
  })
    .then((response) => {
      if (!response.ok && response.status !== 400) {
        throw new Error('ゲームの開始に失敗しました')
      } else {
        const ws = new WebSocket(`${wsBaseUrl}/games/ws`)
        wating_matching.value = true
        ws.onmessage = (event) => {
          const data = JSON.parse(event.data) as { gameId: string; playerId: number }
          router.push({ name: 'game', params: { gameId: data.gameId } })
        }
      }
    })
    .catch((error) => {
      console.error('ゲームの開始に失敗しました:', error)
      wating_matching.value = false
    })
}

type MyInfo = {
  name: string
  iconUrl: string
}

type PlayerRanking = {
  name: string
  iconUrl: string
  wins: number
  losses: number
  totalScore: number
}

type RankingInfo = {
  count: number
  ranking: PlayerRanking[]
}

const myInfo = ref<MyInfo | undefined>(undefined)
const rankingInfo = ref<RankingInfo | undefined>(undefined)
const myRank = ref<number | undefined>(undefined)

const findMyRanking = () => {
  if (!myInfo.value || !rankingInfo.value) return
  for (let i: number = 0; i < rankingInfo.value.count; i++) {
    if (rankingInfo.value.ranking[i].name == myInfo.value.name) {
      myRank.value = i + 1
      return
    }
  }
}

onMounted(async () => {
  try {
    //自分のnameを取得
    const meRes = await fetch(`${httpBaseUrl}/users/me`)
    if (!meRes.ok) throw new Error('ユーザー情報の取得に失敗しました')
    myInfo.value = await meRes.json()

    //ランキングを取得
    const rankingRes = await fetch(`${httpBaseUrl}/ranking`)
    if (!rankingRes.ok) throw new Error('ランキングの取得に失敗しました')
    rankingInfo.value = await rankingRes.json()

    findMyRanking()
  } catch (error) {
    console.error(error)
  }
})
</script>

<template>
  <div class="game-title">hasTEN</div>
  <CommonButton class="game-start" @click="gameMatching" :disabled="wating_matching"
    >ゲームを始める</CommonButton
  >
  <div v-if="wating_matching" class="spinner">
    <div class="dot"></div>
    <div class="dot"></div>
    <div class="dot"></div>
  </div>
  <div class="square">
    <div v-if="rankingInfo" class="ranking-list">
      <RankingRow :rank="1" :name="rankingInfo.ranking[0].name"></RankingRow>
      <RankingRow :rank="2" :name="rankingInfo.ranking[1].name"></RankingRow>
      <RankingRow :rank="3" :name="rankingInfo.ranking[2].name"></RankingRow>
      <RankingRow :rank="4" :name="rankingInfo.ranking[3].name"></RankingRow>
      <RankingRow :rank="5" :name="rankingInfo.ranking[4].name"></RankingRow>

      <div v-if="myRank" class="ranking-row" style="border-top: solid 1px #fff; padding-top: 1rem">
        <RankingRow :rank="myRank" :name="rankingInfo.ranking[myRank - 1].name"></RankingRow>
      </div>
    </div>
  </div>
</template>

<style scoped>
.spinner {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1000;
  align-items: center;
  width: 100vw;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  background-color: rgba(128, 128, 128, 0.3);
}

.dot {
  width: 2rem;
  height: 2rem;
  margin: 0 0.5rem;
  background-color: var(--theme-primary);
  border-radius: 50%;
  animation: bounce 1.4s infinite ease-in-out both;
}

.dot:nth-child(1) {
  animation-delay: -0.32s;
}

.dot:nth-child(2) {
  animation-delay: -0.16s;
}

@keyframes bounce {
  0%,
  80%,
  100% {
    transform: scale(0);
  }
  40% {
    transform: scale(1);
  }
}

.game-title {
  font-size: 8rem;
  text-align: center;
  display: block;
  padding-top: 5rem;
  margin-bottom: 4rem;
  color: var(--theme-tertiary);
}

.game-start {
  color: var(--theme-text-white);
  background-color: var(--theme-primary);
  display: flex;
  font-size: 3.25rem;
  width: 28.9375rem;
  height: 7.0625rem;
  align-items: center;
  justify-content: center;
  margin: 0 auto;
}

.square {
  background-color: var(--theme-surface);
  width: 97.625rem;
  min-height: 34.0625rem;
  align-items: center;
  margin-top: 4rem;
}

.ranking-list {
  display: flex;
  flex-direction: column;
  gap: 0.625rem;
  padding-top: 1rem;
  padding-bottom: 1.5rem;
}

.ranking-row {
  display: flex;
  align-items: center;
  gap: 20rem;
}
</style>
