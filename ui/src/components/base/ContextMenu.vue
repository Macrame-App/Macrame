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
  <div class="context-menu">
    <div class="context-menu__trigger" @click="toggle">
      <slot name="trigger" />
    </div>
    <div :class="`context-menu__content ${menuOpen ? 'open' : ''}`">
      <slot name="content" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUpdated } from 'vue'

defineExpose({ toggle })

const props = defineProps({
  open: Boolean,
})

const menuOpen = ref(false)

onMounted(() => {
  menuOpen.value = props.open
})

function toggle() {
  menuOpen.value = !menuOpen.value
}
</script>

<style>
@reference "@/assets/main.css";

.context-menu {
  @apply relative;

  .context-menu__content {
    @apply absolute
    top-full
    -translate-y-full
    opacity-0
    pointer-events-none
    mt-2
    min-w-full
    grid
    border
    border-white/50
    bg-slate-100/60
    backdrop-blur-3xl
    text-slate-800
    rounded-md
    z-50
    transition-all;

    &.open {
      @apply translate-y-0
      opacity-100
      pointer-events-auto;
    }
  }
}

.context-menu ul {
  @apply text-slate-800
  divide-y
  divide-slate-300;

  li {
    @apply flex
    gap-2
    items-center
    p-2
    hover:bg-black/10
    cursor-pointer;

    svg {
      @apply size-5;
    }
  }
}
</style>
