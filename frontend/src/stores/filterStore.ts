import { defineStore } from 'pinia'
import type { Filters } from '@/models'

import { useStorage } from '@vueuse/core'

export const useFilterStore = defineStore('filter', () => {
  const filters = useStorage<Filters>('filters', {
    search: '',
    background: '',
    style: '',
    world: '',
    isOriginal: null,
    isCompleted: null,
    hasSystem: null,
  })

  const sortBy = useStorage<'name' | 'author' | 'created' | 'ctime'>('filter-name', 'ctime')
  const sortOrder = useStorage<'asc' | 'desc'>('filter-order', 'desc')
  const showFilter = useStorage<boolean>('filter-showFilter', true)
  const currentPage = useStorage<number>('filter-currentPage', 1)

  const toggleFilterVisibility = () => {
    showFilter.value = !showFilter.value
  }

  const toggleSortOrder = () => {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'desc'
  }

  const resetFilters = () => {
    Object.assign(filters.value, {
      search: '',
      background: '',
      style: '',
      world: '',
      isOriginal: null,
      isCompleted: null,
      hasSystem: null,
    })
  }

  return {
    filters,
    sortBy,
    sortOrder,
    showFilter,
    currentPage,

    resetFilters,
    toggleSortOrder,
    toggleFilterVisibility,
  }
})
