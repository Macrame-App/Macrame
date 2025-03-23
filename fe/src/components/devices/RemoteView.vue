<template>
  <div class="server-overview">
    <AlertComp type="info">
      <div class="grid">
        <strong>This is a remote device.</strong>
        <em>UUID: {{ device.uuid() }} </em>
      </div>
    </AlertComp>

    <div class="mcrm-block block__light grid gap-4">
      <h4 class="text-lg flex gap-4 items-center"><IconServer />Server</h4>
      <p>
        Connected to: <strong>{{ server.host }}</strong>
      </p>
      <AlertComp v-if="server.status === 'unauthorized'" type="warning">Not authorized</AlertComp>
      <AlertComp v-if="server.status === 'authorized'" type="success">Authorized</AlertComp>
      <AlertComp v-if="server.status === 'requested'" type="info">
        <div class="grid gap-2">
          <strong>Access requested</strong>
          <p>
            Navigate to <em class="font-semibold">http://localhost:6970/devices</em> on your pc to
            authorize.
          </p>
        </div>
      </AlertComp>

      <ButtonComp
        v-if="server.status === 'unauthorized'"
        variant="primary"
        @click="requestAccess()"
      >
        <IconKey />
        Request access
      </ButtonComp>
      <ButtonComp variant="danger" v-if="server.status === 'authorized'">
        <IconPlugConnectedX />
        Disconnect
      </ButtonComp>
    </div>
  </div>
</template>

<script setup>
import { IconKey, IconPlugConnectedX, IconServer } from '@tabler/icons-vue'
import AlertComp from '../base/AlertComp.vue'
import ButtonComp from '../base/ButtonComp.vue'
import { onMounted, reactive } from 'vue'
import { useDeviceStore } from '@/stores/device'

const device = useDeviceStore()

const server = reactive({
  host: '',
  status: false,
  access: false,
})

onMounted(async () => {
  server.host = window.location.host

  device.$subscribe((mutation, state) => {
    if (Object.keys(state.server).length) server.status = state.server.status
  })

  const status = await device.checkServerAccess()
  server.status = status
})

function requestAccess() {
  const request = device.requestServerAccess()

  request
    .then((data) => {
      console.log('1', data)
    })
    .then((data) => {
      console.log('2', data)
    })
}
</script>

<style scoped>
@reference "@/assets/main.css";

.server-overview {
  @apply grid
  gap-4
  content-start;
}
</style>
