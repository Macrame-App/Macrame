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
  <div id="macros" class="panel">
    <h1 class="panel__title">Macros</h1>
    <div class="panel__content !p-0 !overflow-hidden">
      <div class="macro-panel__content">
        <MacroOverview />
        <MacroRecorder />
      </div>
    </div>
  </div>
</template>

<script setup>
import MacroOverview from '@/components/macros/MacroOverview.vue'
import MacroRecorder from '../components/macros/MacroRecorder.vue'
</script>

<style scoped>
@reference "@/assets/main.css";

.macro-panel__content {
  @apply grid
  grid-cols-[25ch_1fr]
  gap-6
  pt-2;
}
</style>
