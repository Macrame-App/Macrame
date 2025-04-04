<template>
  <div class="device-overview">
    <AlertComp type="info">
      <div class="grid">
        <strong>This is a server!</strong>
        <em>UUID: {{ device.uuid() }} </em>
      </div>
    </AlertComp>

    <div class="mcrm-block block__light flex flex-wrap items-start gap-4">
      <h4 class="w-full flex gap-4 items-center justify-between mb-4">
        <span class="flex gap-4"> <IconDevices />Remote devices </span>
        <ButtonComp variant="primary" @click="device.serverGetRemotes()"><IconReload /></ButtonComp>
      </h4>
      <!-- {{ Object.keys(remote.devices).length }} -->
      <template v-if="Object.keys(remote.devices).length > 0">
        <div
          class="mcrm-block block__dark block-size__sm w-64 grid !gap-4 content-start"
          v-for="(remoteDevice, id) in remote.devices"
          :key="id"
        >
          <div class="grid gap-2">
            <h5 class="grid grid-cols-[auto_1fr] gap-2">
              <IconDeviceUnknown v-if="remoteDevice.settings.type == 'unknown'" />
              <IconDeviceMobile v-if="remoteDevice.settings.type == 'mobile'" />
              <IconDeviceTablet v-if="remoteDevice.settings.type == 'tablet'" />
              <IconDeviceDesktop v-if="remoteDevice.settings.type == 'desktop'" />
              <span class="w-full truncate">
                {{ remoteDevice.settings.name }}
              </span>
            </h5>
            <em>{{ id }}</em>
          </div>
          <template v-if="remoteDevice.key">
            <AlertComp type="success">Authorized</AlertComp>
            <ButtonComp variant="danger" @click="unlinkDevice(id)">
              <IconLinkOff />Unlink device
            </ButtonComp>
          </template>
          <template v-else>
            <AlertComp type="warning">Unauthorized</AlertComp>
            <ButtonComp variant="primary" @click="startLink(id)">
              <IconLink />Link device
            </ButtonComp>
          </template>
          <template v-if="remote.pinlink.uuid == id">
            <AlertComp type="info">One time pin: {{ remote.pinlink.pin }}</AlertComp>
          </template>
        </div>
      </template>
      <template v-else>
        <div class="grid w-full gap-4">
          <em class="text-slate-300">No remote devices</em>
        </div>
      </template>
      <DialogComp ref="pinDialog">
        <template #content>
          <div class="grid gap-4">
            <h3>Pin code</h3>
            <span class="text-4xl font-mono tracking-wide">{{ remote.pinlink.pin }}</span>
          </div>
        </template>
      </DialogComp>
    </div>
  </div>
</template>

<script setup>
// TODO
// - startLink -> responsePin also in device block
// - startLink -> poll removal of pin file, if removed close dialog, update device list
// - Make unlink work

import { onMounted, reactive, ref } from 'vue'
import AlertComp from '../base/AlertComp.vue'
import { useDeviceStore } from '@/stores/device'
import {
  IconDevices,
  IconDeviceDesktop,
  IconDeviceMobile,
  IconDeviceTablet,
  IconDeviceUnknown,
  IconLink,
  IconLinkOff,
  IconReload,
} from '@tabler/icons-vue'
import ButtonComp from '../base/ButtonComp.vue'
import DialogComp from '../base/DialogComp.vue'
import axios from 'axios'
import { appUrl } from '@/services/ApiService'

const device = useDeviceStore()

const pinDialog = ref()

const remote = reactive({ devices: [], pinlink: false })

onMounted(() => {
  device.serverGetRemotes()

  device.$subscribe((mutation, state) => {
    if (Object.keys(state.remote).length) remote.devices = device.remote
  })
})

async function startLink(deviceUuid) {
  const pin = await device.serverStartLink(deviceUuid)

  remote.pinlink = { uuid: deviceUuid, pin: pin }
  pinDialog.value.toggleDialog(true)

  pollLink()

  setTimeout(() => {
    resetPinLink()
  }, 60000)
}

function pollLink() {
  const pollInterval = setInterval(() => {
    axios.post(appUrl() + '/device/link/poll', { uuid: remote.pinlink.uuid }).then((data) => {
      if (!data.data) {
        clearInterval(pollInterval)
        resetPinLink()
        device.serverGetRemotes()
      }
    })
  }, 1000)
}

function resetPinLink() {
  remote.pinlink = false
  if (pinDialog.value) pinDialog.value.toggleDialog(false)
}

function unlinkDevice(id) {
  axios.post(appUrl() + '/device/link/remove', { uuid: id }).then((data) => {
    if (data.data) device.serverGetRemotes()
  })
}
</script>

<style scoped>
@reference "@/assets/main.css";

.device-overview {
  @apply grid
  gap-4
  content-start;
}
</style>
