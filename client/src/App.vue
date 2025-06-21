<template>
  <main>
    <div :class="$style.container">
      <router-view />
    </div>
  </main>
</template>

<style module>
.container {
  max-width: fit-content;
  margin: auto;
}
</style>
