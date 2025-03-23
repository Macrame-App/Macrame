<template>
  <div class="macro-edit__dialogs" v-if="macroRecorder.state.edit !== false">
    <div
      class="flex gap-2"
      v-if="macroRecorder.state.editKey !== false && typeof macroRecorder.getEditKey() === 'object'"
    >
      <ContextMenu ref="ctxtMenu">
        <template #trigger>
          <ButtonComp variant="dark" size="sm"> <IconPlus /> Insert </ButtonComp>
        </template>
        <template #content>
          <ul>
            <li @click="insert.position = 'before'"><IconArrowLeftCircle /> Before</li>
            <li @click="insert.position = 'after'"><IconArrowRightCircle /> After</li>
          </ul>
        </template>
      </ContextMenu>

      <DialogComp
        v-if="insert.position !== null"
        :open="insert.position !== null"
        @on-open="onOpenDialog"
        @on-close="onCloseDialog"
      >
        <template #content>
          <InsertKeyDialog :position="insert.position" />
        </template>
      </DialogComp>

      <DialogComp
        :id="`edit-key-${macroRecorder.state.editKey}`"
        @on-open="onOpenDialog"
        @on-close="onCloseDialog"
      >
        <template #trigger>
          <ButtonComp variant="secondary" size="sm"> <IconPencil />Edit </ButtonComp>
        </template>
        <template #content>
          <EditKeyDialog />
        </template>
      </DialogComp>

      <DialogComp @on-open="onOpenDialog" @on-close="onCloseDialog">
        <template #trigger>
          <ButtonComp size="sm" variant="danger"> <IconTrash />Delete </ButtonComp>
        </template>
        <template #content>
          <DeleteKeyDialog />
        </template>
      </DialogComp>
    </div>
    <DialogComp
      v-if="
        macroRecorder.state.editDelay !== false && typeof macroRecorder.getEditDelay() === 'object'
      "
      @on-open="onOpenDialog"
      @on-close="onCloseDialog"
    >
      <template #trigger>
        <ButtonComp variant="secondary" size="sm"> <IconAlarm />Edit </ButtonComp>
      </template>
      <template #content>
        <EditDelayDialog />
      </template>
    </DialogComp>
  </div>
</template>

<script setup>
import {
  IconAlarm,
  IconArrowLeftCircle,
  IconArrowRightCircle,
  IconPencil,
  IconPlus,
  IconTrash,
} from '@tabler/icons-vue'
import DialogComp from '@/components/base/DialogComp.vue'
import ButtonComp from '@/components/base/ButtonComp.vue'

import { useMacroRecorderStore } from '@/stores/macrorecorder'
import EditKeyDialog from '../components/EditKeyDialog.vue'
import EditDelayDialog from '../components/EditDelayDialog.vue'
import DeleteKeyDialog from '../components/DeleteKeyDialog.vue'
import ContextMenu from '@/components/base/ContextMenu.vue'
import InsertKeyDialog from '../components/InsertKeyDialog.vue'
import { onMounted, reactive, ref } from 'vue'

const macroRecorder = useMacroRecorderStore()

const insert = reactive({ position: null })
const ctxtMenu = ref()

onMounted(() => {
  macroRecorder.$subscribe((mutation) => {
    if (mutation.events && mutation.events.key == 'editKey' && mutation.events.newValue === false) {
      insert.position = null
    }
  })
})

function onOpenDialog() {
  if (insert.position !== null) ctxtMenu.value.toggle()
}
function onCloseDialog() {
  macroRecorder.state.editKey = false
  macroRecorder.state.editDelay = false

  insert.position = null
}
</script>

<style scoped>
@reference "@/assets/main.css";

.macro-edit__dialogs {
  @apply flex
  flex-grow
  justify-end;
}
</style>
