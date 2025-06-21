<script setup lang="ts">
import type { Card } from '@/lib/type'
import GameCard from '@/components/GameCard.vue'

const { cards, cardSize = 'medium' } = defineProps<{
  cards: Card[]
  cardSize?: 'small' | 'medium'
}>()

const eachSlice = <T,>(array: T[], size: number): T[][] =>
  new Array(Math.ceil(array.length / size))
    .fill(null)
    .map((_, i) => array.slice(i * size, (i + 1) * size))

const cardColumns = eachSlice(cards, 2)
</script>

<template>
  <div class="hand-cards" :style="{ gap: cardSize == 'medium' ? '1.2rem' : '0.75rem' }">
    <div
      class="hand-card-rows"
      :style="{ gap: cardSize == 'medium' ? '1.2rem' : '0.75rem' }"
      v-for="(cardColumn, i) in cardColumns"
      :key="`hand-card-row-${i}`"
    >
      <GameCard v-for="card in cardColumn" :size="cardSize" :key="card.id">
        {{ card.value }}
      </GameCard>
    </div>
  </div>
</template>

<style scoped>
.hand-cards {
  display: flex;
}

.hand-card-rows {
  display: flex;
  flex-direction: column;
}
</style>
