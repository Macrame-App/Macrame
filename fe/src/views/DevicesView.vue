<template>
  <div id="devices-view" class="panel">
    <h1 class="panel__title">
      {{ isLocal() ? 'Remote devices' : 'Server' }}
    </h1>
    <div class="grid gap-8 pr-2 panel__content">
      <ServerView />
      <RemoteView />
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
