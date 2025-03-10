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
</style>
