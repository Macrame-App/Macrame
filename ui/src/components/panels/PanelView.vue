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
    <div class="panel-container" ref="panelContainer" v-html="viewPanel.html"></div>
  </div>
</template>

<script setup>
import { isLocal } from '@/services/ApiService'
import { RunMacro } from '@/services/MacroService'
import {
  CheckLocalPanel,
  PanelButtonListeners,
  PanelDialogListeners,
  RemovePanelScripts,
  RemovePanelStyle,
  SavePanelToLocal,
  SetPanelStyle,
  StripPanelHTML,
} from '@/services/PanelService'
import { usePanelStore } from '@/stores/panel'
import { useSettingStore } from '@/stores/settings'
import { onMounted, onUnmounted, ref } from 'vue'

const panel = usePanelStore()
const settings = useSettingStore()

const props = defineProps({
  dirname: String,
})

const panelContainer = ref(null)

const viewPanel = ref({})

const wakeLock = ref(null)

onMounted(async () => {
  requestWakeLock()

  const currentPanel = await panel.get(props.dirname)
  viewPanel.value = currentPanel

  if (!isLocal() && settings.get('openLastPanel') && !CheckLocalPanel()) SavePanelToLocal()

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

  wakeLock.value.release()
})

const viewPanelListeners = () => {
  const callback = (button) => {
    RunMacro(viewPanel.value.macros[button.id])
  }

  PanelButtonListeners(panelContainer.value, callback)
  PanelDialogListeners(panelContainer.value)
}

const requestWakeLock = async () => {
  try {
    if ('wakeLock' in navigator) {
      wakeLock.value = await navigator.wakeLock.request('screen')
    } else {
      console.warn('Wake Lock API not supported')
    }
  } catch (err) {
    console.error(`${err.name}, ${err.message}`)
  }
}
</script>

<style scoped>
@reference "@/assets/main.css";

#panel-view {
  @apply fixed
  inset-0
  size-full
  bg-black;

  .panel-container {
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
