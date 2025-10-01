<template>
  <div v-if="total > pageSize" class="flex justify-center mt-12">
    <div class="flex items-center space-x-2 bg-card border border-border rounded-xl p-2 shadow-sm">
      <!-- 上一页 -->
      <button
        @click="currentPage = Math.max(1, currentPage - 1)"
        :disabled="currentPage === 1"
        class="px-3 py-2 rounded-lg text-muted-foreground hover:text-foreground hover:bg-accent disabled:opacity-30 disabled:cursor-not-allowed transition-all duration-200"
        :class="{ 'hover:bg-transparent': currentPage === 1 }"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M15 19l-7-7 7-7"
          />
        </svg>
      </button>

      <!-- 页码按钮 -->
      <div class="flex items-center space-x-1">
        <button
          v-for="page in visiblePages"
          :key="page"
          @click="currentPage = page"
          class="min-w-[2.5rem] px-3 py-2 rounded-lg text-sm font-medium transition-all duration-200"
          :class="
            currentPage === page
              ? 'bg-primary text-primary-foreground shadow-sm'
              : 'text-muted-foreground hover:text-foreground hover:bg-accent'
          "
        >
          {{ page }}
        </button>
      </div>

      <!-- 下一页 -->
      <button
        @click="currentPage = Math.min(totalPages, currentPage + 1)"
        :disabled="currentPage === totalPages"
        class="px-3 py-2 rounded-lg text-muted-foreground hover:text-foreground hover:bg-accent disabled:opacity-30 disabled:cursor-not-allowed transition-all duration-200"
        :class="{ 'hover:bg-transparent': currentPage === totalPages }"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const currentPage = defineModel<number>('currentPage', { default: 1 })

const pageSize = defineModel<number>('pageSize', { default: 12 })

const total = defineModel<number>('total', { required: true })

const totalPages = computed(() => Math.ceil(total.value / pageSize.value))

const visiblePages = computed(() => {
  const pages: number[] = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, start + 4)

  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  return pages
})

const resetPage = () => {
  currentPage.value = 1
}
const setPage = (page: number) => {
  currentPage.value = page
}

defineExpose({
  resetPage,
  setPage,
})
</script>
