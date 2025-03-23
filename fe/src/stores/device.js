import { ref } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios'
import { appUrl } from '@/services/ApiService'
import { v4 as uuidv4 } from 'uuid'

export const useDeviceStore = defineStore('device', () => {
  // Properties - State values
  const current = ref({
    uuid: false,
  })

  const remote = ref([])
  const server = ref({
    status: false,
  })

  const uuid = () => {
    if (!current.value.uuid && localStorage.getItem('deviceId')) {
      current.value.uuid = localStorage.getItem('deviceId')
    } else if (!current.value.uuid) {
      current.value.uuid = setDeviceId()
    }
    return current.value.uuid
  }

  const setDeviceId = () => {
    const uuid = uuidv4()
    localStorage.setItem('deviceId', uuid)
    return uuid
  }

  const getRemoteDevices = async () => {
    // remote.value = { penis: 'penis' }
    // return
    axios.post(appUrl() + '/device/list').then((data) => {
      if (data.data.devices) remote.value = data.data.devices
    })
  }

  const checkServerAccess = async () => {
    const check = await axios.post(appUrl() + '/device/access/check', { uuid: uuid() })
    server.value.access = check.data
    return check.data
  }

  const requestServerAccess = async () => {
    const request = await axios.post(appUrl() + '/device/access/request', { uuid: uuid() })
    return request
  }

  return {
    remote,
    server,
    uuid,
    setDeviceId,
    getRemoteDevices,
    checkServerAccess,
    requestServerAccess,
  }
})
