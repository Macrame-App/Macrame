<template>
  <div
    id="server-dashboard"
    :class="`${server.remoteCount == 0 ? 'no__devices' : 'devices__found'} ${server.macroCount == 0 ? 'no__macros' : 'macros__found'}`"
  >
    <div id="devices" class="dashboard-block mcrm-block block__light">
      <div class="icon__container">
        <IconDevices />
      </div>
      <h4>{{ server.remoteCount }} {{ server.remoteCount != 1 ? 'Devices' : 'Device' }}</h4>
      <template v-if="server.remoteCount == 0">
        <p><em>No devices found.</em></p>
        <ButtonComp variant="primary" href="/devices"> <IconLink /> Link a device</ButtonComp>
      </template>
      <template v-else>
        <p>Unlink a device or add new devices.</p>
        <ButtonComp variant="primary" href="/devices"><IconDevices /> View devices</ButtonComp>
      </template>
    </div>
    <div id="macros" class="dashboard-block mcrm-block block__light">
      <div class="icon__container">
        <IconKeyboard />
      </div>
      <h4>{{ server.macroCount }} {{ server.macroCount != 1 ? 'Macros' : 'Macro' }}</h4>
      <template v-if="server.macroCount == 0">
        <p><em>No macros found.</em></p>
        <ButtonComp variant="secondary" href="/macros"> <IconLayoutGrid /> Create macro</ButtonComp>
      </template>
      <template v-else>
        <p>Edit and view your macros.</p>
        <ButtonComp variant="secondary" href="/macros"><IconKeyboard /> View macros</ButtonComp>
      </template>
    </div>

    <div id="panels" class="dashboard-block mcrm-block block__light">
      <div class="icon__container">
        <IconLayoutGrid />
      </div>
      <h4>{{ server.panelCount }} {{ server.panelCount != 1 ? 'Panels' : 'Panel' }}</h4>
      <template v-if="server.panelCount == 0">
        <p><em>No panels found. </em></p>
        <p>Learn how to create a panel <a href="#" target="_blank">here</a>.</p>
      </template>
      <template v-else>
        <p>Link macros to panels or view a panel.</p>
        <ButtonComp variant="danger" href="/panels"> <IconLayoutGrid /> View panels </ButtonComp>
      </template>
    </div>
  </div>
</template>

<script setup>
import { useDeviceStore } from '@/stores/device'
import { usePanelStore } from '@/stores/panel'
import { IconDevices, IconKeyboard, IconLayoutGrid, IconLink } from '@tabler/icons-vue'
import { onMounted, reactive } from 'vue'
import ButtonComp from '../base/ButtonComp.vue'
import { GetMacroList } from '@/services/MacroService'

const device = useDeviceStore()
const panel = usePanelStore()

const server = reactive({
  ip: '',
  port: '',
  fullPath: '',
  remoteCount: 0,
  macroCount: 0,
  panelCount: 0,
})

onMounted(async () => {
  const serverIP = await device.serverGetIP()
  server.ip = serverIP
  // server.port = window.__CONFIG__.MCRM__PORT
  // server.fullPath = `http://${server.ip}:${server.port}`

  const remoteCount = await device.serverGetRemotes(true)
  server.remoteCount = remoteCount

  const macroCount = await GetMacroList(true)
  server.macroCount = macroCount

  const panelCount = await panel.getList(true)
  server.panelCount = panelCount

  console.log(server)
})
</script>

<style scoped>
@reference "@/assets/main.css";

#server-dashboard {
  @apply grid
  grid-cols-1
  grid-rows-3
  md:grid-cols-3
  md:grid-rows-1
  gap-4
  w-fit
  h-fit
  pt-8;

  &.no__devices #devices {
    @apply row-start-1 md:col-start-1;
  }

  &.no__macros.devices__found #devices {
    @apply row-start-3 md:col-start-3;
  }

  &.devices__found #devices {
    @apply row-start-3 md:col-start-3;
  }

  &.no__devices.no__macros #macros {
    @apply row-start-2 md:col-start-2;
  }

  &.no__macros #macros {
    @apply row-start-1 md:col-start-1;
  }

  &.macros__found #macros {
    @apply row-start-2 md:col-start-2;
  }

  &.no__devices.macros__found #macros {
    @apply row-start-3 md:col-start-3;
  }
}
</style>
