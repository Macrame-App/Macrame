<template>
  <div class="app-background">
    <img src="./assets/bg-gradient.svg" aria-hidden="true" />
    <img src="@/assets/Macrame-Logo-white.svg" class="logo" aria-hidden="true" />
  </div>
  <MainMenu />
  <RouterView />
</template>

<script setup>
import MainMenu from '@/components/base/MainMenu.vue'
import { RouterView } from 'vue-router'
</script>

<style scoped>
@reference "@/assets/main.css";

.app-background {
  @apply fixed
  inset-0
  size-full
  overflow-hidden
  pointer-events-none
  opacity-40
  z-[-1];

  img {
    @apply absolute
    size-full
    object-cover;
  }

  .logo {
    @apply absolute
    top-[10%]
    left-[10%]
    scale-[1.8]
    p-28
    opacity-35
    mix-blend-overlay;
  }
}
</style>
