import { ref } from 'vue'
import { defineStore } from 'pinia'

import { filterKey, isRepeat, invalidMacro, translateJSON } from '../services/MacroRecordService'
import axios from 'axios'
import { appUrl } from '@/services/ApiService'

export const useMacroRecorderStore = defineStore('macrorecorder', () => {
  // Properties - State values
  const state = ref({
    record: false,
    edit: false,
    editKey: false,
    editDelay: false,
    validationErrors: false,
  })

  const macroName = ref('')

  const steps = ref([])

  const delay = ref({
    start: 0,
    fixed: false,
  })

  // Getters - Computed values
  const getEditKey = () => steps.value[state.value.editKey]
  const getAdjacentKey = (pos, includeDelay = false) => {
    let returnVal = false

    const mod = pos == 'before' ? -1 : 1
    const keyIndex = state.value.editKey + 2 * mod
    const delayIndex = includeDelay ? state.value.editKey + 1 * mod : false

    if (steps.value[keyIndex]) returnVal = steps.value[keyIndex]
    if (delayIndex && steps.value[delayIndex])
      returnVal = {
        delay: steps.value[delayIndex],
        key: steps.value[keyIndex],
        delayIndex: delayIndex,
      }

    return returnVal
  }

  const getEditDelay = () => steps.value[state.value.editDelay]

  // Setters - Actions
  const recordStep = (e, direction = false, key = false) => {
    if ((e.ctrlKey, e.shiftKey, e.altKey, e.metaKey)) e.preventDefault()

    const lastStep = steps.value[steps.value.length - 1]

    let stepVal = {}

    if (typeof e === 'object' && !isRepeat(lastStep, e, direction)) {
      if (key === false) recordDelay()

      stepVal = {
        type: 'key',
        key: e.key,
        code: e.code,
        location: e.location,
        direction: direction,
        keyObj: filterKey(e),
      }
    } else if (direction && key !== false) {
      stepVal = steps.value[key]
      stepVal.direction = direction
    } else if (typeof e === 'number') {
      stepVal = { type: 'delay', value: parseFloat(e.toFixed()) }
    }

    if (key !== false) steps.value[key] = stepVal
    else steps.value.push(stepVal)

    console.log(steps.value)
  }

  const recordDelay = () => {
    if (delay.value.fixed !== false)
      recordStep(delay.value.fixed) // Record fixed delay
    else if (delay.value.start == 0)
      delay.value.start = performance.now() // Record start of delay
    else {
      recordStep(performance.now() - delay.value.start) // Record end of delay
      delay.value.start = performance.now() // Reset start
    }
  }

  const insertKey = (e, direction, adjacentDelayIndex) => {
    let min, max, newKeyIndex, newDelayIndex

    if (adjacentDelayIndex === null) {
      min = state.value.editKey == 0 ? 0 : state.value.editKey
      max = state.value.editKey == 0 ? 1 : false

      newKeyIndex = max === false ? min + 2 : min
      newDelayIndex = min + 1
    } else if (state.value.editKey < adjacentDelayIndex) {
      min = state.value.editKey
      max = adjacentDelayIndex
      newKeyIndex = min + 2
      newDelayIndex = min + 1
    } else {
      min = adjacentDelayIndex
      max = state.value.editKey
      newKeyIndex = min + 1
      newDelayIndex = min + 2
    }

    if (max !== false) {
      for (let i = Object.keys(steps.value).length - 1; i >= max; i--) {
        steps.value[i + 2] = steps.value[i]
      }
    }

    recordStep(e, direction, newKeyIndex)
    recordStep(10, false, newDelayIndex)

    state.value.editKey = false
  }

  const deleteEditKey = () => {
    if (state.value.editKey === 0) steps.value.splice(state.value.editKey, 2)
    else steps.value.splice(state.value.editKey - 1, 2)
    state.value.editKey = false
  }

  const restartDelay = () => {
    delay.value.start = performance.now()
  }

  const changeName = (name) => {
    macroName.value = name
  }

  const changeDelay = (fixed) => {
    delay.value.fixed = fixed

    formatDelays()
  }

  const formatDelays = () => {
    steps.value = steps.value.map((step) => {
      if (step.type === 'delay' && delay.value.fixed !== false) step.value = delay.value.fixed
      return step
    })
  }

  const toggleEdit = (type, key) => {
    if (type === 'key') {
      state.value.editKey = key
      state.value.editDelay = false
    }

    if (type === 'delay') {
      state.value.editKey = false
      state.value.editDelay = key
    }
  }

  const resetEdit = () => {
    state.value.edit = false
    state.value.editKey = false
    state.value.editDelay = false
  }

  const reset = () => {
    state.value.record = false
    delay.value.start = 0
    steps.value = []

    if (state.value.edit) resetEdit()
  }

  const check = async () => {
    const resp = await axios.post(appUrl() + '/macro/check', {
      macro: macroName.value,
    })

    return resp.data
  }

  const save = async () => {
    state.value.validationErrors = invalidMacro(steps.value)

    if (state.value.validationErrors) return false

    const resp = await axios.post(appUrl() + '/macro/record', {
      name: macroName.value,
      steps: steps.value,
    })

    return resp.status == 200
  }

  const open = async (macroFileName, name) => {
    const openResp = await axios.post(appUrl() + '/macro/open', {
      macro: macroFileName,
    })

    if (openResp.data) steps.value = translateJSON(openResp.data)

    // console.log(macroName)

    macroName.value = name
  }

  return {
    state,
    macroName,
    steps,
    delay,
    getEditKey,
    getAdjacentKey,
    getEditDelay,
    recordStep,
    insertKey,
    deleteEditKey,
    restartDelay,
    changeName,
    changeDelay,
    toggleEdit,
    resetEdit,
    reset,
    check,
    save,
    open,
  }
})
