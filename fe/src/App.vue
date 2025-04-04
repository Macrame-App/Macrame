<template>
  <div class="app-background">
    <img src="./assets/img/bg-gradient.svg" aria-hidden="true" />
    <img src="@/assets/img/Macrame-Logo-white.svg" class="logo" aria-hidden="true" />
  </div>
  <MainMenu />
  <RouterView />
</template>

<script setup>
import MainMenu from '@/components/base/MainMenu.vue'
import { onMounted } from 'vue'
import { RouterView } from 'vue-router'
import { useDeviceStore } from './stores/device'

const device = useDeviceStore()

onMounted(() => {
  // Setting device uuid from localstorage
  // If not present in LocalStorage a new uuidV4 will be generated
  device.uuid()
})
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
