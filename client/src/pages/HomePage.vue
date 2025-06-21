<script setup lang="ts">
import CommonButton from '@/components/CommonButton.vue'
</script>

<template>
  <div class="game-title">（タイトル）</div>
  <CommonButton class="game-start">ゲームを始める</CommonButton>
  <div class="square"></div>
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
  height: 34.0625rem;
  align-items: center;
  margin-top: 4rem;
}
</style>
