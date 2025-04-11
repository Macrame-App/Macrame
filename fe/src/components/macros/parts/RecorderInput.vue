<template>
  <div :class="`recorder-input__container ${macroRecorder.state.record && 'record'}`">
    <input
      v-if="macroRecorder.state.record"
      :class="`macro-recorder__input ${macroRecorder.state.record && 'record'}`"
      type="text"
      ref="macroInput"
      @keydown.prevent="macroRecorder.recordStep($event, 'down')"
      @keyup.prevent="macroRecorder.recordStep($event, 'up')"
    />
  </div>
</template>

<script setup>
import { useMacroRecorderStore } from '@/stores/macrorecorder'
import { ref, onUpdated } from 'vue'

const macroInput = ref(null)

const macroRecorder = useMacroRecorderStore()

onUpdated(() => {
  if (macroRecorder.state.record) {
    macroInput.value.focus()
    if (macroRecorder.delay.start !== 0) macroRecorder.restartDelay()
  }
})
</script>

<style scoped>
@reference "@/assets/main.css";

.recorder-input__container,
.macro-recorder__input {
  @apply absolute
    inset-0
    size-full
    opacity-0
    hidden;

  &.record {
    @apply block;
  }
}
</style>
