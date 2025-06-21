<script setup lang="ts">
defineProps<{ myExpressions: string[]; opponentExpressions: string[] }>()
</script>

<template>
  <div class="exp-box">
    <div class="left-exp-box">
      <span v-for="(exp, index) in myExpressions" :key="index">{{ exp }}</span>
    </div>
    <div class="right-exp-box">
      <span v-for="(exp, index) in opponentExpressions" :key="index">{{ exp }}</span>
    </div>
  </div>
</template>

<style scoped>
.exp-box {
  min-height: 24.875rem;
  width: 102.6875rem;
  display: flex;
  background-color: var(--theme-surface);
  gap: 9.0625rem;
}
.left-exp-box {
  margin: 2.5rem 0;
  flex: 1;
  display: flex;
  background-color: rgba(0, 0, 0, 0);
  text-align: right;
  gap: 0.625rem;
  font-size: 3.25rem;
  flex-direction: column;
}
.right-exp-box {
  margin: 2.5rem 0;
  flex: 1;
  display: flex;
  background-color: rgba(0, 0, 0, 0);
  text-align: left;
  gap: 0.625rem;
  font-size: 3.25rem;
  flex-direction: column;
}
</style>
