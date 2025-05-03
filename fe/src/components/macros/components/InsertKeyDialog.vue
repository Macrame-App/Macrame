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
  <div id="insert-key-dialog" class="dialog__content w-96">
    <h4 class="text-slate-50 mb-4">Insert key {{ position }}</h4>
    <p v-if="inputFocus" class="text-center">[Press a key]</p>
    <input
      class="size-0 opacity-0"
      type="text"
      min="0"
      max="1"
      ref="insertKeyInput"
      placeholder="New key"
      @focusin="inputFocus = true"
      @focusout="inputFocus = false"
      @keydown.prevent="handleInsertKey($event)"
      autofocus
    />
    <div class="insert-output" :class="position == 'before' ? 'flex-row-reverse' : ''">
      <MacroKey v-if="keyObjs.selected" :key-obj="keyObjs.selected" />
      <hr class="spacer" />
      <DelaySpan :preset="true" :value="10" />
      <hr class="spacer" />
      <MacroKey
        v-if="keyObjs.insert"
        class="insert"
        :key-obj="keyObjs.insert"
        :direction="keyObjs.insertDirection"
        @click="insertKeyInput.focus()"
      />
      <MacroKey v-if="!keyObjs.insert" :empty="true" @click="insertKeyInput.focus()" />
      <template v-if="keyObjs.adjacentDelay">
        <hr class="spacer" />
        <DelaySpan :value="keyObjs.adjacentDelay.value" />
      </template>
      <template v-if="keyObjs.adjacent">
        <hr class="spacer" />
        <MacroKey :key-obj="keyObjs.adjacent" />
      </template>
    </div>
    <div class="insert-key__direction">
      <ButtonComp
        variant="secondary"
        :class="keyObjs.insertDirection === 'down' ? 'selected' : ''"
        size="sm"
        @click.prevent="keyObjs.insertDirection = 'down'"
      >
        &darr; Down
      </ButtonComp>
      <ButtonComp
        variant="secondary"
        :class="keyObjs.insertDirection === 'up' ? 'selected' : ''"
        size="sm"
        @click.prevent="keyObjs.insertDirection = 'up'"
      >
        &uarr; Up
      </ButtonComp>
    </div>
    <div class="flex justify-end">
      <ButtonComp variant="primary" size="sm" @click="insertKey()">Insert key</ButtonComp>
    </div>
  </div>
</template>

<script setup>
import MacroKey from './MacroKey.vue'
import DelaySpan from './DelaySpan.vue'
import ButtonComp from '@/components/base/ButtonComp.vue'

import { onMounted, reactive, ref } from 'vue'
import { useMacroRecorderStore } from '@/stores/macrorecorder'
import { filterKey } from '@/services/MacroRecordService'

const props = defineProps({
  position: String,
})

const macroRecorder = useMacroRecorderStore()

const keyObjs = reactive({
  selected: null,
  insert: null,
  insertEvent: null,
  insertDirection: 'down',
  adjacent: null,
  adjacentDelay: null,
  adjacentDelayIndex: null,
})

const insertKeyInput = ref(null)
const inputFocus = ref(false)

onMounted(() => {
  keyObjs.selected = filterKey(macroRecorder.getEditKey())

  const adjacentKey = macroRecorder.getAdjacentKey(props.position, true)
  if (adjacentKey) keyObjs.adjacent = filterKey(adjacentKey.key)
  if (adjacentKey.delay) {
    keyObjs.adjacentDelay = adjacentKey.delay
    keyObjs.adjacentDelayIndex = adjacentKey.delayIndex
  }
})

const handleInsertKey = (e) => {
  keyObjs.insert = filterKey(e)
  keyObjs.insertEvent = e
}

const insertKey = () => {
  macroRecorder.insertKey(keyObjs.insertEvent, keyObjs.insertDirection, keyObjs.adjacentDelayIndex)
}
</script>

<style scoped>
@reference "@/assets/main.css";

.insert-output {
  @apply flex 
  justify-center 
  items-center 
  w-full 
  mb-4;
}
.insert-key__direction {
  @apply flex
  justify-center
  gap-2
  mt-6;
}
button.selected {
  @apply bg-sky-500
  ring-2 
  ring-offset-1 
  ring-sky-500;
}
</style>
