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
          <template v-if="server.link == 'checking'">
            <div class="grid grid-cols-[2rem_1fr] gap-2">
              <IconReload class="animate-spin" />
              Checking server for link...
            </div>
          </template>
          <template v-if="server.link === false">
            <ButtonComp variant="subtle" @click="pingLink()" class="w-fit">
              <IconReload />Check for server link
            </ButtonComp>
          </template>
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
    <DialogComp ref="linkPinDialog">
      <template #content>
        <div class="grid gap-4">
          <h3>Enter server link pin:</h3>
          <input class="input" type="number" v-model="server.linkPin" />
          <ButtonComp variant="primary" @click="getDeviceKey()">Enter</ButtonComp>
        </div>
      </template>
    </DialogComp>
  </div>
</template>

<script setup>
import { IconKey, IconPlugConnectedX, IconReload, IconServer } from '@tabler/icons-vue'
import AlertComp from '../base/AlertComp.vue'
import ButtonComp from '../base/ButtonComp.vue'
import { onMounted, reactive, ref } from 'vue'
import { useDeviceStore } from '@/stores/device'
import { deviceType, deviceModel, deviceVendor } from '@basitcodeenv/vue3-device-detect'
import DialogComp from '../base/DialogComp.vue'

const device = useDeviceStore()

const linkPinDialog = ref()

const server = reactive({
  host: '',
  status: false,
  access: false,
  link: false,
  linkPin: '',
})

onMounted(async () => {
  server.host = window.location.host

  device.$subscribe((mutation, state) => {
    if (Object.keys(state.server).length) server.status = state.server.status
  })

  const status = await device.remoteCheckServerAccess()
  server.status = status
})

function requestAccess() {
  let deviceName = `${deviceVendor() ? deviceVendor() : 'Unknown'} ${deviceVendor() ? deviceModel() : deviceType()}`

  device.remoteRequestServerAccess(deviceName, deviceType()).then((data) => {
    if (data.data) pingLink()
  })
}

function pingLink() {
  server.link = 'checking'
  device.remotePingLink((link) => {
    server.link = true
    linkPinDialog.value.toggleDialog(true)
    console.log(link, 'opendialog')
  })
}

function getDeviceKey() {
  device.remoteHandshake(server.linkPin)
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
