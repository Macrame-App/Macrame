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
