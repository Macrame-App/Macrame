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
  <div id="devices-view" class="panel">
    <h1 class="panel__title">
      {{ isLocal() ? 'Remote devices' : 'Server' }}
    </h1>
    <div class="grid gap-8 pr-2 panel__content">
      <ServerView v-if="isLocal()" />
      <RemoteView v-else />
    </div>
  </div>
</template>

<script setup>
import ServerView from '@/components/devices/ServerView.vue'
import RemoteView from '@/components/devices/RemoteView.vue'
import { isLocal } from '@/services/ApiService'
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
