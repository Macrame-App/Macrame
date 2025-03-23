<template>
  <div class="server-overview">
    <AlertComp type="info"> This is a remote device. </AlertComp>

    <div class="mcrm-block block__light grid gap-4">
      <h4 class="text-lg flex gap-4 items-center"><IconServer />Server</h4>
      <p>
        Connected to: <strong>{{ server.host }}</strong>
      </p>
    </div>
  </div>
</template>

<script setup>
import { IconServer } from '@tabler/icons-vue'
import AlertComp from '../base/AlertComp.vue'
import { onMounted, reactive } from 'vue'
import { useDeviceStore } from '@/stores/device'

const device = useDeviceStore()

const server = reactive({
  host: '',
  access: false,
})

onMounted(() => {
  server.host = window.location.host

  device.checkServerAccess()
})

// console.log(window.location.host)
</script>

<style scoped>
@reference "@/assets/main.css";

.server-overview {
  @apply grid
  gap-4
  content-start;
}
</style>
