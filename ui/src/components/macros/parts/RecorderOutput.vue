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
  <div
    :class="`macro-recorder__output ${macroRecorder.state.record && 'record'} ${macroRecorder.state.edit && 'edit'}`"
  >
    <template v-for="(step, key) in macroRecorder.steps">
      <!-- Key element -->
      <template v-if="step.type === 'key'">
        <MacroKey
          :key="key"
          :key-obj="step.keyObj"
          :direction="step.direction"
          :active="macroRecorder.state.editKey === key"
          @click="macroRecorder.state.edit ? macroRecorder.toggleEdit('key', key) : false"
        />
      </template>

      <!-- Delay element -->
      <template v-else-if="step.type === 'delay'">
        <DelaySpan
          :key="key"
          :value="step.value"
          :active="macroRecorder.state.editDelay === key"
          @click="macroRecorder.toggleEdit('delay', key)"
        />
      </template>

      <!-- Spacer element -->
      <hr class="spacer" />
    </template>
  </div>
</template>

<script setup>
import { useMacroRecorderStore } from '@/stores/macrorecorder'
import MacroKey from '../components/MacroKey.vue'
import DelaySpan from '../components/DelaySpan.vue'

const macroRecorder = useMacroRecorderStore()
</script>

<style scoped>
@reference "@/assets/main.css";

.macro-recorder__output {
  @apply flex
  flex-wrap
  items-center
  gap-y-4
  p-4
  absolute
  top-0 left-0
  h-fit;
}

hr.spacer:last-of-type {
  @apply hidden;
}
</style>
