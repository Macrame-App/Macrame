<template>
  <div class="device-overview">
    <AlertComp variant="info">
      <strong>This is a server!</strong>
      <em>UUID: {{ device.uuid() }} </em>
    </AlertComp>

    <div class="flex flex-wrap items-start gap-4 mcrm-block block__light">
      <h4 class="flex items-center justify-between w-full gap-4 mb-4">
        <span class="flex gap-4" v-if="Object.keys(remote.devices).length > 0">
          <IconDevices />{{ Object.keys(remote.devices).length }}
          {{ Object.keys(remote.devices).length > 1 ? 'Devices' : 'Device' }}
        </span>

        <ButtonComp variant="primary" @click="device.serverGetRemotes()"><IconReload /></ButtonComp>
      </h4>
      <template v-if="Object.keys(remote.devices).length > 0">
        <template v-for="(remoteDevice, id) in remote.devices" :key="id">
          <div
            v-if="id !== device.uuid()"
            class="mcrm-block block__dark block-size__sm w-64 grid !gap-4 content-start"
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
              <AlertComp variant="success">Authorized</AlertComp>
              <ButtonComp variant="danger" @click="unlinkDevice(id)">
                <IconLinkOff />Unlink device
              </ButtonComp>
            </template>

            <template v-else>
              <AlertComp variant="warning">Unauthorized</AlertComp>
              <ButtonComp variant="primary" @click="startLink(id)">
                <IconLink />Link device
              </ButtonComp>
            </template>

            <template v-if="remote.pinlink.uuid == id">
              <AlertComp variant="info">One time pin: {{ remote.pinlink.pin }}</AlertComp>
            </template>
          </div>
        </template>
      </template>

      <template v-else>
        <div class="grid w-full gap-4">
          <em class="text-slate-300">No remote devices</em>
        </div>
      </template>

      <AccordionComp
        class="w-full mt-8 border-t border-t-white/50"
        title="How to connect a device?"
      >
        <div class="grid py-4">
          <ul class="space-y-1">
            <li>
              To connect a device, open <strong>http://{{ server.ip }}:{{ server.port }}</strong> in
              a browser on the device.
            </li>
            <li>Open the menu, and click on <strong>Server.</strong></li>
            <li>
              The device will automatically request access, if you see "Access requested" on the
              device.
            </li>
            <li>
              <div class="inline-flex items-center gap-2">
                Click the
                <span class="p-1 border rounded-sm"><IconReload class="size-4" /></span> to reload
                the devices.
              </div>
            </li>
            <li>
              <div class="inline-flex flex-wrap items-center gap-2 w-fit">
                Click on
                <span class="flex items-center gap-1 p-1 text-sm border rounded-sm">
                  <IconLink class="size-4" /> Link device
                </span>
                to generate a one-time-pin to link the device.
              </div>
            </li>
            <li>
              Enter the pin that is shown on this server in the dialog that will appear on the
              device.
            </li>
            <li>Congratulations! You have linked a device! (Hopefully)</li>
          </ul>
        </div>
      </AccordionComp>

      <DialogComp ref="pinDialog">
        <template #content>
          <div class="grid gap-4">
            <h3>Pin code</h3>
            <span class="font-mono text-4xl tracking-wide">{{ remote.pinlink.pin }}</span>
          </div>
        </template>
      </DialogComp>
    </div>
  </div>
</template>

<script setup>
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
import AccordionComp from '../base/AccordionComp.vue'

const device = useDeviceStore()

const pinDialog = ref()

const server = reactive({
  ip: '',
  port: '',
})
const remote = reactive({ devices: [], pinlink: false })

onMounted(async () => {
  device.serverGetRemotes()

  device.$subscribe((mutation, state) => {
    if (state.remote !== remote.devices) remote.devices = device.remote
  })

  const serverIP = await device.serverGetIP()
  server.ip = serverIP

  server.port = import.meta.env.VITE_MCRM__PORT
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
