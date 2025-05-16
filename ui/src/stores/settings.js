import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useSettingStore = defineStore('settings', () => {
  const list = ref({})

  const get = (key, all = false) => {
    if (key === undefined) return list.value

    if (!all) return list.value[key].value
    else return list.value[key]
  }

  const set = (key, value) => {
    if (value === 'true' || value === 'false') value = value === 'true'

    list.value[key].value = value

    saveSettings()
  }

  const loadSettings = (returnObj = false) => {
    const settings = localStorage.getItem('settings')

    if (settings) list.value = JSON.parse(settings)
    else list.value = loadDefaultSettings()

    if (returnObj) return list.value
  }

  const loadDefaultSettings = () => {
    const defaultSettings = {
      openLastPanel: {
        title: 'Open last panel',
        label: 'Open last panel',
        description: 'Open the last panel on startup',
        type: 'checkbox',
        value: true,
        remote: true,
      },
      mcrmPort: {
        title: 'Macrame port',
        label: 'Port',
        description:
          'The port that is used by Macrame, changing this will require a restart of the application.',
        type: 'number',
        value: window.__CONFIG__.MCRM__PORT,
        remote: false,
        disabled: true,
      },
    }

    return defaultSettings
  }

  const saveSettings = () => {
    localStorage.setItem('settings', JSON.stringify(list.value))
  }

  const resetSettings = () => {
    localStorage.removeItem('settings')
    list.value = loadDefaultSettings()
  }

  return {
    list,
    get,
    set,
    loadSettings,
    loadDefaultSettings,
    saveSettings,
    resetSettings,
  }
})
