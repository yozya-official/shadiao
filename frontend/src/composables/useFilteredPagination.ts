import { ref, computed, watch } from 'vue'

interface UseFilteredPaginationOptions<T> {
  data: Ref<T[]> | T[]
  pageSize?: number
  filterFn?: (item: T, searchQuery: string) => boolean
  sortFn?: (a: T, b: T) => number
}

export function useFilteredPagination<T>({
  data,
  pageSize = 12,
  filterFn,
  sortFn,
}: UseFilteredPaginationOptions<T>) {
  const searchQuery = ref('')
  const currentPage = ref(1)

  // 筛选后的数据
  const filteredData = computed(() => {
    const dataArray = unref(data)

    if (!searchQuery.value || !filterFn) {
      return dataArray
    }

    return dataArray.filter((item) => filterFn(item, searchQuery.value))
  })

  // 排序后的数据
  const sortedData = computed(() => {
    if (!sortFn) {
      return filteredData.value
    }
    return [...filteredData.value].sort(sortFn)
  })

  // 总页数
  const totalPages = computed(() => Math.ceil(sortedData.value.length / pageSize))

  // 当前页数据
  const pagedData = computed(() => {
    const start = (currentPage.value - 1) * pageSize
    const end = start + pageSize
    return sortedData.value.slice(start, end)
  })

  // 重置搜索时回到第一页
  watch(searchQuery, () => {
    currentPage.value = 1
  })

  // 清空搜索
  const clearSearch = () => {
    searchQuery.value = ''
  }

  return {
    searchQuery,
    currentPage,
    pagedData,
    filteredData,
    totalPages,
    total: computed(() => sortedData.value.length),
    clearSearch,
  }
}
