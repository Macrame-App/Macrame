<template>
  <div id="devices-view" class="panel">
    <h1 class="panel__title">
      Devices <span class="text-sm">{{ isLocal() ? 'remote' : 'servers' }}</span>
    </h1>
    <div class="panel__content">
      <DeviceOverview v-if="!isLocal()" />
      <ServerOverview v-else />
    </div>
  </div>
</template>

<script setup>
import DeviceOverview from '@/components/devices/DeviceOverview.vue'
import ServerOverview from '@/components/devices/ServerOverview.vue'
import { isLocal } from '@/services/ApiService'
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
