<script setup lang="ts">
const { cardSize = 'medium' } = defineProps<{
  cardSize?: 'small' | 'medium'
}>()
</script>

<template>
  <div class="hand-cards" :style="{ gap: cardSize == 'medium' ? '1.2rem' : '0.75rem' }">
    <slot />
  </div>
</template>

<style scoped>
.hand-cards {
  display: grid;
  grid-template-rows: repeat(2, min-content);
  grid-auto-flow: column;
}

.hand-card-rows {
  display: flex;
  flex-direction: column;
}
</style>
