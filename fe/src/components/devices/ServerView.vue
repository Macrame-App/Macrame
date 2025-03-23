<template>
  <div class="device-overview">
    <AlertComp type="info">
      <div class="grid">
        <strong>This is a server.</strong>
        <em>UUID: {{ device.uuid() }} </em>
      </div>
    </AlertComp>

    <div class="mcrm-block block__light flex flex-wrap items-start">
      <!-- {{ remote.devices }} -->
      <div
        class="mcrm-block block__dark block-size__sm w-64 grid !gap-4 content-start"
        v-for="(device, id) in remote.devices"
        :key="id"
      >
        <div class="grid gap-2">
          <h5 class="grid grid-cols-[auto_1fr] gap-2">
            <IconDeviceUnknown v-if="device.settings.type == 'unknown'" />
            <IconDeviceMobile v-if="device.settings.type == 'phone'" />
            <IconDeviceTablet v-if="device.settings.type == 'tablet'" />
            <span class="w-full truncate">
              {{ device.settings.name }}
            </span>
          </h5>
          <em>{{ id }}</em>
        </div>
        <template v-if="device.key">
          <AlertComp type="success">Registered</AlertComp>
          <ButtonComp variant="danger"> <IconDevicesMinus />Unregister device </ButtonComp>
        </template>
        <template v-else>
          <AlertComp type="warning">Not registered</AlertComp>
          <ButtonComp variant="primary"> <IconDevicesPlus />Register device </ButtonComp>
        </template>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive } from 'vue'
import AlertComp from '../base/AlertComp.vue'
import { useDeviceStore } from '@/stores/device'
import {
  IconDeviceMobile,
  IconDevicesMinus,
  IconDevicesPlus,
  IconDeviceTablet,
  IconDeviceUnknown,
} from '@tabler/icons-vue'
import ButtonComp from '../base/ButtonComp.vue'

const device = useDeviceStore()

const remote = reactive({ devices: [] })

onMounted(() => {
  device.getRemoteDevices()

  device.$subscribe((mutation, state) => {
    if (Object.keys(state.remote).length) remote.devices = device.remote
  })
})
</script>

<style scoped>
@reference "@/assets/main.css";

.device-overview {
  @apply grid
  gap-4
  content-start;
}
</style>
