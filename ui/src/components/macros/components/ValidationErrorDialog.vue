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
  <div id="validation-error__dialog" class="dialog__content">
    <h4 class="mb-4 text-slate-50">There's an error in your macro</h4>

    <div class="grid gap-4" v-if="(errors && errors.up.length > 0) || errors.down.length > 0">
      <div v-if="errors.down.length > 0">
        <p>
          The following keys have been <strong>pressed</strong> down, but
          <strong>not released</strong>.
        </p>
        <ul>
          <li v-for="key in errors.down" :key="key">{{ key.toUpperCase() }}</li>
        </ul>
      </div>

      <div v-if="errors.up.length > 0">
        <p>
          The following keys have been <strong>released</strong>, but
          <strong>not pressed</strong> down.
        </p>
        <ul>
          <li v-for="key in errors.up" :key="key">{{ key.toUpperCase() }}</li>
        </ul>
      </div>
    </div>
    <div class="flex justify-end mt-4">
      <ButtonComp size="sm" variant="danger" @click="macroRecorder.state.validationErrors = false">
        Close
      </ButtonComp>
    </div>
  </div>
</template>

<script setup>
import ButtonComp from '@/components/base/ButtonComp.vue'
import { useMacroRecorderStore } from '@/stores/macrorecorder'
import { onMounted, reactive } from 'vue'

const macroRecorder = useMacroRecorderStore()

const errors = reactive({
  up: [],
  down: [],
})

onMounted(() => {
  macroRecorder.$subscribe((mutation) => {
    if (mutation.events && mutation.events.key == 'validationErrors') {
      errors.up = mutation.events.newValue !== false ? macroRecorder.state.validationErrors.up : []
      errors.down =
        mutation.events.newValue !== false ? macroRecorder.state.validationErrors.down : []
    }
  })
})
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
