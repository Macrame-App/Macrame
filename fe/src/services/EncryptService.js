import { useDeviceStore } from '@/stores/device'
import { AES, enc, pad } from 'crypto-js'
import { isLocal } from './ApiService'

export const encryptAES = (key, str) => {
  key = keyPad(key)

  let iv = enc.Utf8.parse(window.__CONFIG__.MCRM__IV)
  let encrypted = AES.encrypt(str, key, {
    iv: iv,
    padding: pad.Pkcs7,
  })
  return encrypted.toString()
}

export const decryptAES = (key, str) => {
  key = keyPad(key)

  let iv = enc.Utf8.parse(window.__CONFIG__.MCRM__IV)
  let encrypted = AES.decrypt(str.toString(), key, {
    iv: iv,
    padding: pad.Pkcs7,
  })
  return encrypted.toString(enc.Utf8)
}

export const AuthCall = (data) => {
  if (isLocal()) return data

  const device = useDeviceStore()

  return {
    uuid: device.uuid(),
    d: encryptAES(device.key(), JSON.stringify(data)),
  }
}

function keyPad(key) {
  let returnKey = key

  if (key.length == 4) {
    returnKey = key + window.__CONFIG__.MCRM__SALT
  }

  return enc.Utf8.parse(returnKey)
}

export const getDateStr = () => {
  const date = new Date()
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  return `${year}${month}${day}`
}
