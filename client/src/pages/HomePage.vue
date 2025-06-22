<script setup lang="ts">
import CommonButton from '@/components/CommonButton.vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const httpBaseUrl = import.meta.env.VUE_APP_HTTP_BASEURL || 'http://localhost:8080'
const wsBaseUrl = import.meta.env.VUE_APP_WS_BASEURL || 'ws://localhost:8080'
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
      if (!response.ok) {
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
</script>

<template>
  <div class="game-title">（タイトル）</div>
  <CommonButton class="game-start" @click="gameMatching" :disabled="wating_matching"
    >ゲームを始める</CommonButton
  >
  <div v-if="wating_matching" class="spinner">
    <div class="dot"></div>
    <div class="dot"></div>
    <div class="dot"></div>
  </div>
  <div class="square"></div>
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
  height: 34.0625rem;
  align-items: center;
  margin-top: 4rem;
}
</style>
