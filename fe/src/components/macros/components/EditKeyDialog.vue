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
  <div id="edit-key-dialog" class="dialog__content">
    <h4 class="text-slate-50 mb-4">Press a key</h4>
    <div class="flex justify-center" @click="$refs.newKeyInput.focus()">
      <MacroKey
        v-if="editable.key.keyObj"
        :key-obj="editable.key.keyObj"
        :direction="editable.key.direction"
      />

      <template v-if="typeof editable.newKey.keyObj === 'object'">
        <span class="px-4 flex items-center text-white"> >>> </span>
        <MacroKey :key-obj="editable.newKey.keyObj" :direction="editable.newKey.direction" />
      </template>
    </div>
    <form class="grid gap-4" submit.prevent>
      <input
        class="size-0 opacity-0"
        type="text"
        min="0"
        max="1"
        ref="newKeyInput"
        placeholder="New key"
        autofocus
        @keydown.prevent="handleNewKey($event)"
      />
      <div class="flex gap-2 justify-center">
        <ButtonComp
          variant="secondary"
          :class="editable.newKey.direction === 'down' ? 'selected' : ''"
          size="sm"
          @click.prevent="handleNewDirection('down')"
        >
          &darr; Down
        </ButtonComp>
        <ButtonComp
          variant="secondary"
          :class="editable.newKey.direction === 'up' ? 'selected' : ''"
          size="sm"
          @click.prevent="handleNewDirection('up')"
        >
          &uarr; Up
        </ButtonComp>
      </div>
      <div class="flex justify-end">
        <ButtonComp variant="primary" size="sm" @click.prevent="changeKey()">
          Change key
        </ButtonComp>
      </div>
    </form>
  </div>
</template>

<script setup>
import MacroKey from './MacroKey.vue'
import ButtonComp from '@/components/base/ButtonComp.vue'
import { useMacroRecorderStore } from '@/stores/macrorecorder'
import { filterKey } from '@/services/MacroRecordService'

import { reactive, ref, onMounted } from 'vue'

const editable = reactive({
  key: {},
  newKey: {},
})

const macroRecorder = useMacroRecorderStore()

const newKeyInput = ref(null)

onMounted(() => {
  editable.key = macroRecorder.getEditKey()
  editable.newKey.direction = editable.key.direction
})

const handleNewKey = (e) => {
  editable.newKey.e = e
  editable.newKey.keyObj = filterKey(e)
}

const handleNewDirection = (direction) => {
  editable.newKey.direction = direction
  editable.newKey.keyObj = filterKey(editable.key)
}

const changeKey = () => {
  macroRecorder.recordStep(
    editable.newKey.e,
    editable.newKey.direction,
    macroRecorder.state.editKey,
  )

  macroRecorder.state.editKey = false
}
</script>

<style scoped>
@reference "@/assets/main.css";

button.selected {
  @apply ring-2 ring-offset-1 ring-sky-500 bg-sky-500;
}
</style>
