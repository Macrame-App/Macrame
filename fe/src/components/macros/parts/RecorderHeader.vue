<template>
  <div class="macro-recorder__header">
    <div :class="`recording__buttons ${macroRecorder.state.edit ? 'disabled' : ''}`">
      <ButtonComp
        v-if="!macroRecorder.state.record"
        variant="primary"
        @click="macroRecorder.state.record = true"
      >
        <IconPlayerRecordFilled class="text-red-500" />Start recording
      </ButtonComp>
      <ButtonComp
        v-if="macroRecorder.state.record"
        variant="danger"
        @click="macroRecorder.state.record = false"
      >
        <IconPlayerStopFilled class="text-white" />Stop recording
      </ButtonComp>
    </div>
    <div :class="`edit__buttons ${macroRecorder.state.record ? 'disabled' : ''}`">
      <div>
        <ButtonComp
          v-if="!macroRecorder.state.edit"
          variant="secondary"
          @click="macroRecorder.state.edit = true"
        >
          <IconPencil />Edit macro
        </ButtonComp>
        <ButtonComp
          v-if="macroRecorder.state.edit"
          variant="dark"
          @click="macroRecorder.resetEdit()"
        >
          <IconPlayerStopFilled />Stop editing
        </ButtonComp>
      </div>
      <FixedDelayMenu v-if="macroRecorder.state.edit" />
      <EditDialogs />
    </div>
  </div>
</template>

<script setup>
import { IconPencil, IconPlayerRecordFilled, IconPlayerStopFilled } from '@tabler/icons-vue'
import ButtonComp from '@/components/base/ButtonComp.vue'
import FixedDelayMenu from '../components/FixedDelayMenu.vue'

import { useMacroRecorderStore } from '@/stores/macrorecorder'
import EditDialogs from './EditDialogs.vue'

const macroRecorder = useMacroRecorderStore()
</script>

<style scoped>
@reference "@/assets/main.css";

.macro-recorder__header {
  @apply grid
  grid-cols-[auto_1fr]
  items-end
  gap-4;

  > div {
    @apply flex gap-4 items-end;

    &.disabled {
      @apply opacity-50 pointer-events-none cursor-not-allowed;
    }
  }
}
</style>
