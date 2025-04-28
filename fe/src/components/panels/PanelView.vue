<template>
  <div id="panel-view">
    <div class="panel-preview__content" ref="panelView" v-html="viewPanel.html"></div>
  </div>
</template>

<script setup>
import { RunMacro } from '@/services/MacroService'
import {
  PanelButtonListeners,
  PanelDialogListeners,
  RemovePanelScripts,
  RemovePanelStyle,
  SetPanelStyle,
  StripPanelHTML,
} from '@/services/PanelService'
import { usePanelStore } from '@/stores/panel'
import { onMounted, onUnmounted, ref } from 'vue'

const panel = usePanelStore()

const props = defineProps({
  dirname: String,
})

const panelView = ref(null)

const viewPanel = ref({})

onMounted(async () => {
  const currentPanel = await panel.get(props.dirname)
  viewPanel.value = currentPanel

  viewPanel.value.html = StripPanelHTML(viewPanel.value.html, viewPanel.value.aspectRatio)
  SetPanelStyle(viewPanel.value.style)

  setTimeout(() => {
    viewPanelListeners()

    if (typeof window.onPanelLoaded === 'function') {
      window.onPanelLoaded()
    }
  }, 50)
})

onUnmounted(() => {
  RemovePanelStyle()
  RemovePanelScripts()
})

const viewPanelListeners = () => {
  const callback = (button) => {
    RunMacro(viewPanel.value.macros[button.id])
  }

  PanelButtonListeners(panelView.value, callback)
  PanelDialogListeners(panelView.value)
}
</script>

<style scoped>
@reference "@/assets/main.css";

#panel-view {
  @apply fixed
  inset-0
  size-full
  bg-black;

  .panel-preview__content {
    @apply relative
      grid
      justify-center
      size-full;

    #panel-html__body {
      @apply size-full
        max-w-full max-h-full;
    }
  }
}
</style>
