<template>
  <div class="app-background">
    <img src="./assets/img/bg-gradient.svg" aria-hidden="true" />
    <img src="@/assets/img/Macrame-Logo-white.svg" class="logo" aria-hidden="true" />
  </div>
  <MainMenu />
  <RouterView />
  <AlertComp
    v-if="!isLocal && !handshake && route.fullPath !== '/devices'"
    variant="warning"
    :page-wide="true"
    href="/devices"
  >
    <h4>Not authorized!</h4>
    <p>Click here to start authorization and open the "Devices" page on your PC.</p>
  </AlertComp>
</template>

<script setup>
import MainMenu from '@/components/base/MainMenu.vue'
import { onMounted, ref } from 'vue'
import { RouterView, useRoute } from 'vue-router'
import { useDeviceStore } from './stores/device'
import { isLocal } from './services/ApiService'
import AlertComp from './components/base/AlertComp.vue'

const device = useDeviceStore()

const route = useRoute()
const handshake = ref(false)

onMounted(async () => {
  // Setting device uuid from localstorage
  // If not present in LocalStorage a new uuidV4 will be generated
  device.uuid()

  const hsReq = await device.remoteHandshake()
  handshake.value = hsReq

  device.$subscribe(() => {
    if (device.key()) handshake.value = true
  })
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
    opacity-35
    mix-blend-overlay;
  }
}
</style>
