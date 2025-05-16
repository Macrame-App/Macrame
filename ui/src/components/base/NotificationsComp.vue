<template>
  <div class="mcrm-notifications">
    <ToastComp
      v-for="(notification, id) in notificationList"
      :key="id"
      :id="id"
      :title="notification.title"
      :message="notification.message"
      :variant="notification.variant"
      :time="notification.time"
      :closable="notification.closable"
      :notification="id"
    />
  </div>
</template>

<script setup>
import { useNoticationStore } from '@/stores/notifications'
import { computed, onMounted } from 'vue'
import ToastComp from './ToastComp.vue'

const notifications = useNoticationStore()

const notificationList = computed(() => notifications.list)

// onMounted(() => {
//   notifications.$subscribe((mutation, state) => {
//     console.log(mutation, state)
//   })
// })
</script>

<style scoped>
@reference "@/assets/main.css";

.mcrm-notifications {
  @apply grid
  gap-4
  fixed
  bottom-0
  right-0
  p-4
  z-50
  w-[30ch];
}
</style>
