import { defineStore } from 'pinia'
import { useStorage } from '@vueuse/core'

import { watch } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const apiKey = useStorage('apikey', '')

  function setApiKey(key: string) {
    apiKey.value = key
  }
  watch(apiKey, () => {})

  return {
    apiKey,
    setApiKey,
  }
})
