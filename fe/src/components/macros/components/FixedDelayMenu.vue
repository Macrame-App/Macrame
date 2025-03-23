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
