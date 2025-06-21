<script setup lang="ts">
defineProps<{
  currentCount: number
  maxCount: number
}>()
</script>

<template>
  <div>
    <span class="label">Hand card</span>
    <div class="count">
      <span :class="['current-count', { 'reach-max': currentCount >= maxCount }]">
        {{ currentCount }}</span
      >/{{ maxCount }}
    </div>
  </div>
</template>

<style scoped>
.label {
  font-size: 1.75rem;
}

.count {
  font-size: 4rem;
}

.current-count {
  font-size: 5.25rem;
  color: var(--theme-tertiary);
}

.current-count.reach-max {
  color: var(--theme-danger);
}
</style>
