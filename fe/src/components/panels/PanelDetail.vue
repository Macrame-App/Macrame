<template>
  <div></div>
</template>

<script setup>
import { appUrl } from '@/services/ApiService'
import axios from 'axios'
import { nextTick, onMounted, reactive, ref } from 'vue'

const panel = reactive({
  style: '',
  styleEl: null,
  html: '',
})

const testPanel = ref()

onMounted(() => {
  axios.post(appUrl() + '/panel/get').then(async (data) => {
    // console.log(data.data.html)
    if (data.data) {
      setPanelStyle(data.data.css)
      panel.html = data.data.html

      await nextTick()

      addButtonEventListeners()
    }
  })
})

const setPanelStyle = (styleStr) => {
  const styleEl = document.createElement('style')
  styleEl.setAttribute('custom_panel_style', true)
  styleEl.innerHTML = styleStr
  document.head.appendChild(styleEl)

  panel.styleEl = styleEl
}

const addButtonEventListeners = () => {
  testPanel.value.querySelectorAll('[mcrm__button]').forEach((button) => {
    button.addEventListener('click', () => {
      console.log(button.id)
      if (button.id == 'button_1') {
        axios.post(appUrl() + '/macro/play', { macro: 'task_manager' })
      }
    })
  })
}
</script>

<style>
@reference "@/assets/main.css";

[mcrm__button] {
  @apply cursor-pointer;
}
</style>
