/*
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
*/

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

  const serverGetIP = async () => {
    const request = await axios.post(appUrl() + '/device/server/ip')
    return `http://${request.data}:${window.__CONFIG__.MCRM__PORT}`
  }

  // Server application
  const serverGetRemotes = async (count = false) => {
    const request = await axios.post(appUrl() + '/device/list')

    if (!request.data.devices) return false

    remote.value = request.data.devices

    if (!count) return remote.value
    else return Object.keys(remote.value).length
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
    const pingInterval = setInterval(() => {
      axios.post(appUrl() + '/device/link/ping', { uuid: uuid() }).then((data) => {
        if (data.data) {
          clearInterval(pingInterval)
          cb(data.data)
        }
      })
    }, 1000)
  }

  const remoteHandshake = async (keyStr = false) => {
    if (!keyStr) keyStr = key()

    if (!keyStr) return false

    const handshake = await axios.post(appUrl() + '/device/handshake', {
      uuid: uuid(),
      shake: encryptAES(keyStr, getDateStr()),
    })

    if (!handshake.data) removeDeviceKey()

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
    serverGetIP,
    serverGetRemotes,
    serverStartLink,
    remoteCheckServerAccess,
    remoteRequestServerAccess,
    remotePingLink,
    remoteHandshake,
  }
})
