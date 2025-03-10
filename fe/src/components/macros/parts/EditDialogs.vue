<template>
  <div class="macro-edit__dialogs" v-if="macroRecorder.state.edit !== false">
    <div
      class="flex gap-2"
      v-if="macroRecorder.state.editKey !== false && typeof macroRecorder.getEditKey() === 'object'"
    >
      <ContextMenu>
        <template #trigger>
          <ButtonComp variant="success" size="sm"> <IconPlus /> Insert key </ButtonComp>
        </template>
        <template #content>
          <ul>
            <li @click="insertPosition = 'before'"><IconArrowLeftCircle /> Before</li>
            <li @click="insertPosition = 'after'"><IconArrowRightCircle /> After</li>
          </ul>
        </template>
      </ContextMenu>
      <DialogComp v-if="insertPosition !== null" :open="true" @on-close="insertPosition = null">
        <template #content>
          <InsertKeyDialog :position="insertPosition" />
        </template>
      </DialogComp>

      <DialogComp>
        <template #trigger>
          <ButtonComp size="sm" variant="danger" @click="console.log('delete')">
            <IconTrash />Delete key
          </ButtonComp>
        </template>
        <template #content>
          <DeleteKeyDialog />
        </template>
      </DialogComp>
      <DialogComp
        :id="`edit-key-${macroRecorder.state.editKey}`"
        @on-close="macroRecorder.state.editKey = false"
      >
        <template #trigger>
          <ButtonComp variant="primary" size="sm"> <IconKeyboard />Edit key </ButtonComp>
        </template>
        <template #content>
          <EditKeyDialog />
        </template>
      </DialogComp>
    </div>
    <DialogComp
      v-if="
        macroRecorder.state.editDelay !== false && typeof macroRecorder.getEditDelay() === 'object'
      "
      @on-close="macroRecorder.state.editDelay = false"
    >
      <template #trigger>
        <ButtonComp variant="primary" size="sm"> <IconAlarm />Edit delay </ButtonComp>
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
  IconKeyboard,
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
import { ref } from 'vue'

const macroRecorder = useMacroRecorderStore()

const insertPosition = ref(null)
</script>

<style scoped>
@reference "@/assets/main.css";

.macro-edit__dialogs {
  @apply flex
  flex-grow
  justify-end;
}
</style>
