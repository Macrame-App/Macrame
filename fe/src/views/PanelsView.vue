<template>
  <div id="macros" class="panel">
    <h1 class="panel__title">
      Panels <span class="text-sm">{{ isLocal() ? 'remote' : 'servers' }}</span>
    </h1>
    <div class="panel__content !p-0">
      <div class="macro-panel__content">
        <PanelsOverview />
      </div>
    </div>
  </div>
</template>

<script setup>
import PanelsOverview from '@/components/panels/PanelsOverview.vue'
import { isLocal } from '@/services/ApiService'
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
