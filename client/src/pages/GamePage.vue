<script setup lang="ts"></script>

<template>
  <h1>this is GamePage</h1>
  <div class="paren-row">
    <button class="paren">(</button>
    <button class="paren">)</button>
  </div>
</template>


<style>
  .paren-row {
    display: flex;
    justify-content: center;
    gap: 24px;
    margin-top: 6px;
  }
  .paren {
    background: #b8e6d1;
    color: #232323;
    font-size: 2.1rem;
    width: 62px;
    height: 46px;
    border-radius: 22px;
    border: none;
    font-weight: 500;
    margin: 0 4px;
    cursor: pointer;
    transition: background 0.2s;
  }
  .paren:hover {
    background: #9ad2bc;
  }
</style>