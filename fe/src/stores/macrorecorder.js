import { ref } from 'vue'
import { defineStore } from 'pinia'

import { filterKey, isRepeat } from '../services/MacroRecordService'

export const useMacroRecorderStore = defineStore('macrorecorder', () => {
  // Properties - State values
  const state = ref({
    record: false,
    edit: false,
    editKey: false,
    editDelay: false,
  })

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
      returnVal = { delay: steps.value[delayIndex], key: steps.value[keyIndex] }

    return returnVal
  }

  const getEditDelay = () => steps.value[state.value.editDelay]

  // Setters - Actions
  const recordStep = (e, direction = false, key = false) => {
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

  const deleteEditKey = () => {
    steps.value.splice(state.value.editKey, 2)
    state.value.editKey = false
  }

  const restartDelay = () => {
    delay.value.start = performance.now()
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
    steps.value = []

    if (state.value.edit) resetEdit()
  }

  return {
    state,
    steps,
    delay,
    getEditKey,
    getAdjacentKey,
    getEditDelay,
    recordStep,
    deleteEditKey,
    restartDelay,
    changeDelay,
    toggleEdit,
    resetEdit,
    reset,
  }
})
