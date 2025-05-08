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
  <kbd :class="`${active ? 'active' : ''} ${empty ? 'empty' : ''}`">
    <template v-if="keyObj">
      <sup v-if="keyObj.loc">
        {{ keyObj.loc }}
      </sup>
      <span :innerHTML="keyObj.str" />
      <span class="dir">{{ dir.value === 'down' ? '&darr;' : '&uarr;' }}</span>
    </template>

    <template v-else-if="empty">
      <span>[ ]</span>
    </template>
  </kbd>
</template>

<script setup>
import { onMounted, onUpdated, reactive } from 'vue'

const props = defineProps({
  keyObj: Object,
  direction: String,
  active: Boolean,
  empty: Boolean,
})

const dir = reactive({
  value: false,
})

onMounted(() => {
  if (props.empty) return
  setDirection()
})

onUpdated(() => {
  setDirection()
})

const setDirection = () => {
  if (props.direction) dir.value = props.direction
  else dir.value = props.keyObj.direction
}
</script>

<style>
@reference "@/assets/main.css";

kbd {
  @apply flex 
  items-center
  gap-2
  pl-4 pr-2 py-1 
  h-9 
  bg-slate-700
  font-mono
  font-bold
  text-lg
  text-white
  whitespace-nowrap
  uppercase
  rounded-md 
  border
  border-slate-500
  transition-all
  shadow-slate-500;
  box-shadow: 0 0.2rem 0 0.2rem var(--tw-shadow-color);

  &:has(sup) {
    @apply pl-2;
  }

  sup {
    @apply text-slate-200 text-xs font-light mt-1;
  }

  span.dir {
    @apply text-slate-200 pl-1;
  }

  &.empty {
    @apply pl-3 pr-3 
    bg-sky-400/50
    border-sky-300
    shadow-sky-600
    tracking-widest
    cursor-pointer;
  }
  &.insert {
    @apply bg-yellow-500/50 
    border-yellow-300 
    shadow-yellow-600 
    cursor-pointer;
  }
}

:has(kdb):not(.edit) kbd {
  @apply pointer-events-none cursor-default;
}

.edit kbd {
  @apply cursor-pointer pointer-events-auto;

  &:hover,
  &.active {
    @apply bg-sky-900 border-sky-400 shadow-sky-700;
  }
}
</style>
