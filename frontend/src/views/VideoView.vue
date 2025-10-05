<template>
  <PageHeader title="沙雕动画收集站" subtitle="发现更多精彩作品，探索无限创意世界">
    <template #icon>
      <svg class="size-8" fill="white" viewBox="0 0 20 20">
        <path
          d="M2 6a2 2 0 012-2h6a2 2 0 012 2v6a2 2 0 01-2 2H4a2 2 0 01-2-2V6zM14.553 7.106A1 1 0 0014 8v4a1 1 0 00.553.894l2 1A1 1 0 0018 13V7a1 1 0 00-1.447-.894l-2 1z"
        ></path>
      </svg>
    </template>
  </PageHeader>
  <div class="container px-6 py-8 relative">
    <!-- 筛选器区域 -->
    <div
      v-if="showFilter"
      class="bg-card/90 backdrop-blur-sm border rounded-3xl shadow p-6 mb-6 hover:shadow-2xl transition-all duration-300"
    >
      <div class="flex items-center justify-between mb-8">
        <div class="flex items-center">
          <div class="size-4 bg-gradient-to-br from-primary to-chart-2 rounded-full mr-3"></div>
          <h3 class="">筛选器</h3>
        </div>

        <div class="flex gap-2">
          <button
            @click="showFilter = false"
            class="btn bg-accent/20 border hover:bg-accent/40 text-accent-foreground rounded-xl transition-all duration-200 font-medium flex items-center space-x-1 group"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="size-4"
              width="24"
              height="24"
              viewBox="0 0 24 24"
            >
              <path
                fill="none"
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="m15 18l-.722-3.25M2 8a10.645 10.645 0 0 0 20 0m-2 7l-1.726-2.05M4 15l1.726-2.05M9 18l.722-3.25"
              />
            </svg>
            <span class="hidden md:inline">隐藏过滤器</span>
          </button>
          <button
            @click="resetFilters"
            class="btn bg-accent/20 border hover:bg-accent/40 text-accent-foreground rounded-xl transition-all duration-200 font-medium flex items-center space-x-1 group"
          >
            <svg
              class="size-4 group-hover:rotate-180 transition-transform duration-300"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"
              ></path>
            </svg>
            <span class="hidden md:inline">重置筛选</span>
          </button>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <!-- 搜索框 -->
        <div class="space-y-2">
          <label class="flex items-center space-x-2 font-semibold">
            <svg class="size-4 text-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
              ></path>
            </svg>
            <span>搜索</span>
          </label>
          <input
            v-model="filters.search"
            type="text"
            placeholder="搜索作品名称..."
            class="w-full input input-primary"
          />
        </div>

        <!-- 背景设定筛选 -->
        <div class="space-y-2">
          <label for="background" class="flex items-center space-x-2 font-semibold">
            <svg class="size-4 text-chart-2" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zM4.332 8.027a6.012 6.012 0 011.912-2.706C6.512 5.73 6.974 6 7.5 6A1.5 1.5 0 019 7.5V8a2 2 0 004 0 2 2 0 011.523-1.943A5.977 5.977 0 0116 10c0 .34-.028.675-.083 1H15a2 2 0 00-2 2v2.197A5.973 5.973 0 0110 16v-2a2 2 0 00-2-2 2 2 0 01-2-2 2 2 0 00-1.668-1.973z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <span>背景设定</span>
          </label>
          <select id="background" v-model="filters.background" class="w-full select input-primary">
            <option value="">全部背景</option>
            <option
              v-for="(icon, background) in backgroundOptions"
              :key="background"
              :value="background"
            >
              {{ icon }} {{ background }}
            </option>
          </select>
        </div>

        <!-- 世界筛选 -->
        <div class="space-y-2">
          <label for="world" class="flex items-center space-x-2 font-semibold">
            <svg class="size-4 text-chart-3" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zM4.332 8.027a6.012 6.012 0 011.912-2.706C6.512 5.73 6.974 6 7.5 6A1.5 1.5 0 019 7.5V8a2 2 0 004 0 2 2 0 011.523-1.943A5.977 5.977 0 0116 10c0 .34-.028.675-.083 1H15a2 2 0 00-2 2v2.197A5.973 5.973 0 0110 16v-2a2 2 0 00-2-2 2 2 0 01-2-2 2 2 0 00-1.668-1.973z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <span>世界设定</span>
          </label>
          <select id="world" v-model="filters.world" class="w-full select select-primary">
            <option value="">全部世界</option>
            <option v-for="(icon, world) in worldOptions" :key="world" :value="world">
              {{ icon }} {{ world }}
            </option>
          </select>
        </div>

        <!-- 风格筛选 -->
        <div class="space-y-2">
          <label for="style" class="flex items-center space-x-2 font-semibold">
            <svg class="size-4 text-chart-4" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M7 4a3 3 0 016 0v4a3 3 0 11-6 0V4zm4 10.93A7.001 7.001 0 0017 8a1 1 0 10-2 0A5 5 0 015 8a1 1 0 00-2 0 7.001 7.001 0 006 6.93V17H6a1 1 0 100 2h8a1 1 0 100-2h-3v-2.07z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <span>风格</span>
          </label>
          <select id="style" v-model="filters.style" class="w-full select select-primary">
            <option value="">全部风格</option>
            <option v-for="(icon, style) in styleOptions" :key="style" :value="style">
              {{ icon }} {{ style }}
            </option>
          </select>
        </div>
      </div>

      <!-- 状态筛选 -->
      <div class="space-y-3">
        <label class="flex items-center space-x-2 font-semibold">
          <svg class="size-4 text-chart-5" fill="currentColor" viewBox="0 0 20 20">
            <path
              d="M3 4a1 1 0 011-1h12a1 1 0 011 1v2a1 1 0 01-1 1H4a1 1 0 01-1-1V4zM3 10a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H4a1 1 0 01-1-1v-6zM14 9a1 1 0 00-1 1v6a1 1 0 001 1h2a1 1 0 001-1v-6a1 1 0 00-1-1h-2z"
            ></path>
          </svg>
          <span>状态筛选</span>
        </label>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <label
            class="group cursor-pointer bg-muted/20 hover:bg-muted/40 rounded-xl p-3 border transition-all duration-200 flex items-center justify-between"
          >
            <div class="flex items-center space-x-3">
              <div class="size-6 rounded-lg bg-primary/10 flex items-center justify-center">
                <svg class="size-4 text-primary" fill="currentColor" viewBox="0 0 20 20">
                  <path
                    d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
                  ></path>
                </svg>
              </div>
              <span class="font-medium">原创作品</span>
            </div>
            <div class="relative">
              <input
                v-model="filters.isOriginal"
                type="checkbox"
                :class="{ 'no-checked': filters.isOriginal === false }"
                class="size-4 text-primary border-2 rounded focus:ring-2 focus:ring-primary/50 transition-all duration-200"
              />
            </div>
          </label>

          <label
            class="group cursor-pointer bg-muted/20 hover:bg-muted/40 rounded-xl p-3 border transition-all duration-200 flex items-center justify-between"
          >
            <div class="flex items-center space-x-3">
              <div class="size-6 rounded-lg bg-chart-2/10 flex items-center justify-center">
                <svg class="size-4 text-chart-2" fill="currentColor" viewBox="0 0 20 20">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd"
                  ></path>
                </svg>
              </div>
              <span class="font-medium">已完结</span>
            </div>
            <div class="relative">
              <input
                v-model="filters.isCompleted"
                type="checkbox"
                :class="{ 'no-checked': filters.isCompleted === false }"
                class="size-4 text-chart-2 border-2 rounded focus:ring-2 focus:ring-chart-2/50 transition-all duration-200"
              />
            </div>
          </label>

          <label
            class="group cursor-pointer bg-muted/20 hover:bg-muted/40 rounded-xl p-3 border transition-all duration-200 flex items-center justify-between"
          >
            <div class="flex items-center space-x-3">
              <div class="size-6 rounded-lg bg-chart-3/10 flex items-center justify-center">
                <svg class="size-4 text-chart-3" fill="currentColor" viewBox="0 0 20 20">
                  <path
                    fill-rule="evenodd"
                    d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z"
                    clip-rule="evenodd"
                  ></path>
                </svg>
              </div>
              <span class="font-medium">有系统</span>
            </div>
            <div class="relative">
              <input
                v-model="filters.hasSystem"
                type="checkbox"
                :class="{ 'no-checked': filters.hasSystem === false }"
                class="size-4 text-chart-3 border-2 rounded focus:ring-2 focus:ring-chart-3/50 transition-all duration-200"
              />
            </div>
          </label>
        </div>
      </div>
    </div>

    <!-- 结果统计和排序 -->
    <div
      class="flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8 bg-card/60 backdrop-blur-sm border rounded-2xl p-6"
    >
      <div class="text-muted-foreground text-lg pb-2">
        找到
        <span class="font-bold text-primary text-2xl">{{ filteredVideos.length }}</span>
        部精彩作品
      </div>

      <div class="flex items-center space-x-3">
        <select v-model="sortBy" class="px-4 select pr-8 py-2 input-primary" aria-label="排序方式">
          <option value="name">按名称排序</option>
          <option value="author">按作者排序</option>
          <option value="ctime">按创建时间</option>
          <option value="created">按投稿时间</option>
        </select>
        <button
          @click="sortOrder = sortOrder === 'asc' ? 'desc' : 'asc'"
          class="p-2 bg-accent/20 hover:bg-accent/40 border rounded-xl transition-all duration-200 text-accent-foreground"
          :title="sortOrder === 'asc' ? '升序' : '降序'"
        >
          <svg
            class="size-4 transition-transform duration-200"
            :class="{ 'rotate-180': sortOrder === 'desc' }"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M3 4h13M3 8h9m-9 4h6m4 0l4-4m0 0l4 4m-4-4v12"
            ></path>
          </svg>
        </button>
        <button
          v-if="!showFilter"
          @click="showFilter = true"
          class="btn ml-auto bg-accent/20 border hover:bg-accent/40 text-accent-foreground rounded-xl transition-all duration-200 font-medium flex items-center space-x-2 group"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="size-4"
            width="24"
            height="24"
            viewBox="0 0 24 24"
          >
            <g
              fill="none"
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
            >
              <path
                d="M2.062 12.348a1 1 0 0 1 0-.696a10.75 10.75 0 0 1 19.876 0a1 1 0 0 1 0 .696a10.75 10.75 0 0 1-19.876 0"
              />
              <circle cx="12" cy="12" r="3" />
            </g>
          </svg>
          <span>显示过滤器</span>
        </button>
      </div>
    </div>

    <!-- 视频结果 -->
    <div v-if="filteredVideos.length" class="relative">
      <VideoContainer
        v-model:videos="pagedVideos"
        :show-author="true"
        ref="videoContainerRef"
      ></VideoContainer>

      <DataPagination
        v-model:total="filteredVideos.length"
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
      ></DataPagination>

      <!-- 装饰性渐变叠加 -->
      <div class="absolute inset-0 pointer-events-none">
        <div class="absolute top-0 left-0 w-32 h-32 bg-primary/5 rounded-full blur-2xl"></div>
        <div class="absolute bottom-20 right-20 w-40 h-40 bg-chart-2/5 rounded-full blur-3xl"></div>
      </div>
    </div>

    <!-- 空状态 -->
    <div v-else class="text-center py-20">
      <div class="w-32 h-32 mx-auto mb-8 rounded-full bg-muted/20 flex items-center justify-center">
        <svg
          class="w-16 h-16 text-muted-foreground"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="1.5"
            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
          ></path>
        </svg>
      </div>
      <h3 class="text-2xl font-bold mb-4">没有找到匹配的作品</h3>
      <p class="text-muted-foreground mb-8 max-w-md mx-auto">
        尝试调整筛选条件或者重置来发现更多内容
      </p>
      <button
        @click="resetFilters"
        class="px-8 py-4 bg-primary text-primary-foreground rounded-xl hover:bg-primary/90 transition-all duration-200 font-semibold shadow-sm hover:shadow-xl"
      >
        <svg class="size-4 inline mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"
          ></path>
        </svg>
        重置筛选器
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import VideoContainer from '@/components/VideoContainer.vue'
import { backgroundOptions, styleOptions, worldOptions } from '@/stores/options'

const videoStore = useVideoStore()
const filterStore = useFilterStore()
const { videos } = storeToRefs(videoStore)
const { filters, showFilter, sortBy, sortOrder, currentPage } = storeToRefs(filterStore)
const { resetFilters } = filterStore
const videoContainerRef = ref<InstanceType<typeof VideoContainer> | null>(null)

const pageSize = ref(12)

const sortedVideos = computed(() => {
  const result = [...filteredVideos.value]

  result.sort((a, b) => {
    let aValue: string | Date | number
    let bValue: string | Date | number

    switch (sortBy.value) {
      case 'name':
        aValue = a.title
        bValue = b.title
        break
      case 'author':
        aValue = a.author?.uid || 0
        bValue = b.author?.uid || 0
        break
      case 'ctime':
        aValue = a.ctime || new Date()
        bValue = b.ctime || new Date()
      case 'created':
        aValue = a.createdAt || new Date()
        bValue = b.createdAt || new Date()
        break
      default:
        return 0
    }

    if (sortOrder.value === 'asc') {
      return aValue > bValue ? 1 : -1
    } else {
      return aValue < bValue ? 1 : -1
    }
  })

  return result
})

const filteredVideos = ref<Video[]>([])

watch(
  () => ({
    ...filters.value,
    sortBy: sortBy.value,
    sortOrder: sortOrder.value,
    videos: videos.value,
  }),
  () => {
    currentPage.value = 1
    let result = [...videos.value]

    // 搜索过滤
    if (filters.value.search) {
      const keyword = filters.value.search.toLowerCase()
      result = result.filter((v) => {
        return (
          v.title.toLowerCase().includes(keyword) || v.author?.name.toLowerCase().includes(keyword)
        )
      })
    }

    // 背景/世界/风格过滤
    if (filters.value.background) {
      result = result.filter((v) => v.background()?.name === filters.value.background)
    }
    if (filters.value.world) {
      result = result.filter((v) => v.world()?.name === filters.value.world)
    }
    if (filters.value.style) {
      result = result.filter((v) => v.style().some((tag) => tag.name === filters.value.style))
    }

    // 布尔标签过滤
    const booleanFilters: (keyof Filters)[] = ['isOriginal', 'isCompleted', 'hasSystem']
    for (const key of booleanFilters) {
      const filterVal = filters.value[key]
      if (filterVal !== null && filterVal !== '') {
        result = result.filter((v: Video) => {
          // @ts-expect-error tag 布尔筛选器
          const val = v[key]() as boolean | null
          return val === filterVal
        })
      }
    }

    filteredVideos.value = result
  },
  { deep: true, immediate: true },
)

const pagedVideos = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = currentPage.value * pageSize.value
  return sortedVideos.value.slice(start, end)
})

onMounted(async () => {
  await videoStore.loadVideos()
})
</script>

<style scoped>
/* 过滤器不为null 且 为false时 */
.no-checked {
  position: relative;
}

.no-checked::before {
  content: '✖';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  color: var(--destructive);
  font-size: 12px;
  font-weight: bold;
  z-index: 10;
}

.no-checked::after {
  content: '';
  position: absolute;
  inset: 0;
  background: var(--background);
  border: 2px solid var(--destructive);
  border-radius: 4px;
}

/* 状态筛选卡片悬停效果 */
.group:hover {
  transform: translateY(-1px);
}
</style>
