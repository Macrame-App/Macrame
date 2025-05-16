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
  <div id="settings" class="panel">
    <h1 class="flex items-end justify-between !w-full panel__title">
      <div>Settings</div>
      <ButtonComp variant="subtle" size="sm" @click="settings.resetSettings()">
        <IconRestore /> Reset
      </ButtonComp>
    </h1>

    <div class="panel__content">
      <div class="settings-container">
        <template
          class="setting-block mcrm-block block__dark"
          v-for="(setting, name) in settingList"
          :key="setting"
        >
          <div
            v-if="(setting.remote && !isLocal()) || (!setting.remote && isLocal())"
            class="setting-block mcrm-block block__dark"
          >
            <h4>{{ setting.title }}</h4>
            <p class="text-sm">
              <em>{{ setting.description }}</em>
            </p>
            <FormInput
              :type="setting.type"
              :label="setting.label"
              :name="name"
              :value="setting.value"
              :options="setting.options"
              :disabled="setting.disabled"
              :horizontal="true"
              @onChange="updateSetting(name, setting.type, $event)"
            />
          </div>
        </template>
      </div>
    </div>
  </div>
</template>

<script setup>
import ButtonComp from '@/components/base/ButtonComp.vue'
import { IconRestore } from '@tabler/icons-vue'
import FormInput from '@/components/form/FormInput.vue'

import { isLocal } from '@/services/ApiService'

import { useSettingStore } from '@/stores/settings'
import { useNoticationStore } from '@/stores/notifications'

import { computed, onMounted } from 'vue'

const settings = useSettingStore()
const notifications = useNoticationStore()

const settingList = computed(() => settings.list)

onMounted(() => {})

function updateSetting(name, type, target) {
  if (type == 'checkbox' || type == 'radio') target.value = target.checked

  settings.set(name, target.value)

  notifications.add({
    title: 'Settings updated',
    message: 'Settings have been updated',
    variant: 'success',
    time: 5000,
    closable: true,
  })
}
</script>

<style scoped>
@reference "@/assets/main.css";

.settings-container {
  @apply grid
  grid-cols-1
  sm:grid-cols-2
  lg:grid-cols-3
  items-start
  gap-4
  pt-8;

  .setting-block {
    @apply grid
    gap-3
    content-start;
  }
}
</style>
