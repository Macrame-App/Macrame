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
  </div>
</template>

<script setup>
import ButtonComp from '@/components/base/ButtonComp.vue'
import { IconRestore } from '@tabler/icons-vue'

import { useMacroRecorderStore } from '@/stores/macrorecorder'

const macroRecorder = useMacroRecorderStore()
</script>

<style scoped>
@reference "@/assets/main.css";
</style>
