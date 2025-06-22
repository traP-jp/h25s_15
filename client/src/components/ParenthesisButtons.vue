<script setup lang="ts">
const emits = defineEmits<{ (e: 'left'): void; (e: 'right'): void }>()
</script>

<template>
  <div class="paren-row">
    <button class="paren" @click="() => emits('left')">(</button>
    <button class="paren" @click="() => emits('right')">)</button>
  </div>
</template>

<style scoped>
.paren-row {
  display: flex;
  justify-content: center;
  gap: 3.125rem;
}
.paren {
  background: var(--theme-secondary);
  color: var(--theme-text-black);
  font-size: 3rem;
  width: 8.0625rem;
  height: 4.375rem;
  border-radius: 3.125rem;
  border: none;
  font-weight: 400;
  padding: 0.3125rem 3.125rem;
}
</style>
