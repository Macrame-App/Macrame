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
  <div v-if="loading" class="loading-component">
    <span v-if="text">
      {{ text }}
    </span>
    <IconLoader3 class="duration-1000 animate-spin" />
  </div>
</template>

<script setup>
import { IconLoader3 } from '@tabler/icons-vue'

defineProps({
  loading: Boolean,
  text: String,
})
</script>

<style>
@reference "@/assets/main.css";

&:has(.loading-component) {
  @apply relative;
}

.loading-component {
  @apply absolute
  inset-0
  size-full
  flex gap-2
  flex-col
  justify-center
  items-center
  text-sm
  bg-black/50
  backdrop-blur-md;
}
</style>
