<template>
  <div class="score-display">
    <div class="user-panel user1">
      <div class="avatar"></div>
      <span class="username">user1</span>
      <span class="score">20</span>
    </div>
    <div class="vs-separator">VS</div>
    <div class="user-panel user2">
      <span class="score">10</span>
      <span class="username">user2</span>
      <div class="avatar"></div>
    </div>
  </div>
</template>

<style scoped>
.score-display {
  display: flex;
  width: 100vw;
  height: 9.25rem;
  align-items: center;
  color: var(--theme-text-white);
  gap: 2.375rem;
}

.user-panel {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 3.125rem;
  padding: 1.4375rem 1.875rem;
}

.user-panel.user1 {
  justify-content: end;
  background: linear-gradient(to right, var(--theme-background), var(--theme-primary));
}

.user-panel.user2 {
  background: linear-gradient(to left, var(--theme-background), var(--theme-danger));
}

.avatar {
  width: 6.0625rem;
  height: 6.0625rem;
  background-color: var(--theme-text-white);
  border-radius: 50%;
}

.username {
  font-size: 3rem;
  white-space: nowrap;
}

.score {
  font-size: 5.25rem;
}

.vs-separator {
  font-size: 4rem;
}
</style>
