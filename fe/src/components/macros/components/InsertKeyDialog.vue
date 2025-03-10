<template>
  <div id="insert-key-dialog" class="dialog__content w-80">
    <h4 class="text-slate-50 mb-4">Insert key {{ position }}</h4>
    <div class="flex justify-center w-full mb-4">
      <MacroKey v-if="keyObjs.selected" :key-obj="keyObjs.selected" />
      <MacroKey v-if="keyObjs.adjacent" :key-obj="keyObjs.adjacent" />
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue'
import { useMacroRecorderStore } from '@/stores/macrorecorder'
import MacroKey from './MacroKey.vue'
import { filterKey } from '@/services/MacroRecordService'

const props = defineProps({
  position: String,
})

const macroRecorder = useMacroRecorderStore()

const keyObjs = reactive({
  selected: null,
  insert: null,
  adjacent: null,
  adjacentDelay: null,
})

onMounted(() => {
  keyObjs.selected = filterKey(macroRecorder.getEditKey())

  const adjacentKey = macroRecorder.getAdjacentKey(props.position, true)
  if (adjacentKey) keyObjs.adjacent = filterKey(adjacentKey.key)
  if (adjacentKey.delay) keyObjs.adjacentDelay = adjacentKey.delay
})
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
