<script setup lang="ts">
const props = defineProps<{
  is_mine: boolean
  now_score: number
}>()
</script>

<template>
  <div
    class="score-display"
    :style="{
      border: props.is_mine
        ? '0.5rem dashed var(--theme-primary, #5980ff)'
        : '0.5rem dashed var(--theme-danger, #dd333c)',
      width: props.is_mine ? '15rem' : '12rem',
      height: props.is_mine ? '10rem' : '8rem',
    }"
  >
    <span
      class="score-value"
      :style="{
        fontSize: props.is_mine ? '6rem' : '4.8rem',
      }"
    >
      {{ now_score }}
    </span>

    <span
      class="score-unit"
      :style="{
        fontSize: props.is_mine ? '4rem' : '3.2rem',
      }"
      >pt</span
    >
  </div>
</template>

<style scoped>
.score-display {
  /* 背景色に定義した変数を使用 */
  background-color: var(--score-box-background, #30384d);

  /* 点線ボーダーに定義した変数を使用 */

  border-radius: 0.75rem;

  display: flex;
  justify-content: center;
  align-items: center;
  box-sizing: border-box;
}

.score-value {
  font-family: 'Inter', 'Helvetica Neue', Arial, sans-serif;
  color: var(--theme-text-white, #ffffff); /* 文字色に定義した変数を使用 */
  line-height: 1;
  margin-right: 0.5rem;
}

.score-unit {
  font-family: 'Inter', 'Helvetica Neue', Arial, sans-serif;
  color: var(--theme-text-white, #ffffff); /* 単位の文字色にも定義した変数を使用 */
  line-height: 1;
  padding-bottom: 0.25rem;
}
</style>
