<!--
Macrame is a program that enables the user to create keyboard macros and button panels. 
The macros are saved as simple JSON files and can be linked to the button panels. The panels can 
be created with HTML and CSS.

Copyright (C) 2025 Jesse Malotaux

This program is free software: you can redistribute it and/or modify 
it under the terms of the GNU General Public License as published by 
the Free Software Foundation, either version 3 of the License, or 
(at your option) any later version.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
GNU General Public License for more details.

You should have received a copy of the GNU General Public License 
along with this program. If not, see <https://www.gnu.org/licenses/>.
-->

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
