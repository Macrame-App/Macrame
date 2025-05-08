<!--
Macrame is a program that enables the user to create keyboard macros and button panels. 
The macros are saved as simple JSON files and can be linked to the button panels. The panels can 
be created with HTML and CSS.

Copyright (C) 2025 Jesse Malotaux

This program is free software: you can redistribute it and/or modify 
it under the terms of the GNU General Public License as published by 
the Free Software Foundation, either version 3 of the License, or 
(at your option) any later version.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
GNU General Public License for more details.

You should have received a copy of the GNU General Public License 
along with this program. If not, see <https://www.gnu.org/licenses/>.
-->

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
