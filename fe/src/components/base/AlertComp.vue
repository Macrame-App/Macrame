<template>
  <div :class="`alert alert__${type}`">
    <IconInfoCircle v-if="type === 'info'" />
    <IconCheck v-if="type === 'success'" />
    <IconExclamationCircle v-if="type === 'warning'" />
    <IconAlertTriangle v-if="type === 'error'" />
    <slot />
  </div>
</template>

<script setup>
import {
  IconAlertTriangle,
  IconCheck,
  IconExclamationCircle,
  IconInfoCircle,
} from '@tabler/icons-vue'

defineProps({
  type: String, // info, success, warning, error
})
</script>

<style scoped>
@reference "@/assets/main.css";

.alert {
  @apply flex
  items-center
  gap-4
  p-4
  border
  border-white/10
  bg-white/10
  rounded-md
  backdrop-blur-md;

  &.alert__info {
    @apply text-sky-400 bg-sky-400/10;
  }

  &.alert__success {
    @apply text-lime-400 bg-lime-400/10;
  }

  &.alert__warning {
    @apply text-amber-400 bg-amber-400/10;
  }

  &.alert__error {
    @apply text-rose-400 bg-rose-400/10;
  }
}
</style>
