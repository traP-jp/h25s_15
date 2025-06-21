<script setup lang="ts">
const { score, opponent } = defineProps<{
  score: number
  opponent?: boolean
}>()
</script>

<template>
  <div
    class="score-display"
    :style="{
      outlineColor: !opponent ? 'var(--theme-primary)' : 'var(--theme-danger)',
      outlineOffset: '-0.5rem',
      width: !opponent ? '15.4375rem' : '11.75rem',
      height: !opponent ? '10.1875rem' : '7.75rem',
    }"
  >
    <span>
      <span :style="{ fontSize: !opponent ? '5.25rem' : '4rem' }">
        {{ score }}
      </span>
      <span :style="{ fontSize: !opponent ? '3rem' : '2.5rem' }">pt</span>
    </span>
  </div>
</template>

<style scoped>
.score-display {
  display: flex;
  border-radius: 1.25rem;
  justify-content: center;
  outline: 0.5rem dashed;
  align-items: center;
  box-sizing: border-box;
  color: var(--theme-text-white);
}
</style>
