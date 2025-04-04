<template>
  <div id="devices-view" class="panel">
    <h1 class="panel__title">
      Devices <span class="text-sm">{{ isLocal() ? 'remote' : 'servers' }}</span>
    </h1>
    <div class="panel__content grid gap-8">
      <ServerView v-if="isLocal()" />
      <RemoteView v-else />
    </div>
  </div>
</template>

<script setup>
import ServerView from '@/components/devices/ServerView.vue'
import RemoteView from '@/components/devices/RemoteView.vue'
import { isLocal } from '@/services/ApiService'
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
