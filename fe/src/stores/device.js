import { ref } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios'
import { appUrl, encrypt } from '@/services/ApiService'
import { v4 as uuidv4 } from 'uuid'
import { encryptAES, getDateStr } from '@/services/EncryptService'

export const useDeviceStore = defineStore('device', () => {
  // Properties - State values
  const current = ref({
    uuid: false,
    key: false,
  })

  const remote = ref([])
  const server = ref({
    status: false,
  })

  // Current device
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

  const key = () => {
    if (!current.value.key && localStorage.getItem('deviceKey')) {
      current.value.key = localStorage.getItem('deviceKey')
    }
    return current.value.key
  }

  const setDeviceKey = (key) => {
    current.value.key = key
    localStorage.setItem('deviceKey', key)
  }

  const removeDeviceKey = () => {
    current.value.key = false
    localStorage.removeItem('deviceKey')
  }

  // Server application
  const serverGetRemotes = async (remoteUuid) => {
    axios.post(appUrl() + '/device/list', { uuid: remoteUuid }).then((data) => {
      if (data.data.devices) remote.value = data.data.devices
    })
  }

  const serverStartLink = async (deviceUuid) => {
    const request = await axios.post(appUrl() + '/device/link/start', { uuid: deviceUuid })
    return request.data
  }

  // Remote application
  const remoteCheckServerAccess = async () => {
    const check = await axios.post(appUrl() + '/device/access/check', { uuid: uuid() })
    server.value.access = check.data
    return check.data
  }

  const remoteRequestServerAccess = async (deviceName, deviceType) => {
    const request = await axios.post(appUrl() + '/device/access/request', {
      uuid: uuid(),
      name: deviceName,
      type: deviceType,
    })
    return request
  }
  const remotePingLink = async (cb) => {
    // const linkRequest = await axios.post(appUrl() + '/device/link/ping', { uuid: deviceUuid })
    // if (linkRequest.data)
    const pingInterval = setInterval(() => {
      axios.post(appUrl() + '/device/link/ping', { uuid: uuid() }).then((data) => {
        if (data.data) {
          clearInterval(pingInterval)
          cb(data.data)
        }
      })
    }, 1000)
  }

  const remoteHandshake = async (key) => {
    const handshake = await axios.post(appUrl() + '/device/handshake', {
      uuid: uuid(),
      shake: encryptAES(key, getDateStr()),
    })
    console.log(handshake)

    return handshake.data
  }

  return {
    remote,
    server,
    uuid,
    setDeviceId,
    key,
    setDeviceKey,
    removeDeviceKey,
    serverGetRemotes,
    serverStartLink,
    remoteCheckServerAccess,
    remoteRequestServerAccess,
    remotePingLink,
    remoteHandshake,
  }
})
