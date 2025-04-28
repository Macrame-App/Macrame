<template>
  <div class="macro-overview mcrm-block block__dark">
    <h4 class="border-b-2 border-transparent">Saved Macros</h4>
    <div class="macro-overview__list">
      <LoadComp :loading="macros.loading" text="Loading macros..." />
      <div class="macro-item" v-for="(macro, i) in macros.list" :key="i">
        <ButtonComp
          :variant="macroRecorder.macroName === macro.name ? 'secondary' : 'dark'"
          class="overview__macro-open"
          size="sm"
          @click="macroRecorder.openMacro(macro.macroname, macro.name)"
        >
          <IconKeyboard /> <span>{{ macro.name }}</span>
        </ButtonComp>
        <div class="overview__macro-delete">
          <ButtonComp
            class="!text-red-500 hover:!text-red-300"
            variant="ghost"
            size="sm"
            @click="startDelete(macro.name)"
          >
            <IconTrash />
          </ButtonComp>
        </div>
      </div>
    </div>
    <DialogComp ref="deleteDialog">
      <template #content>
        <div class="grid gap-2">
          <h4 class="pr-4">Are you sure you want to delete:</h4>
          <h3 class="mb-2 text-center text-sky-500">{{ macroToBeDeleted }}</h3>
          <div class="flex justify-between">
            <ButtonComp size="sm" variant="subtle" @click="deleteDialog.toggleDialog(false)">
              No
            </ButtonComp>
            <ButtonComp size="sm" variant="danger" @click="deleteMacro()">Yes</ButtonComp>
          </div>
        </div>
      </template>
    </DialogComp>
  </div>
</template>

<script setup>
// TODO
// - delete macro

import { IconKeyboard, IconTrash } from '@tabler/icons-vue'
import ButtonComp from '../base/ButtonComp.vue'
import { onMounted, reactive, ref } from 'vue'
import { GetMacroList } from '@/services/MacroService'
import LoadComp from '../base/LoadComp.vue'
import { useMacroRecorderStore } from '@/stores/macrorecorder'
import DialogComp from '../base/DialogComp.vue'

const macros = reactive({
  loading: true,
  list: [],
})

const macroRecorder = useMacroRecorderStore()

const macroToBeDeleted = ref('')
const deleteDialog = ref()

onMounted(() => {
  loadMacroList()
})

const loadMacroList = async () => {
  const list = await GetMacroList()
  macros.list = list
  macros.loading = false
}

const startDelete = (macroFilename) => {
  macroToBeDeleted.value = macroFilename
  deleteDialog.value.toggleDialog(true)
}

const deleteMacro = async () => {
  const resp = await macroRecorder.deleteMacro(macroToBeDeleted.value)

  if (resp) {
    deleteDialog.value.toggleDialog(false)

    if (macroToBeDeleted.value === macroRecorder.macroName) macroRecorder.resetMacro()

    macroToBeDeleted.value = ''
    loadMacroList()
  }
}
</script>

<style scoped>
@reference "@/assets/main.css";

.macro-overview {
  @apply relative
  grid
  grid-rows-[auto_1fr];

  &::after {
    @apply content-['']
    absolute
    top-0
    left-full
    h-full
    w-px
    bg-slate-600;
  }

  .macro-overview__list {
    @apply flex
    flex-col 
    pr-1
    -mr-1
    gap-1
    h-[calc(100vh-11.7rem)]
    overflow-auto;
  }

  .macro-item {
    @apply grid items-center grid-cols-[1fr_0fr] transition-[grid-template-columns] delay-0 duration-300;

    &:hover {
      @apply grid-cols-[1fr_auto] delay-500;
    }

    button.overview__macro-open {
      @apply w-full grid grid-cols-[1rem_1fr] justify-items-start;

      span {
        @apply truncate w-full text-left;
      }
    }

    div.overview__macro-delete {
      @apply grid overflow-hidden transition;
    }
  }
}
</style>
