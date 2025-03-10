<template>
  <div id="delete-key-dialog" class="dialog__content">
    <h4 class="text-slate-50 mb-4">Delete key</h4>
    <div class="flex justify-center w-full mb-4">
      <MacroKey v-if="keyObj" :key-obj="keyObj" />
    </div>
    <p class="text-sm text-slate-300">Are you sure you want to delete this key?</p>
    <div class="flex justify-end gap-2 mt-6">
      <ButtonComp variant="danger" size="sm" @click="macroRecorder.deleteEditKey()">
        Delete key
      </ButtonComp>
    </div>
  </div>
</template>

<script setup>
import ButtonComp from '@/components/base/ButtonComp.vue'
import { useMacroRecorderStore } from '@/stores/macrorecorder'
import { onMounted, ref } from 'vue'
import MacroKey from './MacroKey.vue'
import { filterKey } from '@/services/MacroRecordService'

const macroRecorder = useMacroRecorderStore()

const keyObj = ref(null)

onMounted(() => {
  keyObj.value = filterKey(macroRecorder.getEditKey())
  console.log(macroRecorder.getEditKey());
  console.log(keyObj.value);
  console.log('---------');
})
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
'
