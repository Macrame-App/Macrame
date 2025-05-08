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
  <ContextMenu ref="ctxtMenu">
    <template #trigger>
      <ButtonComp variant="secondary" size="sm"> <IconTimeDuration15 />Fixed delay </ButtonComp>
    </template>
    <template #content>
      <ul>
        <li @click="changeDelay(0)">0ms</li>
        <li @click="changeDelay(15)">15ms</li>
        <li @click="changeDelay(50)">50ms</li>
        <li @click="changeDelay(100)">100ms</li>
        <li>
          <DialogComp>
            <template #trigger>
              <span>Custom delay</span>
            </template>

            <template #content>
              <h4 class="text-slate-50 mb-4">Custom delay</h4>
              <form
                class="grid gap-4 w-44"
                @submit.prevent="changeDelay(parseInt($refs.customDelayInput.value))"
              >
                <div>
                  <input
                    type="number"
                    step="10"
                    min="0"
                    max="3600000"
                    ref="customDelayInput"
                    placeholder="100"
                  />
                  <span>ms</span>
                </div>
                <div class="flex justify-end">
                  <ButtonComp variant="primary" size="sm">Set custom delay</ButtonComp>
                </div>
              </form>
            </template>
          </DialogComp>
        </li>
      </ul>
    </template>
  </ContextMenu>
</template>

<script setup>
import ContextMenu from '@/components/base/ContextMenu.vue'
import { IconTimeDuration15 } from '@tabler/icons-vue'
import ButtonComp from '@/components/base/ButtonComp.vue'
import DialogComp from '@/components/base/DialogComp.vue'

import { useMacroRecorderStore } from '@/stores/macrorecorder'

import { ref } from 'vue'

const macroRecorder = useMacroRecorderStore()

const ctxtMenu = ref()

function changeDelay(num) {
  macroRecorder.changeDelay(num)
  ctxtMenu.value.toggle()
}
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
