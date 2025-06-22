<script setup lang="ts">
defineProps<{ rank: number; name: string }>()
</script>

<template>
  <div class="ranking-row">
    <span class="ranking-number">- {{ rank }} -</span>
    <span class="ranking-name">{{ name }}</span>
  </div>
</template>

<style scoped>
.ranking-row {
  display: flex;
  align-items: center;
  gap: 20rem;
}

.ranking-number {
  color: var(--theme-text-white);
  font-size: 4rem;
  width: 17rem;
  text-align: left;
  padding-left: 2.5rem;
}

.ranking-name {
  color: var(--theme-primary);
  font-size: 4rem;
  flex: 1;
  text-align: left;
}
</style>
