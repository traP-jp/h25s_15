<script setup lang="ts">
import { computed } from 'vue'
const props = defineProps<{
  max_value: number
  now_value: number
  turn: number
}>()

const percent = computed(() => {
  if (props.max_value === 0) return 0
  const value = (props.now_value / props.max_value) * 100
  return Math.max(0, Math.min(100, value))
})

const radius = 45
const center = 60
const circumference = 2 * Math.PI * radius
const dashOffset = computed(() => {
  const offset = circumference - (circumference * percent.value) / 100
  return offset < 0 ? 0 : offset
})
</script>

<template>
  <div class="progress-container">
    <svg class="progress-ring" width="100%" height="100%" viewBox="0 0 120 120">
      <circle
        class="progress-ring__progress"
        :r="radius"
        :cx="center"
        :cy="center"
        :stroke-dasharray="circumference"
        :stroke-dashoffset="dashOffset"
      />
    </svg>
    <div class="progress-text">#{{ turn }}</div>
  </div>
</template>

<style scoped>
.progress-container {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 14.5rem;
  height: 14.5rem;
}

.progress-ring {
  transform: rotate(-90deg);
}

.progress-ring__progress {
  fill: none;
  stroke: var(--theme-primary);
  stroke-width: 10;
  transition: stroke-dashoffset 0.4s ease;
}

.progress-text {
  position: absolute;
  font-size: 4rem;
  color: var(--theme-text-white);
  font-weight: 600;
}
</style>
