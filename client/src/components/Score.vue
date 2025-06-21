<script>
// ロジック
</script>

<template>
  <div class="score-display">
    <span class="score-value">19</span>
    <span class="score-unit">pt</span>
  </div>
</template>

<style scoped>
.score-display {
  /* 背景色に定義した変数を使用 */
  background-color: var(--score-box-background, #30384d);

  /* 点線ボーダーに定義した変数を使用 */
  border: 0.5rem dashed var(--theme-primary, #5980ff);
  border-radius: 0.75rem;

  display: flex;
  justify-content: center;
  align-items: center;
  width: 15rem;
  height: 10rem;
  box-sizing: border-box;
}

.score-value {
  font-family: 'Inter', 'Helvetica Neue', Arial, sans-serif;
  font-size: 6rem;
  color: var(--theme-text-white, #ffffff); /* 文字色に定義した変数を使用 */
  line-height: 1;
  margin-right: 0.5rem;
}

.score-unit {
  font-family: 'Inter', 'Helvetica Neue', Arial, sans-serif;
  font-size: 4rem;
  color: var(--theme-text-white, #ffffff); /* 単位の文字色にも定義した変数を使用 */
  line-height: 1;
  padding-bottom: 0.25rem;
}
</style>
