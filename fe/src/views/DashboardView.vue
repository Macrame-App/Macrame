<template>
  <div id="dashboard" class="panel">
    <div class="panel__title">
      <h1>Dashboard</h1>
      <div>
        <em v-if="isLocal()">This is the server dashboard.</em>
        <em v-else>This is the remote dashboard.</em>
      </div>
    </div>

    <div class="panel__content">
      <ServerView v-if="isLocal()" />
      <RemoteView v-else />
    </div>
  </div>
</template>

<script setup>
import { isLocal } from '@/services/ApiService'
import ServerView from '@/components/dashboard/ServerView.vue'
import RemoteView from '@/components/dashboard/RemoteView.vue'
</script>

<style>
@reference "@/assets/main.css";

.dashboard-block {
  @apply md:!row-start-1
  grid 
  justify-items-center
  gap-4;

  &#devices .icon__container,
  &#server .icon__container {
    @apply bg-sky-300/30
    text-sky-400
    border-sky-300/60;
  }

  &#macros .icon__container {
    @apply bg-amber-300/30
    text-amber-400
    border-amber-300/60;
  }

  &#panels .icon__container {
    @apply bg-rose-300/30
    text-rose-400
    border-rose-300/60;
  }

  .icon__container {
    @apply flex
    justify-center
    items-center
    size-16
    aspect-square
    rounded-full      
    border;

    svg {
      @apply size-8;
    }
  }

  p {
    @apply opacity-50
    w-42
    text-center;
  }
}
</style>
