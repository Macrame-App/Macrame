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
  <nav id="main-menu">
    <button id="menu-toggle" :class="menuOpen ? 'open' : ''" @click="menuOpen = !menuOpen">
      <img
        class="p-1 logo"
        :class="{ 'opacity-0': menuOpen }"
        src="@/assets/img/Macrame-Logo-gradient.svg"
        aria-hidden="true"
      />
      <IconX :class="{ 'opacity-0': !menuOpen }" />
    </button>
    <ul :class="menuOpen ? 'open' : ''">
      <li>
        <RouterLink @click="menuOpen = false" to="/"> <IconHome />Dashboard </RouterLink>
      </li>
      <li>
        <RouterLink @click="menuOpen = false" to="/panels"> <IconLayoutGrid />Panels </RouterLink>
      </li>
      <li v-if="isLocal()">
        <RouterLink @click="menuOpen = false" to="/macros"> <IconKeyboard />Macros </RouterLink>
      </li>
      <li>
        <RouterLink @click="menuOpen = false" to="/devices">
          <IconDevices />{{ isLocal() ? 'Devices' : 'Server' }}
        </RouterLink>
      </li>
      <!-- <li>
        <RouterLink @click="menuOpen = false" to="/settings">
          <IconSettings />Settings
        </RouterLink>
      </li> -->
    </ul>
  </nav>
</template>

<script setup>
import { RouterLink } from 'vue-router'
import {
  IconDevices,
  IconHome,
  IconKeyboard,
  IconLayoutGrid,
  IconSettings,
  IconX,
} from '@tabler/icons-vue'
import { ref } from 'vue'
import { isLocal } from '@/services/ApiService'

const menuOpen = ref(false)
</script>

<style>
@reference "@/assets/main.css";
nav {
  @apply relative flex z-50;

  button {
    @apply absolute
    top-4 left-4
    size-12
    rounded-full
    aspect-square
    bg-white/20 hover:bg-white/40
    border-0
    cursor-pointer
    transition-colors
    backdrop-blur-md;

    .logo,
    svg {
      @apply absolute
      inset-1/2
      -translate-1/2
      transition-opacity
      duration-400
      ease-in-out;
    }

    .logo {
      @apply w-full;
    }
  }

  ul {
    @apply absolute
    top-20 left-0
    -translate-x-full
    grid
    list-none
    rounded-xl
    overflow-hidden
    bg-white/10
    backdrop-blur-md
    divide-y
    divide-slate-600
    transition-transform
    duration-300
    ease-in-out;

    &.open {
      @apply left-4 translate-x-0;
    }

    li {
      a {
        @apply flex
        items-center
        gap-2
        px-4 py-2
        text-white
        no-underline
        border-transparent
        transition-colors;

        svg {
          @apply text-white/40 transition-colors;
        }

        &:hover {
          @apply bg-white/20;

          svg {
            @apply text-white;
          }
        }

        &.router-link-active {
          @apply text-sky-300
          bg-sky-200/20;
        }
      }
    }
  }
}
</style>
