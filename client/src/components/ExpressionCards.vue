<script setup lang="ts">
import CommonButton from './CommonButton.vue'

const emits = defineEmits<{ (e: 'delete'): void; (e: 'submit'): void }>()
</script>

<template>
  <div class="expression-card-container">
    <slot />
    <div class="expression-buttons">
      <CommonButton
        @click="() => emits('delete')"
        size="small"
        variant="outline"
        theme="danger"
        :style="{
          flex: 1,
          padding: '0.525rem',
          height: '3.0625rem',
        }"
      >
        ←Delete
      </CommonButton>
      <CommonButton @click="() => emits('submit')" size="small">Submit!➤</CommonButton>
    </div>
  </div>
</template>

<style scoped>
.expression-card-container {
  display: flex;
  gap: 2.5rem;
  width: 98.0625rem;
  height: 10.1875rem;
  padding: 0.9375rem;
  justify-content: center;
  align-items: center;
  background: var(--theme-surface);
}

.expression-buttons {
  display: flex;
  gap: 0.4375rem;
  flex-direction: column;
}
</style>
