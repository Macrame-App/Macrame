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
  <div id="dashboard" class="panel">
    <div class="panel__title">
      <h1>Dashboard</h1>
      <div>
        <em v-if="isLocal()">This is the server dashboard.</em>
        <em v-else>This is the remote dashboard.</em>
      </div>
    </div>

    <div class="panel__content !h-fit !gap-y-16">
      <ServerView v-if="isLocal()" />
      <RemoteView v-else />
      <div class="grid gap-2 text-slate-300">
        <h3>About Macrame</h3>
        <p>Macrame is an open-source application designed to turn any device into a customizable button panel. Whether you're optimizing your workflow or enhancing your gaming experience, Macrame makes it simple to create and link macros to your button panels.</p>
        <p>For more information, including details on licensing, visit <a href="https://macrame.github.io" target="_blank">https://macrame.github.io</a></p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { isLocal } from '@/services/ApiService'
import ServerView from '@/components/dashboard/ServerView.vue'
import RemoteView from '@/components/dashboard/RemoteView.vue'
</script>

<style>
@reference "@/assets/main.css";

.dashboard-block {
  @apply md:!row-start-1
  grid 
  justify-items-center
  gap-4;

  &#devices .icon__container,
  &#server .icon__container {
    @apply bg-sky-300/30
    text-sky-400
    border-sky-300/60;
  }

  &#macros .icon__container {
    @apply bg-amber-300/30
    text-amber-400
    border-amber-300/60;
  }

  &#panels .icon__container {
    @apply bg-rose-300/30
    text-rose-400
    border-rose-300/60;
  }

  .icon__container {
    @apply flex
    justify-center
    items-center
    size-16
    aspect-square
    rounded-full      
    border;

    svg {
      @apply size-8;
    }
  }

  p {
    @apply opacity-50
    w-42
    text-center;
  }
}
</style>
