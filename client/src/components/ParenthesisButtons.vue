<template>
  <div class="paren-row">
    <button class="paren">(</button>
    <button class="paren">)</button>
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
  border-radius: 100%;
  border: none;
  font-weight: 400;
  margin: 0 0.25rem;
}
</style>
