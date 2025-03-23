<template>
  <div class="button-group">
    <slot />
  </div>
</template>

<script setup>
defineProps({
  variant: String,
})
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
