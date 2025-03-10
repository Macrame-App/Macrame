<template>
  <ContextMenu>
    <template #trigger>
      <ButtonComp variant="secondary" size="sm"> <IconAlarmFilled />Fixed delay </ButtonComp>
    </template>
    <template #content>
      <ul>
        <li @click="macroRecorder.changeDelay(0)">0ms</li>
        <li @click="macroRecorder.changeDelay(15)">15ms</li>
        <li @click="macroRecorder.changeDelay(50)">50ms</li>
        <li @click="macroRecorder.changeDelay(100)">100ms</li>
        <li>
          <DialogComp>
            <template #trigger>
              <span>Custom delay</span>
            </template>

            <template #content>
              <h4 class="text-slate-50 mb-4">Custom delay</h4>
              <form
                class="grid gap-4 w-44"
                @submit.prevent="macroRecorder.changeDelay(parseInt($refs.customDelayInput.value))"
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
import { IconAlarmFilled } from '@tabler/icons-vue'
import ButtonComp from '@/components/base/ButtonComp.vue'
import DialogComp from '@/components/base/DialogComp.vue'

import { useMacroRecorderStore } from '@/stores/macrorecorder'

import { ref } from 'vue'

const macroRecorder = useMacroRecorderStore()

const delayMenu = ref(false)
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
