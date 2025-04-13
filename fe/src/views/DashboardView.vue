<template>
  <div id="dashboard" class="panel">
    <h1 class="panel__title">Dashboard</h1>

    <div class="panel__content">
      <div class="grid gap-1 opacity-50 h-fit">
        <em v-if="isLocal()">This is the server dashboard.</em>
        <em v-else>This is the server dashboard.</em>
      </div>
      <div v-if="isLocal()">
        <h4>Start: Authenticate a device</h4>
        <ul>
          <li>
            Open <strong>{{ server.ip }}:{{ server.port }}</strong> in a browser on your phone.
            <div class="p-4 qr-container">
              <canvas ref="serverQr"></canvas>
            </div>
          </li>
          <li>Navigate to the Server page. An authentication request will start automatically.</li>
          <li>
            Open the
            <RouterLink to="/devices"> Devices</RouterLink>
            page in this window.
          </li>
          <li>
            <div class="inline-flex items-center gap-2">
              The device will appear, if not click the
              <span class="p-1 border rounded-sm"><IconReload class="size-4" /></span> button.
            </div>
          </li>
        </ul>
        <h4>Using a panel</h4>
        <ul>
          <li>Once authenticated you can access the Test Panel on your device.</li>
          <li>Open the Panels page on your device and click the Test Panel.</li>
        </ul>
      </div>
      <div v-else>
        <h4>Start: Authenticate this device</h4>
        <ul>
          <li>
            Open
            <RouterLink to="/devices"> Server</RouterLink>
            to start.
          </li>
          <li>An authentication request will start automatically.</li>
        </ul>
        <h4>Using a panel</h4>
        <ul>
          <li>
            Once authenticated you can access the
            <RouterLink to="/panel/view/test_panel"> Test Panel</RouterLink>
            on your device.
          </li>
          <li>
            Open the
            <RouterLink to="/panels">Panels page</RouterLink>
            you can edit the panel.
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup>
import { RouterLink } from 'vue-router'
import { isLocal } from '@/services/ApiService'
import { useDeviceStore } from '@/stores/device'
import { onMounted, reactive, ref } from 'vue'
import { IconReload } from '@tabler/icons-vue'

import QRCode from 'qrcode'

const device = useDeviceStore()

const server = reactive({
  ip: '',
  port: '',
  fullPath: '',
})

const serverQr = ref()

onMounted(async () => {
  const serverIP = await device.serverGetIP()
  server.ip = serverIP
  server.port = window.__CONFIG__.MCRM__PORT
  server.fullPath = `http://${server.ip}:${server.port}`

  QRCode.toCanvas(serverQr.value, server.fullPath, (error) => {
    console.log(error)
  })
})
</script>

<style lang="scss" scoped></style>
