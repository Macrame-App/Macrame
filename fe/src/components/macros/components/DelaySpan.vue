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
  <span :class="`delay ${active ? 'active' : ''} ${preset ? 'preset' : ''}`">
    <template v-if="value < 10000"> {{ value }} <i>ms</i> </template>
    <template v-else> >10 <i>s</i> </template>
  </span>
</template>

<script setup>
import { IconTimeDuration10 } from '@tabler/icons-vue'

defineProps({
  value: Number,
  active: Boolean,
  preset: Boolean,
})
</script>

<style scoped>
@reference "@/assets/main.css";

span.delay {
  @apply flex 
  items-center
  px-2 py-1
  bg-slate-500
  border
  border-slate-400
  text-slate-950
  font-sans
  font-semibold
  rounded-sm
  text-sm
  cursor-default;

  &.preset {
    @apply text-amber-400
    border-amber-300/80
    bg-amber-100/60;
  }

  i {
    @apply pl-1
    font-normal 
    not-italic 
    opacity-80;
  }
}

.edit span.delay {
  @apply cursor-pointer;

  &:hover,
  &.active {
    @apply bg-lime-700 border-lime-500 text-lime-200;
  }
}
</style>
