<template>
  <div class="server-overview">
    <AlertComp type="info">
      <div class="grid">
        <strong>This is a remote device.</strong>
        <em>UUID: {{ device.uuid() }} </em>
      </div>
    </AlertComp>

    <div class="mcrm-block block__light grid gap-4">
      <h4 class="text-lg flex gap-4 items-center justify-between">
        <span class="flex gap-4"><IconServer />Server</span>
        <ButtonComp variant="primary" @click="checkServerStatus()"><IconReload /></ButtonComp>
      </h4>

      <p>
        Connected to: <strong>{{ server.host }}</strong>
      </p>

      <!-- Alerts -->
      <AlertComp v-if="server.status === 'authorized'" type="success">Authorized</AlertComp>
      <AlertComp v-if="server.status === 'unlinked'" type="warning">Not linked</AlertComp>
      <AlertComp v-if="server.status === 'unauthorized'" type="info">
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
      <ButtonComp
        variant="danger"
        v-if="server.status === 'authorized'"
        @click="disonnectFromServer()"
      >
        <IconPlugConnectedX />
        Disconnect
      </ButtonComp>
    </div>
    <DialogComp ref="linkPinDialog">
      <template #content>
        <div class="grid gap-4 w-64">
          <h3>Server link pin:</h3>
          <form class="grid gap-4" @submit.prevent="decryptKey()">
            <input
              class="input"
              id="input-pin"
              type="text"
              pattern="[0-9]{4}"
              v-model="server.inputPin"
            />
            <ButtonComp variant="primary">Enter</ButtonComp>
          </form>
        </div>
      </template>
    </DialogComp>
  </div>
</template>

<script setup>
// TODO
// - [Delete local key button]
// - if not local key
// - - if !checkAccess -> requestAccess -> put settings.json (go)
// - - if checkAccess -> pingLink -> check for device.tmp (go)
// - - if [devicePin] -> handshake -> save key local, close dialog, update server status

import { IconKey, IconPlugConnectedX, IconReload, IconServer } from '@tabler/icons-vue'
import AlertComp from '../base/AlertComp.vue'
import ButtonComp from '../base/ButtonComp.vue'
import { onMounted, onUpdated, reactive, ref } from 'vue'
import { useDeviceStore } from '@/stores/device'
import { deviceType, deviceModel, deviceVendor } from '@basitcodeenv/vue3-device-detect'
import DialogComp from '../base/DialogComp.vue'
import { AuthCall, decryptAES } from '@/services/EncryptService'
import axios from 'axios'
import { appUrl } from '@/services/ApiService'

const device = useDeviceStore()

const linkPinDialog = ref()

const server = reactive({
  host: '',
  status: false,
  link: false,
  inputPin: '',
  encryptedKey: '',
  key: '',
})

onMounted(async () => {
  server.host = window.location.host
})

onUpdated(() => {
  if (!server.status) checkServerStatus()
})

async function checkServerStatus(request = true) {
  const status = await device.remoteCheckServerAccess()

  server.status = status

  if (status === 'unlinked' || status === 'unauthorized') {
    if (request) requestAccess()
    return true
  }

  if (!device.key()) {
    server.status = 'unauthorized'
    return true
  }

  const handshake = await device.remoteHandshake(device.key())

  if (handshake) server.key = device.key()
  else {
    device.removeDeviceKey()
    server.status = 'unlinked'
    if (request) requestAccess()
  }

  return true
}

function requestAccess() {
  let deviceName = `${deviceVendor() ? deviceVendor() : 'Unknown'} ${deviceVendor() ? deviceModel() : deviceType()}`

  device.remoteRequestServerAccess(deviceName, deviceType()).then((data) => {
    if (data.data) (server.status = data.data), pingLink()
  })
}

function pingLink() {
  server.link = 'checking'

  device.remotePingLink((encryptedKey) => {
    server.link = true
    server.encryptedKey = encryptedKey

    linkPinDialog.value.toggleDialog(true)
  })
}

async function decryptKey() {
  const decryptedKey = decryptAES(server.inputPin, server.encryptedKey)
  const handshake = await device.remoteHandshake(decryptedKey)

  if (handshake) {
    device.setDeviceKey(decryptedKey)
    server.key = decryptedKey
    linkPinDialog.value.toggleDialog(false)
    server.status = 'authorized'
  }
}

function disonnectFromServer() {
  axios.post(appUrl() + '/device/link/remove', AuthCall({ uuid: device.uuid() })).then((data) => {
    if (data.data) checkServerStatus(false)
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

#input-pin {
}
</style>
