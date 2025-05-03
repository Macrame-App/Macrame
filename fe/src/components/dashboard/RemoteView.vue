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
  <div id="remote-dashboard">
    <div id="panels" class="dashboard-block mcrm-block block__light" v-if="server.handshake">
      <div class="icon__container">
        <IconLayoutGrid />
      </div>
      <h4>{{ server.panelCount }} {{ server.panelCount != 1 ? 'Panels' : 'Panel' }}</h4>
      <template v-if="server.panelCount == 0">
        <p><em>No panels found. </em></p>
        <p>Learn how to create a panel <a href="#" target="_blank">here</a>.</p>
      </template>
      <template v-else>
        <p>Start using a panel!</p>
        <ButtonComp variant="danger" href="/panels"> <IconLayoutGrid /> View panels </ButtonComp>
      </template>
    </div>
    <div id="server" class="dashboard-block mcrm-block block__light">
      <div class="icon__container">
        <IconServer />
      </div>
      <h4>Server</h4>
      <template v-if="server.handshake">
        <p>
          Linked with: <strong class="text-center">{{ server.ip }}</strong>
        </p>
        <ButtonComp variant="primary" href="/devices"> <IconServer /> View server</ButtonComp>
      </template>
      <template v-else>
        <p>
          <em>Not linked</em>
        </p>
        <ButtonComp variant="primary" href="/devices"> <IconLink /> Link with server</ButtonComp>
      </template>
    </div>
  </div>
</template>

<script setup>
import { IconLayoutGrid, IconLink, IconServer } from '@tabler/icons-vue'
import { onMounted, reactive } from 'vue'

import ButtonComp from '../base/ButtonComp.vue'

import { useDeviceStore } from '@/stores/device'
import { usePanelStore } from '@/stores/panel'

const device = useDeviceStore()
const panel = usePanelStore()

const server = reactive({
  ip: '',
  handshake: '',
  panelCount: 0,
})

onMounted(async () => {
  const serverIp = await device.serverGetIP()
  server.ip = serverIp

  if (device.key()) server.handshake = true

  device.$subscribe(() => {
    if (device.key()) server.handshake = true
  })

  const panelCount = await panel.getList(true)
  server.panelCount = panelCount
})
</script>

<style scoped>
@reference "@/assets/main.css";

#remote-dashboard {
  @apply grid
  pt-8
  gap-4
  md:w-fit
  h-fit
  content-start;

  &.not__linked #server {
    @apply row-start-1 md:col-start-1;
  }
}
</style>
