<template>
  <div class="macro-recorder__header">
    <div class="w-full grid grid-cols-[auto_1fr_auto] gap-2">
      <h4 class="">Name:</h4>

      <input
        id="macro-name"
        type="text"
        @input.prevent="changeName($event.target.value)"
        placeholder="New macro"
      />
      <div :class="`recording__buttons ${!nameSet || macroRecorder.state.edit ? 'disabled' : ''}`">
        {{ macroRecorder.name }}
        <ButtonComp
          v-if="!macroRecorder.state.record"
          variant="primary"
          @click="macroRecorder.state.record = true"
        >
          <IconPlayerRecordFilled class="text-red-500" />Record
        </ButtonComp>
        <ButtonComp
          v-if="macroRecorder.state.record"
          variant="danger"
          @click="macroRecorder.state.record = false"
        >
          <IconPlayerStopFilled class="text-white" />Stop
        </ButtonComp>
      </div>
    </div>
    <div
      v-if="macroRecorder.steps.length > 0"
      :class="`edit__buttons ${macroRecorder.state.record ? 'disabled' : ''}`"
    >
      <div>
        <ButtonComp
          v-if="!macroRecorder.state.edit"
          variant="secondary"
          @click="macroRecorder.state.edit = true"
        >
          <IconPencil />Edit
        </ButtonComp>
        <ButtonComp
          v-if="macroRecorder.state.edit"
          variant="danger"
          @click="macroRecorder.resetEdit()"
        >
          <IconPlayerStopFilled />Stop
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
import { computed, onMounted, onUpdated, ref } from 'vue'

const macroRecorder = useMacroRecorderStore()

const nameSet = ref(false)

function changeName(name) {
  macroRecorder.changeName(name)
  nameSet.value = name.length > 0
}
</script>

<style scoped>
@reference "@/assets/main.css";

.macro-recorder__header {
  @apply grid 
  gap-4
  w-full;

  .edit__buttons {
    @apply flex
    justify-between
    gap-2
    w-full;
  }

  > div {
    @apply flex gap-2 items-end;
  }
}
</style>
