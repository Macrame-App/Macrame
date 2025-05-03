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

import axios from 'axios'
import { appUrl, isLocal } from './ApiService'
import { AuthCall } from './EncryptService'

export const GetMacroList = async (count = false) => {
  const request = await axios.post(appUrl() + '/macro/list')
  
  if (!request.data) return 0

  if (!count) return sortMacroList(request.data)
  else return request.data.length
}

const sortMacroList = (list) => {
  return [...list].sort((a, b) => a.name.localeCompare(b.name))
}

export const RunMacro = async (macro) => {
  const data = isLocal() ? { macro: macro } : AuthCall({ macro: macro })
  const request = await axios.post(appUrl() + '/macro/play', data)
  return request.data
}

export const CheckMacroListChange = (oldList, newList) => {
  return oldList !== JSON.stringify(newList)
}
