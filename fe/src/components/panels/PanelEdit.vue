<template>
  <div id="panel-edit" class="mcrm-block block__dark !p-0 !gap-0" v-if="editPanel">
    <div class="panel-preview">
      <div class="panel-preview__content" ref="panelPreview" v-html="editPanel.html"></div>
    </div>
    <div class="panel-settings">
      <AccordionComp title="Panel info" ref="infoAccordion">
        <div class="grid grid-cols-[auto_1fr] gap-2 p-4">
          <span>Name:</span><strong class="text-right">{{ editPanel.name }}</strong>

          <span>Aspect ratio:</span><strong class="text-right">{{ editPanel.aspectRatio }}</strong>

          <template v-if="editPanel.macros">
            <span>Linked Macros:</span>
            <strong class="text-right">{{ Object.keys(editPanel.macros).length }}</strong>
          </template>
        </div>
      </AccordionComp>
      <div>
        <AccordionComp
          v-if="editButton.id"
          title="Button"
          ref="buttonAccordion"
          :open="editButton.id != ''"
        >
          <div class="grid gap-4 p-4">
            <div class="grid grid-cols-[auto_1fr] gap-2">
              <span>Button ID:</span>
              <strong class="text-right">{{ editButton.id }}</strong>
            </div>
            <div class="grid">
              <FormSelect
                name="button_macro"
                label="Button macro"
                :search="true"
                :options="macroList"
                :value="editButton.macro"
                @change="checkNewMacro(editButton.id, $event)"
              />
              <div class="grid grid-cols-2 mt-4">
                <ButtonComp
                  v-if="editButton.macro != ''"
                  class="col-start-1 w-fit"
                  size="sm"
                  variant="danger"
                  @click="unlinkMacro(editButton.id)"
                  ref="unlinkButton"
                >
                  <IconTrash /> Unlink
                </ButtonComp>
                <ButtonComp
                  v-if="editButton.changed"
                  class="col-start-2 w-fit justify-self-end"
                  size="sm"
                  variant="primary"
                  @click="linkMacro(editButton.id)"
                  ref="linkButton"
                >
                  <IconLink /> Link
                </ButtonComp>
              </div>
            </div>
          </div>
        </AccordionComp>
      </div>
      <footer class="flex items-end justify-end h-full p-4">
        <ButtonComp v-if="panelMacros.changed" variant="success" @click="savePanelChanges()">
          <IconDeviceFloppy /> Save changes
        </ButtonComp>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { CheckMacroListChange, GetMacroList } from '@/services/MacroService'
import {
  PanelButtonListeners,
  RemovePanelStyle,
  SetPanelStyle,
  StripPanelHTML,
} from '@/services/PanelService'
import { usePanelStore } from '@/stores/panel'
import { onMounted, onUnmounted, onUpdated, reactive, ref } from 'vue'
import AccordionComp from '../base/AccordionComp.vue'
import FormSelect from '../form/FormSelect.vue'
import ButtonComp from '../base/ButtonComp.vue'
import { IconDeviceFloppy, IconLink, IconTrash } from '@tabler/icons-vue'
import axios from 'axios'
import { appUrl } from '@/services/ApiService'

const props = defineProps({
  dirname: String,
})

const panel = usePanelStore()

const panelPreview = ref(false)
const editPanel = ref({})
const panelMacros = reactive({
  old: {},
  changed: false,
})

const macroList = ref({})

const infoAccordion = ref(false)
const buttonAccordion = ref(false)

const unlinkButton = ref(null)
const linkButton = ref(null)

const editButton = reactive({
  id: '',
  macro: '',
  newMacro: '',
  changed: false,
})

onMounted(async () => {
  const currentPanel = await panel.get(props.dirname)
  editPanel.value = currentPanel
  editPanel.value.dir = props.dirname
  editPanel.value.html = StripPanelHTML(editPanel.value.html, editPanel.value.aspectRatio)

  panelMacros.old = JSON.stringify(currentPanel.macros)

  infoAccordion.value.toggleAccordion(true)

  const macros = await GetMacroList()
  macroList.value = Object.assign(
    {},
    ...Object.keys(macros).map((key) => ({
      [key]: { value: macros[key].macroname, label: macros[key].name },
    })),
  )

  SetPanelStyle(editPanel.value.style)

  EditButtonListeners()
})

onUpdated(() => {
  console.log('updated')
})

onUnmounted(() => {
  RemovePanelStyle()
})

function EditButtonListeners() {
  const callback = (button) => {
    infoAccordion.value.toggleAccordion(false)
    setEditButton(button.id)
  }

  PanelButtonListeners(panelPreview.value, callback)
}

function setEditButton(id) {
  editButton.id = id
  editButton.macro = editPanel.value.macros[id] ? editPanel.value.macros[id] : ''
}

function checkNewMacro(id, macro) {
  editButton.changed = editPanel.value.macros[id] != macro
  editButton.newMacro = macro
}

function linkMacro(id) {
  editPanel.value.macros[id] = editButton.newMacro
  editButton.macro = editButton.newMacro
  editButton.newMacro = ''

  panelMacros.changed = CheckMacroListChange(panelMacros.old, editPanel.value.macros)
}

function unlinkMacro(id) {
  delete editPanel.value.macros[id]
  buttonAccordion.value.toggleAccordion(false)
  panelMacros.changed = CheckMacroListChange(panelMacros.old, editPanel.value.macros)
}

function savePanelChanges() {
  const panelData = {
    dir: editPanel.value.dir,
    name: editPanel.value.name,
    description: editPanel.value.description,
    aspectRatio: editPanel.value.aspectRatio,
    macros: editPanel.value.macros,
  }

  axios.post(appUrl() + '/panel/save/json', panelData).then((data) => {
    console.log(data)
  })
}
</script>

<style>
@reference "@/assets/main.css";

[mcrm__button] {
  @apply cursor-pointer;
}

#panel-edit {
  @apply grid
  grid-cols-[1fr_30ch]
  size-full
  overflow-hidden;

  .panel-preview {
    @apply border-r
    border-slate-700;

    .panel-preview__content {
      @apply relative
      grid
      justify-center
      size-full
      p-8;

      #panel-html__body {
        @apply size-full
        max-w-full max-h-full;
      }
    }
  }

  .panel-settings {
    @apply grid
    grid-rows-[auto_auto_1fr]
    bg-black/30;
  }
}
</style>
