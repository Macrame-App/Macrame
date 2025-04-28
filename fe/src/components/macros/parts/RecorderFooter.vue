<template>
  <div class="macro-recorder__footer">
    <ButtonComp
      v-if="macroRecorder.steps.length > 0"
      variant="danger"
      @click="macroRecorder.reset()"
    >
      <IconRestore /> Reset
    </ButtonComp>

    <DialogComp ref="errorDialog">
      <template #content>
        <ValidationErrorDialog />
      </template>
    </DialogComp>
    <DialogComp ref="overwriteDialog">
      <template #content>
        <div class="grid gap-2">
          <h4 class="pr-4">Are you sure you want to overwrite:</h4>
          <h3 class="mb-2 text-center text-sky-500">{{ macroRecorder.macroName }}</h3>
          <div class="flex justify-between">
            <ButtonComp size="sm" variant="subtle" @click="overwriteDialog.toggleDialog(false)"
              >No</ButtonComp
            >
            <ButtonComp size="sm" variant="primary" @click="saveMacro()">Yes</ButtonComp>
          </div>
        </div>
      </template>
    </DialogComp>

    <ButtonComp
      v-if="macroRecorder.steps.length > 0"
      :disabled="macroRecorder.state.record || macroRecorder.state.edit"
      variant="success"
      @click="startCheck()"
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
const overwriteDialog = ref()

onMounted(() => {
  macroRecorder.$subscribe((mutation) => {
    if (mutation.events && mutation.events.key == 'validationErrors') {
      errorDialog.value.toggleDialog(mutation.events.newValue !== false)
    }
  })
})

const startCheck = async () => {
  const checkResp = await macroRecorder.checkMacro()

  if (checkResp) overwriteDialog.value.toggleDialog(true)
  else saveMacro()
}

const saveMacro = async () => {
  overwriteDialog.value.toggleDialog(false)

  const saveResp = await macroRecorder.saveMacro()

  if (!saveResp) errorDialog.value.toggleDialog(true)
  else window.location.reload()
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
