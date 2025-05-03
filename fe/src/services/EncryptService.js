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

export const AuthCall = (data = false) => {
  if (isLocal()) return data

  if (!data) data = {empty: true}

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
