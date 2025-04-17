import { appUrl } from '@/services/ApiService'
import { AuthCall } from '@/services/EncryptService'
import axios from 'axios'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const usePanelStore = defineStore('panel', () => {
  const current = ref({
    dir: false,
    name: false,
    description: false,
    aspectRatio: false,
    macros: false,
    thumb: false,
    html: false,
    style: false,
  })

  const list = ref([])

  const get = async (dir) => {
    const data = AuthCall({ dir: dir })

    const resp = await axios.post(appUrl() + '/panel/get', data)

    if (!resp.data && !current.value.html) return false

    current.value.name = resp.data.name
    current.value.description = resp.data.description
    current.value.aspectRatio = resp.data.aspectRatio
    current.value.macros = resp.data.macros
    current.value.thumb = resp.data.thumb
    current.value.html = resp.data.html
    current.value.style = resp.data.style

    return current.value
  }

  const getList = async () => {
    if (list.value.length > 0) return list.value

    const data = AuthCall()

    const resp = await axios.post(appUrl() + '/panel/list', data)
    list.value = resp.data

    return list.value
  }

  return {
    current,
    list,
    get,
    getList,
  }
})
