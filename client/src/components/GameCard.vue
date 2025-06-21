<script setup lang="ts">
const props = defineProps<{
  selected?: boolean 
  disabled?: boolean
  size?: 'small' | 'medium' | 'large'
}>()

const emit = defineEmits(['click'])

function onClick() {
  emit('click')
}

const size = props.size || 'medium'
</script>

<template>
  <button
    class="card"
    :disabled="props.selected || props.disabled"
    :style="{
      backgroundColor: props.selected ? '#636363' : '#FFFFFF',
      width: size === 'small' ? '2rem' : size === 'medium' ? '4rem' : '5rem',
      border: 'none',
      aspectRatio: 3 / 4,
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
  border-radius: 4px;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
