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
  <div id="delete-key-dialog" class="dialog__content">
    <h4 class="mb-4 text-slate-50">Delete key</h4>
    <div class="flex justify-center w-full mb-4">
      <MacroKey v-if="keyObj" :key-obj="keyObj" />
    </div>
    <p class="text-sm text-slate-300">Are you sure you want to delete this key?</p>
    <div class="flex justify-end gap-2 mt-6">
      <ButtonComp variant="danger" size="sm" @click="macroRecorder.deleteEditKey()">
        Delete key
      </ButtonComp>
    </div>
  </div>
</template>

<script setup>
import ButtonComp from '@/components/base/ButtonComp.vue'
import { useMacroRecorderStore } from '@/stores/macrorecorder'
import { onMounted, ref } from 'vue'
import MacroKey from './MacroKey.vue'
import { filterKey } from '@/services/MacroRecordService'

const macroRecorder = useMacroRecorderStore()

const keyObj = ref(null)

onMounted(() => {
  keyObj.value = filterKey(macroRecorder.getEditKey())
})
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
'
