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
  <div id="edit-delay-dialog" class="dialog__content">
    <h4 class="mb-4 text-slate-50">Edit delay</h4>
    <div v-if="editable.delay.value" class="flex justify-center">
      <DelaySpan class="!text-lg" :value="editable.delay.value" />
    </div>
    <form class="grid gap-4 mt-6" submit.prevent>
      <div v-if="editable.newDelay.value">
        <input
          type="number"
          min="0"
          max="3600000"
          step="10"
          v-model="editable.newDelay.value"
          autofocus
        />
        <span>ms</span>
      </div>
      <div class="flex justify-end">
        <ButtonComp variant="primary" size="sm" @click.prevent="changeDelay()">
          Change delay
        </ButtonComp>
      </div>
    </form>
  </div>
</template>

<script setup>
import ButtonComp from '@/components/base/ButtonComp.vue'
import { reactive, onMounted } from 'vue'
import { useMacroRecorderStore } from '@/stores/macrorecorder'
import DelaySpan from './DelaySpan.vue'

const macroRecorder = useMacroRecorderStore()

const editable = reactive({
  delay: {},
  newDelay: { value: 0 },
})

onMounted(() => {
  editable.delay = macroRecorder.getEditDelay()
  editable.newDelay.value = editable.delay.value
})

const changeDelay = () => {
  if (!editable.newDelay.value) return

  macroRecorder.recordStep(editable.newDelay.value, false, macroRecorder.state.editDelay)
  macroRecorder.state.editDelay = false
}
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
