<template>
  <div id="macros" class="panel">
    <h1>Macros</h1>
    <MacroRecorder />
  </div>
</template>

<script setup>
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

input {
  @apply border rounded-md bg-white/10 p-2;
}
</style>
