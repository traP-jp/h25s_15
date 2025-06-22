<script setup lang="ts">
import { ref } from 'vue';

const props = defineProps<{
  card: {
    id: string;
    type: string;
    value: string | number;
  }
}>()

const imageUrl = ref<string | undefined>(undefined)

if(props.card.type == "item"){
  imageUrl.value = `/${props.card.value}.png`;
}

</script>

<template>
<span v-if="card.type == 'operand' || card.type == 'operator'">{{card.value}}</span>
<img class="image-size" v-else-if="imageUrl" :src="imageUrl" />
</template>

<style scoped>
.image-size{
  width: 100%;
  height: auto;
}
</style>