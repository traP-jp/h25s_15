<script setup lang="ts">
import CommonButton from '@/components/CommonButton.vue'
import RankingRow from '@/components/RankingRow.vue'
import { ref, onMounted } from 'vue'

const httpBaseUrl = import.meta.env.VUE_APP_HTTP_BASEURL || 'http://localhost:8080'

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

rankingInfo.value = {
  count: 6,
  ranking: [
    {
      name: 'ikura-hamu',
      iconUrl: 'https://q.trap.jp/api/v3/public/icon/ikura-hamu',
      wins: 5,
      losses: 2,
      totalScore: 100,
    },
    {
      name: 'player1',
      iconUrl: 'https://q.trap.jp/api/v3/public/icon/player1',
      wins: 3,
      losses: 4,
      totalScore: 80,
    },
    {
      name: 'player2',
      iconUrl: 'https://q.trap.jp/api/v3/public/icon/player2',
      wins: 5,
      losses: 2,
      totalScore: 100,
    },
    {
      name: 'player3',
      iconUrl: 'https://q.trap.jp/api/v3/public/icon/player3',
      wins: 3,
      losses: 4,
      totalScore: 80,
    },
    {
      name: 'player4',
      iconUrl: 'https://q.trap.jp/api/v3/public/icon/player4',
      wins: 5,
      losses: 2,
      totalScore: 100,
    },
    {
      name: 'irinoirino',
      iconUrl: 'https://q.trap.jp/api/v3/public/icon/irinoirino',
      wins: 3,
      losses: 4,
      totalScore: 80,
    },
  ],
}

myInfo.value = {
  name: 'irinoirino',
  iconUrl: 'https://q.trap.jp/api/v3/public/icon/irinoirino',
}

findMyRanking()

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
  <div class="game-title">（タイトル）</div>
  <CommonButton class="game-start">ゲームを始める</CommonButton>
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
.game-title {
  font-size: 8rem;
  text-align: center;
  display: block;
  padding-top: 5rem;
  margin-bottom: 4rem;
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
