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
  gap: 50px;
}
.paren {
  background: var(--theme-secondary);
  color: var(--theme-text-black);
  font-size: 48px;
  width: 129px;
  height: 70px;
  border-radius: 50px;
  border: none;
  font-weight: 500;
  margin: 0 4px;
}
</style>
