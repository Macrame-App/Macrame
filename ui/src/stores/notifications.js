import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useNoticationStore = defineStore('notications', () => {
  const list = ref({})

  const add = (notification) => {
    list.value[Date.now()] = notification
  }

  const remove = (id) => {
    delete list.value[id]
  }

  return {
    list,
    add,
    remove,
  }
})
