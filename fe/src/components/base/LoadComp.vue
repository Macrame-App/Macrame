<template>
  <div v-if="loading" class="loading-component">
    <span v-if="text">
      {{ text }}
    </span>
    <IconLoader3 class="duration-1000 animate-spin" />
  </div>
</template>

<script setup>
import { IconLoader3 } from '@tabler/icons-vue'

defineProps({
  loading: Boolean,
  text: String,
})
</script>

<style>
@reference "@/assets/main.css";

&:has(.loading-component) {
  @apply relative;
}

.loading-component {
  @apply absolute
  inset-0
  size-full
  flex gap-2
  flex-col
  justify-center
  items-center
  text-sm
  bg-black/50
  backdrop-blur-md;
}
</style>
