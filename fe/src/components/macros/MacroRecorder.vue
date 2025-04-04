<template>
  <div class="macro-recorder mcrm-block block__light">
    <div class="recorder-interface">
      <!-- Recorder buttons -->
      <RecorderHeader />

      <!-- Recorder interface container -->
      <div
        :class="`recorder-interface__container ${macroRecorder.state.record && 'record'} ${macroRecorder.state.edit && 'edit'}`"
      >
        <!-- Shows the macro steps as kbd elements with delay and spacers-->
        <RecorderOutput />
        <!-- Input for recording macro steps -->
        <RecorderInput />
      </div>

      <RecorderFooter />
    </div>
  </div>
</template>

<script setup>
import RecorderOutput from './parts/RecorderOutput.vue'
import RecorderInput from './parts/RecorderInput.vue'

import { useMacroRecorderStore } from '@/stores/macrorecorder'
import RecorderHeader from './parts/RecorderHeader.vue'
import RecorderFooter from './parts/RecorderFooter.vue'

const macroRecorder = useMacroRecorderStore()
</script>

<style>
@reference "@/assets/main.css";

.macro-recorder {
  @apply h-full;
}

.recorder-interface {
  @apply grid
  grid-rows-[auto_1fr_auto]
  gap-4
  h-full
  transition-[grid-template-rows];
}

.recorder-interface__container {
  @apply relative
  w-full
  rounded-lg
  bg-slate-950/50
  border
  border-slate-600
  overflow-auto
  transition-colors;

  &.record {
    @apply border-rose-300  bg-rose-400/10;
  }

  &.edit {
    @apply border-sky-300 bg-sky-900/10;
  }
}

#macro-name {
  @apply w-full
  bg-transparent
  py-0
  outline-0
  border-transparent
  border-b-slate-300
  focus:border-transparent
  focus:border-b-sky-400
  focus:bg-sky-400/10
  transition-colors
  text-lg
  rounded-none;
}

.disabled {
  @apply opacity-50 pointer-events-none cursor-not-allowed;
}
</style>
