<script setup lang="ts">
import { computed } from 'vue'

type Theme = 'primary' | 'secondary' | 'tertiary' | 'danger'
type Variant = 'filled' | 'outline'
type Size = 'small' | 'medium' | 'large' | 'xl'

const {
  size = 'medium',
  theme = 'primary',
  variant = 'filled',
  disabled,
} = defineProps<{
  size?: Size
  theme?: Theme
  variant?: Variant
  disabled?: boolean
}>()

const emits = defineEmits<{ (e: 'click'): void }>()

const accent = computed(() => {
  const colors = {
    primary: '--theme-primary',
    secondary: '--theme-secondary',
    tertiary: '--theme-tertiary',
    danger: '--theme-danger',
  } satisfies Record<Theme, string>

  return colors[theme]
})

const fontSize = computed(() => {
  const sizes = {
    small: '1.5rem',
    medium: '2.25rem',
    large: '2.5rem',
    xl: '3.25rem',
  } satisfies Record<Size, string>

  return sizes[size]
})
</script>

<template>
  <button
    :class="['button', variant]"
    :style="{
      fontSize,
      '--accent': `var(${accent})`,
    }"
    @click="() => emits('click')"
    :disabled="disabled"
  >
    <slot>Button</slot>
  </button>
</template>

<style scoped>
.button {
  padding: 1.5625rem 3.125rem;
  border: none;
  box-sizing: border-box;
  border-radius: 0.625rem;
  box-shadow: 0 4px 4px rgba(0, 0, 0, 0.1);
}

.filled {
  color: var(--theme-text-white);
  background-color: var(--accent);
}

.filled:hover {
  color: color-mix(in srgb, var(--theme-text-white) 90%, black 10%);
  background-color: color-mix(in srgb, var(--accent) 90%, black 10%);
}

.filled:active {
  color: color-mix(in srgb, var(--theme-text-white) 80%, black 20%);
  background-color: color-mix(in srgb, var(--accent) 80%, black 20%);
  box-shadow: inset 0 0 8px rgba(0, 0, 0, 0.2);
}

.outline {
  color: var(--accent);
  outline: 0.25rem solid var(--accent);
  background-color: transparent;
  outline-offset: -0.25rem;
}

.outline:hover {
  color: color-mix(in srgb, var(--accent) 90%, black 10%);
  outline-color: color-mix(in srgb, var(--accent) 90%, black 10%);
}

.outline:active {
  color: color-mix(in srgb, var(--accent) 80%, black 20%);
  outline-color: color-mix(in srgb, var(--accent) 80%, black 20%);
}

.button:disabled {
  box-shadow: none;
  color: color-mix(in srgb, var(--theme-text-white) 70%, black 30%);
  background-color: color-mix(in srgb, var(--accent) 70%, black 30%);
  outline-color: color-mix(in srgb, var(--accent) 70%, black 30%);
}
</style>
