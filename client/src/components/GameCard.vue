<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  selected?: boolean
  disabled?: boolean
  size?: 'small' | 'medium' | 'large'
  onClick?: () => void // あるかないかでスタイルを分けたいのでコールバックにする
}>()

const onClick = (event: MouseEvent) => {
  event.preventDefault()
  props.onClick?.()
}

const size = computed(() => props.size || 'medium')
const sizeValue = computed(() =>
  size.value === 'small' ? '4.0625rem' : size.value === 'medium' ? '6.259375rem' : '7.625rem'
)
</script>

<template>
  <button
    class="card"
    :disabled="props.selected || props.disabled"
    :class="{ clickable: props.onClick != null }"
    :style="{
      backgroundColor: props.selected ? '#636363' : '#FFFFFF',
      width: sizeValue,
      fontSize: sizeValue,
    }"
    @click="onClick"
  >
    <slot />
  </button>
</template>

<style scoped>
.card {
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  color: var(--theme-text-black);
  border-radius: 10%;
  display: flex;
  justify-content: center;
  align-items: center;
  border: none;
  aspect-ratio: 122 / 163;

  transition: all ease-in 100ms;
}

.card.clickable:hover {
  transform: scale(108%);
}

.card.clickable:active {
  transform: none;
}
</style>
