import axios from 'axios'
import { appUrl } from './ApiService'

export const ConfigGet = async (key) => {
  console.log(window.__CONFIG__)

  const config = await axios.get(appUrl() + '/config.json')

  if (!config.data) return false

  return config.data[key]
}
