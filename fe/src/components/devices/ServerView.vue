<template>
  <div class="device-overview">
    <AlertComp type="info">
      <div class="grid">
        <strong>This is a server!</strong>
        <em>UUID: {{ device.uuid() }} </em>
      </div>
    </AlertComp>

    <div class="mcrm-block block__light flex flex-wrap items-start">
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
            <AlertComp type="success">Linked</AlertComp>
            <ButtonComp variant="danger"> <IconLinkOff />Unlink device </ButtonComp>
          </template>
          <template v-else>
            <AlertComp type="warning">Not linked</AlertComp>
            <ButtonComp variant="primary" @click="startLink(id)">
              <IconLink />Link device
            </ButtonComp>
          </template>
        </div>
      </template>
      <template v-else>
        <div class="grid w-full gap-4">
          <em class="text-slate-300">No remote devices</em>
          <div class="flex justify-end">
            <ButtonComp variant="primary" @click="device.serverGetRemotes()">
              <IconReload />Check for access requests
            </ButtonComp>
          </div>
        </div>
      </template>
      <DialogComp ref="pinDialog">
        <template #content>
          <div class="grid gap-4">
            <h3>Pin code</h3>
            <span class="text-4xl font-mono tracking-wide">{{ remote.pinlink }}</span>
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
  IconDeviceDesktop,
  IconDeviceMobile,
  IconDevicesMinus,
  IconDevicesPlus,
  IconDeviceTablet,
  IconDeviceUnknown,
  IconLink,
  IconLinkOff,
  IconReload,
} from '@tabler/icons-vue'
import ButtonComp from '../base/ButtonComp.vue'
import DialogComp from '../base/DialogComp.vue'

const device = useDeviceStore()

const pinDialog = ref()

const remote = reactive({ devices: [], pinlink: '' })

onMounted(() => {
  device.serverGetRemotes()

  device.$subscribe((mutation, state) => {
    if (Object.keys(state.remote).length) remote.devices = device.remote
  })
})

async function startLink(deviceUuid) {
  const pin = await device.serverStartLink(deviceUuid)

  remote.pinlink = pin
  pinDialog.value.toggleDialog(true)
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
