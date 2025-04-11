import { useDeviceStore } from '@/stores/device'
import CryptoJS from 'crypto-js'

export const appUrl = () => {
  return window.location.port !== 6970 ? `http://${window.location.hostname}:6970` : ''
}

export const isLocal = () => {
  return window.location.hostname === '127.0.0.1' || window.location.hostname === 'localhost'
}

export const encrypt = (data, key = false) => {
  const pk = !key ? localStorage.getItem('Macrame__pk') : key

  if (pk) {
    return CryptoJS.RSA.encrypt(JSON.stringify(data), pk).toString()
  } else {
    return false
  }
}

export const checkAuth = async () => {
  const device = useDeviceStore()

  const handshake = await device.remoteHandshake()

  if (handshake === true) return true

  if (device.key()) return true

  return false
}
