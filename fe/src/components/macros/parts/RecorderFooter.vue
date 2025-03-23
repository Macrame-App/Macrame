<template>
  <div class="macro-recorder__footer">
    <ButtonComp
      v-if="macroRecorder.steps.length > 0"
      variant="danger"
      size="sm"
      @click="macroRecorder.reset()"
    >
      <IconRestore /> Reset
    </ButtonComp>

    <DialogComp ref="errorDialog">
      <template #content>
        <ValidationErrorDialog />
      </template>
    </DialogComp>

    <ButtonComp
      v-if="macroRecorder.steps.length > 0"
      :disabled="macroRecorder.state.record || macroRecorder.state.edit"
      variant="success"
      size="sm"
      @click="toggleSave()"
    >
      <IconDeviceFloppy />
      Save
    </ButtonComp>
  </div>
</template>

<script setup>
import ButtonComp from '@/components/base/ButtonComp.vue'
import { IconDeviceFloppy, IconRestore } from '@tabler/icons-vue'

import { useMacroRecorderStore } from '@/stores/macrorecorder'
import DialogComp from '@/components/base/DialogComp.vue'
import ValidationErrorDialog from '../components/ValidationErrorDialog.vue'
import { onMounted, ref } from 'vue'

const macroRecorder = useMacroRecorderStore()

const errorDialog = ref()

onMounted(() => {
  macroRecorder.$subscribe((mutation) => {
    if (mutation.events && mutation.events.key == 'validationErrors') {
      errorDialog.value.toggleDialog(mutation.events.newValue !== false)
    }
  })
})

const toggleSave = () => {
  if (!macroRecorder.save()) errorDialog.value.toggleDialog(true)
}
</script>

<style scoped>
@reference "@/assets/main.css";

.macro-recorder__footer {
  @apply flex
  justify-between
  gap-2;
}
</style>
