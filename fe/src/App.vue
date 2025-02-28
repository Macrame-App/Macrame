<template>
  <div class="app-background">
    <img class="bg2" src="./assets/bg_2.jpg" aria-hidden="true" />
    <img class="bg1" src="./assets/bg_1.jpg" aria-hidden="true" />
  </div>
  <MainMenu />
  <RouterView />
</template>

<script setup>
import MainMenu from '@/components/base/MainMenu.vue'
import { RouterView } from 'vue-router'

console.log('app.vue loaded')
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

  .bg1 {
    @apply mix-blend-multiply;
  }

  .bg2 {
    @apply grayscale;
  }

  &::after {
    @apply content-['']
    absolute
    inset-0
    bg-radial
    from-sky-400/40
    to-black/80
    mix-blend-multiply
    pointer-events-none;
  }
}
</style>
