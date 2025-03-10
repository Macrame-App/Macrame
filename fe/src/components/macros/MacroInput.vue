<template>
  <div class="macro-input">
    <div class="macro-input__output" ref="macroOutput"></div>
    <input class="macro-input__input" type="text" ref="macroInput" @keydown.prevent="keyDown" />
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'

const macroOutput = ref(null)
const macroInput = ref(null)

onMounted(() => {
  macroInput.value.focus()
})

const keyDown = (e) => {
  console.log(e)

  const modKeys = {
    Control: 'Ctrl',
    Shift: 'Shift',
    Alt: 'Alt',
    Meta: 'Win',
  }

  const specialKeys = {
    PageUp: 'PgUp',
    PageDown: 'PgDn',
    ScrollLock: 'Scr Lk',
  }

  let key = e.key

  if (e.shiftKey) {
    key = e.key.toLowerCase()
  }

  const newKeyEl = document.createElement('kbd')

  if (e.code === 'Space') {
    newKeyEl.innerHTML = 'Space'
  } else if (e.location === 1 && Object.keys(modKeys).includes(key)) {
    newKeyEl.innerHTML = '<sup>left</sup> ' + (modKeys[key] || key)
  } else if (e.location === 2 && Object.keys(modKeys).includes(key)) {
    newKeyEl.innerHTML = '<sup>right</sup> ' + (modKeys[key] || key)
  } else if (e.location === 3) {
    newKeyEl.innerHTML = '<sup>num</sup> ' + (modKeys[key] || key)
  } else if (Object.keys(specialKeys).includes(key)) {
    newKeyEl.innerHTML = specialKeys[key] || key
  } else {
    newKeyEl.innerHTML = key
  }

  macroOutput.value.appendChild(newKeyEl)
}
</script>

<style>
@reference "@/assets/main.css";

.macro-input {
  @apply relative
  w-full
  h-96
  my-4
  rounded-lg
  bg-slate-900/50
  border
  border-white/10
  overflow-auto;

  .macro-input__input,
  .macro-input__output {
    @apply absolute
    inset-0
    size-full;
  }

  .macro-input__output {
    @apply flex 
    flex-wrap 
    items-start
    gap-2 
    p-4
    h-fit;
  }

  kbd {
    @apply flex 
    items-center
    gap-2
    px-4 py-1 
    h-9 
    bg-white/10
    font-sans
    font-bold
    text-lg
    uppercase
    rounded-md 
    border;

    sup {
      @apply text-xs font-light -ml-1.5 mt-1;
    }
  }
}
</style>
