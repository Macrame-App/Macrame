<template>
  <div class="macro-recorder">
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
</template>

<script setup>
// TODO:
// X refactor filtering the keys
// X add keyup functionality
// X add delay between steps
// X restyle keys and delay elements
// X add lines between steps
// X record macro to object
// X refactor macro output based on object?
// X Make sure keydown is not spamming steps
// X Make record button work as a toggle
// X Make edit button work
// X Make fixed/custom delay work
// X Refactor into multiple components and state store
// X Make edit key function
// X Make edit delay function
// X Make delete key function
// X Make sure delay is paused when not recording.
// X Refactor macro recorder parts.
// X X Layout parts should be parts, smaller parts should be components.
// X Make reset function
// - Make insert button, menu and function

import RecorderOutput from './parts/RecorderOutput.vue'
import RecorderInput from './parts/RecorderInput.vue'

import { useMacroRecorderStore } from '@/stores/macrorecorder'
import RecorderHeader from './parts/RecorderHeader.vue'
import RecorderFooter from './parts/RecorderFooter.vue'

const macroRecorder = useMacroRecorderStore()
</script>

<style>
@reference "@/assets/main.css";

.recorder-interface__container {
  @apply relative
  w-full
  h-96
  my-4
  rounded-lg
  bg-slate-900/50
  border-2
  border-white/10
  overflow-auto
  transition-colors;

  &.record {
    @apply border-rose-300  bg-rose-950/50;
  }

  &.edit {
    @apply border-sky-300 bg-sky-900/50;
  }
}
</style>
