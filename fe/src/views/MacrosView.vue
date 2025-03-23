<template>
  <div id="macros" class="panel">
    <h1 class="panel__title">Macros</h1>
    <div class="panel__content !p-0">
      <div class="macro-panel__content">
        <MacroOverview />
        <MacroRecorder />
      </div>
    </div>
  </div>
</template>

<script setup>
import MacroOverview from '@/components/macros/MacroOverview.vue'
import MacroRecorder from '../components/macros/MacroRecorder.vue'
import { onMounted, ref } from 'vue'

const recordMacro = ref(false)

const macroInput = ref(null)

onMounted(() => {
  // macroInput.value.focus()
})

const keyDown = (e) => {
  console.log(e)
}
</script>

<style scoped>
@reference "@/assets/main.css";

.macro-panel__content {
  @apply grid
  grid-cols-[25ch_1fr]
  gap-6
  pt-2;
}
</style>
