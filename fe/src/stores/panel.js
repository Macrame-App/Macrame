import { appUrl } from '@/services/ApiService'
import axios from 'axios'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const usePanelStore = defineStore('panel', () => {
  const list = ref([])

  const getPanels = async () => {
    console.log(list.value.length)

    if (list.value.length > 0) return list.value

    const resp = await axios.post(appUrl() + '/panel/list')
    list.value = resp.data.data

    return list.value
  }

  return {
    list,
    getPanels,
  }
})
