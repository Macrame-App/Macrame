<template>
  <nav id="main-menu">
    <button id="menu-toggle" :class="menuOpen ? 'open' : ''" @click="menuOpen = !menuOpen">
      <IconMenu2 />
      <IconX />
    </button>
    <ul :class="menuOpen ? 'open' : ''">
      <li>
        <RouterLink @click="menuOpen = false" to="/"> <IconHome />Dashboard </RouterLink>
      </li>
      <li>
        <RouterLink @click="menuOpen = false" to="/panels"> <IconLayoutGrid />Panels </RouterLink>
      </li>
      <li>
        <RouterLink @click="menuOpen = false" to="/macros"> <IconKeyboard />Macros </RouterLink>
      </li>
      <li>
        <RouterLink @click="menuOpen = false" to="/settings"> <IconSettings />Settings </RouterLink>
      </li>
    </ul>
  </nav>
</template>

<script setup>
import { RouterLink } from 'vue-router'
import {
  IconHome,
  IconKeyboard,
  IconLayoutGrid,
  IconMenu2,
  IconSettings,
  IconX,
} from '@tabler/icons-vue'
import { ref } from 'vue'

const menuOpen = ref(false)
</script>

<style scoped>
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
    cursor-pointer
    transition-colors
    backdrop-blur-md;

    svg {
      @apply absolute
      inset-1/2
      -translate-1/2
      transition-opacity
      duration-400
      ease-in-out;
    }

    svg:last-of-type {
      @apply opacity-0;
    }

    &.open {
      svg:first-of-type {
        @apply opacity-0;
      }

      svg:last-of-type {
        @apply opacity-100;
      }
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
