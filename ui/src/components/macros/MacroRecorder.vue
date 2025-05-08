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
  <div class="macro-recorder mcrm-block block__light">
    <div class="recorder-interface">
      <!-- Recorder buttons -->
      <RecorderHeader />

      <!-- Recorder interface container -->
      <div
        :class="`recorder-interface__container ${macroRecorder.state.record && 'record'} ${macroRecorder.state.edit && 'edit'}`"
      >
        <!-- Shows the macro steps as kbd elements with delay and spacers-->
        <RecorderOutput />
        <!-- Input for recording macro steps -->
        <RecorderInput />
      </div>

      <RecorderFooter />
    </div>
  </div>
</template>

<script setup>
import RecorderOutput from './parts/RecorderOutput.vue'
import RecorderInput from './parts/RecorderInput.vue'

import { useMacroRecorderStore } from '@/stores/macrorecorder'
import RecorderHeader from './parts/RecorderHeader.vue'
import RecorderFooter from './parts/RecorderFooter.vue'

const macroRecorder = useMacroRecorderStore()
</script>

<style>
@reference "@/assets/main.css";

.macro-recorder {
  @apply h-full;
}

.recorder-interface {
  @apply grid
  grid-rows-[auto_1fr_auto]
  gap-4
  h-full
  transition-[grid-template-rows];
}

.recorder-interface__container {
  @apply relative
  w-full
  rounded-lg
  bg-slate-950/50
  border
  border-slate-600
  overflow-auto
  transition-colors;

  &.record {
    @apply border-rose-300  bg-rose-400/10;
  }

  &.edit {
    @apply border-sky-300 bg-sky-900/10;
  }
}

#macro-name {
  @apply w-full
  bg-transparent
  py-0
  outline-0
  border-transparent
  border-b-slate-300
  focus:border-transparent
  focus:border-b-sky-400
  focus:bg-sky-400/10
  transition-colors
  text-lg
  rounded-none;
}

.disabled {
  @apply opacity-50 pointer-events-none cursor-not-allowed;
}
</style>
