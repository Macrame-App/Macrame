<template>
  <div
    :class="`alert alert__${variant} ${pageWide ? 'page-wide' : ''}`"
    @click="href ? router.push(href) : null"
  >
    <IconInfoCircle v-if="variant === 'info'" />
    <IconCheck v-if="variant === 'success'" />
    <IconExclamationCircle v-if="variant === 'warning'" />
    <IconAlertTriangle v-if="variant === 'error'" />
    <div class="alert__content">
      <slot />
    </div>
  </div>
</template>

<script setup>
import {
  IconAlertTriangle,
  IconCheck,
  IconExclamationCircle,
  IconInfoCircle,
} from '@tabler/icons-vue'
import { useRouter } from 'vue-router'

defineProps({
  variant: String, // info, success, warning, error
  pageWide: Boolean,
  href: String,
})

const router = useRouter()
</script>

<style scoped>
@reference "@/assets/main.css";

.alert {
  @apply grid
  grid-cols-[1rem_1fr]
  items-start
  gap-4
  p-4
  border
  border-white/10
  bg-white/10
  rounded-md
  backdrop-blur-md;

  &.alert__info {
    @apply text-sky-100 bg-sky-400/40;
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

  &.page-wide {
    @apply fixed
    bottom-0 left-0
    w-full;
  }

  &[href] {
    @apply cursor-pointer;
  }

  .alert__content {
    @apply grid gap-2;
  }
}
</style>
