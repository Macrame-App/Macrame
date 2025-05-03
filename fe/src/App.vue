<!--
Macrame is a program that enables the user to create keyboard macros and button panels. 
The macros are saved as simple JSON files and can be linked to the button panels. The panels can 
be created with HTML and CSS.

Copyright (C) 2025 Jesse Malotaux

This program is free software: you can redistribute it and/or modify 
it under the terms of the GNU General Public License as published by 
the Free Software Foundation, either version 3 of the License, or 
(at your option) any later version.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
GNU General Public License for more details.

You should have received a copy of the GNU General Public License 
along with this program. If not, see <https://www.gnu.org/licenses/>.
-->

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

onMounted(() => {
  // Setting device uuid from localstorage
  // If not present in LocalStorage a new uuidV4 will be generated
  device.uuid()

  if (!isLocal) appHandshake()

  device.$subscribe(() => {
    if (device.key()) handshake.value = true
  })
})

async function appHandshake() {
  const hsReq = await device.remoteHandshake()
  handshake.value = hsReq
}
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
