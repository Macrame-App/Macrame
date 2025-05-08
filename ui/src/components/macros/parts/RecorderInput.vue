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
  <div :class="`recorder-input__container ${macroRecorder.state.record && 'record'}`">
    <input
      v-if="macroRecorder.state.record"
      :class="`macro-recorder__input ${macroRecorder.state.record && 'record'}`"
      type="text"
      ref="macroInput"
      @keydown.prevent="macroRecorder.recordStep($event, 'down')"
      @keyup.prevent="macroRecorder.recordStep($event, 'up')"
    />
  </div>
</template>

<script setup>
import { useMacroRecorderStore } from '@/stores/macrorecorder'
import { ref, onUpdated } from 'vue'

const macroInput = ref(null)

const macroRecorder = useMacroRecorderStore()

onUpdated(() => {
  if (macroRecorder.state.record) {
    macroInput.value.focus()
    if (macroRecorder.delay.start !== 0) macroRecorder.restartDelay()
  }
})
</script>

<style scoped>
@reference "@/assets/main.css";

.recorder-input__container,
.macro-recorder__input {
  @apply absolute
    inset-0
    size-full
    opacity-0
    hidden;

  &.record {
    @apply block;
  }
}
</style>
