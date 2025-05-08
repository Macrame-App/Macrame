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
import CryptoJS from 'crypto-js'

export const appUrl = () => {
  const port = window.location.port == 5173 ? window.__CONFIG__.MCRM__PORT : window.location.port

  return `http://${window.location.hostname}:${port}`
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
