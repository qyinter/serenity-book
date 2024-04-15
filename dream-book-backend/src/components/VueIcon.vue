<template>
  <svg aria-hidden="true" :style="{ width: size, height: size, fill: color }">
    <use :xlink:href="symbolId" />
  </svg>
</template>

<script setup lang="ts">
import { computed } from 'vue'
const props = defineProps({
  name: { type: String, required: true },
  size: { type: String, default: '1rem' },
  color: { type: String, default: '#000' }
})
const symbolId = computed(() => {
  return `#icon-${props.name}`
})
</script>
