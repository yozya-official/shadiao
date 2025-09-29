<template>
  <!-- 视频内容网格 -->
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
    <div
      v-for="video in filteredVideos"
      :key="video.id"
      class="group relative bg-card border border-border rounded-2xl shadow-md hover:shadow-2xl transition-all duration-300 overflow-hidden hover:-translate-y-2 cursor-pointer"
    >
      <!-- 封面图片 -->
      <div class="relative overflow-hidden">
        <img
          :src="video.cover"
          referrerpolicy="no-referrer"
          class="w-full h-36 md:h-40 object-cover transition-transform duration-500 group-hover:scale-110"
          :alt="video.title"
        />
        <!-- 渐变遮罩 -->
        <div
          class="absolute inset-0 bg-gradient-to-t from-background/60 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"
        ></div>

        <!-- 删除按钮 -->
        <button
          @click.stop="delVideoConfirm(video)"
          class="absolute z-20 top-3 right-3 w-8 h-8 flex items-center justify-center rounded-full bg-destructive/90 text-destructive-foreground opacity-0 group-hover:opacity-100 hover:bg-destructive hover:scale-110 transition-all duration-200 shadow-lg backdrop-blur-sm"
          title="删除视频"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            ></path>
          </svg>
        </button>

        <!-- 播放图标覆层 -->
        <div
          @click="openVideo(video)"
          class="absolute z-10 inset-0 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity duration-300"
        >
          <div
            class="w-16 h-16 rounded-full bg-primary/90 backdrop-blur-sm flex items-center justify-center shadow-xl transform group-hover:scale-110 transition-transform duration-300"
          >
            <svg
              class="w-8 h-8 text-primary-foreground ml-1"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path
                d="M6.3 2.841A1.5 1.5 0 004 4.11V15.89a1.5 1.5 0 002.3 1.269l9.344-5.89a1.5 1.5 0 000-2.538L6.3 2.84z"
              ></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- 内容区域 -->
      <div class="p-4 space-y-2">
        <!-- 标题 -->
        <h3
          class="font-bold text-foreground line-clamp-1 group-hover:text-primary transition-colors duration-200"
        >
          <a
            :href="video.url"
            target="_blank"
            rel="noopener noreferrer"
            class="hover:underline"
            @click.stop
          >
            {{ video.title }}
          </a>
        </h3>

        <!-- 作者和日期 -->
        <div class="flex items-center justify-between text-sm">
          <div
            v-if="showAuthor"
            @click="router.push({ name: 'author', params: { authorId: video.author.id } })"
            class="flex items-center space-x-2 flex-1 min-w-0"
          >
            <div
              class="w-7 h-7 rounded-full bg-gradient-to-br from-primary/20 to-chart-2/20 p-0.5 flex-shrink-0"
            >
              <img
                referrerpolicy="no-referrer"
                :src="video.author.avatar"
                :alt="video.author.name"
                class="w-full h-full rounded-full object-cover"
              />
            </div>
            <span class="text-muted-foreground truncate font-medium">{{ video.author.name }}</span>
          </div>
          <time class="text-xs text-muted-foreground/70 whitespace-nowrap ml-2">
            {{ formatDate(video.ctime || new Date().toISOString()) }}
          </time>
        </div>

        <!-- 主要标签 -->
        <div class="flex flex-wrap gap-2">
          <span
            v-if="video.isOriginal"
            class="inline-flex items-center px-2.5 py-1 bg-primary/10 text-primary text-xs rounded-md font-medium border border-primary/20"
          >
            <svg class="w-3 h-3 mr-1" fill="currentColor" viewBox="0 0 20 20">
              <path
                d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
              ></path>
            </svg>
            原创
          </span>
          <span
            v-if="video.isCompleted"
            class="inline-flex items-center px-2.5 py-1 bg-chart-2/10 text-chart-2 text-xs rounded-md font-medium border border-chart-2/20"
          >
            <svg class="w-3 h-3 mr-1" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                clip-rule="evenodd"
              ></path>
            </svg>
            完结
          </span>
          <span
            v-if="video.hasSystem"
            class="inline-flex items-center px-2.5 py-1 bg-chart-3/10 text-chart-3 text-xs rounded-md font-medium border border-chart-3/20"
          >
            <svg class="w-3 h-3 mr-1" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z"
                clip-rule="evenodd"
              ></path>
            </svg>
            系统
          </span>
          <span
            v-if="video.background"
            class="inline-flex items-center px-2.5 py-1 bg-accent/60 text-accent-foreground text-xs rounded-md font-medium"
          >
            {{ backgroundOptions[video.background] }} {{ video.background }}
          </span>
        </div>

        <!-- 风格标签 -->
        <div v-if="video.style.length > 0" class="flex flex-wrap gap-1.5">
          <span
            v-for="s in video.style.slice(0, 3)"
            :key="s"
            class="px-2 py-0.5 bg-muted/40 text-muted-foreground text-xs rounded border border-border/50"
          >
            {{ styleOptions[s] }} {{ s }}
          </span>
          <span
            v-if="video.style.length > 3"
            class="px-2 py-0.5 bg-muted/40 text-muted-foreground text-xs rounded border border-border/50"
          >
            +{{ video.style.length - 3 }}
          </span>
        </div>
      </div>
    </div>
  </div>

  <!-- 分页控制 -->
  <div v-if="props.videos.length > pageSize" class="flex justify-center mt-12">
    <div class="flex items-center space-x-2 bg-card border border-border rounded-xl p-2 shadow-sm">
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
          ></path>
        </svg>
      </button>

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

      <button
        @click="currentPage = Math.min(totalPages, currentPage + 1)"
        :disabled="currentPage === totalPages"
        class="px-3 py-2 rounded-lg text-muted-foreground hover:text-foreground hover:bg-accent disabled:opacity-30 disabled:cursor-not-allowed transition-all duration-200"
        :class="{ 'hover:bg-transparent': currentPage === totalPages }"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M9 5l7 7-7 7"
          ></path>
        </svg>
      </button>
    </div>
  </div>

  <!-- 视频详情模态框 -->
  <Transition
    enter-active-class="transition-all duration-300 ease-out"
    enter-from-class="opacity-0 scale-95"
    enter-to-class="opacity-100 scale-100"
    leave-active-class="transition-all duration-200 ease-in"
    leave-from-class="opacity-100 scale-100"
    leave-to-class="opacity-0 scale-95"
  >
    <div
      v-if="selectedVideo && showVideoModal"
      class="fixed inset-0 z-50 flex items-center justify-center p-4"
      @click="closeVideoModal"
    >
      <div class="fixed inset-0 bg-background/80 backdrop-blur-sm"></div>
      <div
        class="relative bg-card border border-border rounded-3xl shadow-2xl w-full max-w-3xl max-h-[90vh] overflow-y-auto"
        @click.stop
      >
        <!-- 关闭按钮 -->
        <button
          @click="closeVideoModal"
          class="absolute top-4 right-4 z-10 w-10 h-10 flex items-center justify-center rounded-full bg-muted/80 hover:bg-muted text-muted-foreground hover:text-foreground transition-all duration-200"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            ></path>
          </svg>
        </button>

        <div class="p-8 space-y-6">
          <!-- 标题 -->
          <h3 class="text-3xl font-bold text-foreground pr-12">{{ selectedVideo.title }}</h3>

          <!-- 作者信息 -->
          <div
            v-if="showAuthor"
            class="flex cursor-pointer items-center space-x-3 p-4 bg-muted/20 rounded-xl"
            @click="goToAuthor(selectedVideo.author.id)"
          >
            <div class="w-12 h-12 rounded-full bg-gradient-to-br from-primary/20 to-chart-2/20 p-1">
              <img
                :src="selectedVideo.author.avatar"
                :alt="selectedVideo.author.name"
                referrerpolicy="no-referrer"
                class="w-full h-full rounded-full object-cover"
              />
            </div>
            <div>
              <div class="text-foreground font-semibold">{{ selectedVideo.author.name }}</div>
              <div class="text-sm text-muted-foreground">创作者</div>
            </div>
          </div>

          <!-- 推荐理由 -->
          <div
            v-if="selectedVideo.description"
            class="p-4 bg-primary/5 border border-primary/10 rounded-xl"
          >
            <h4 class="font-semibold text-foreground mb-2 flex items-center">
              <svg class="w-5 h-5 text-primary mr-2" fill="currentColor" viewBox="0 0 20 20">
                <path
                  fill-rule="evenodd"
                  d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                  clip-rule="evenodd"
                ></path>
              </svg>
              推荐理由
            </h4>
            <p class="text-foreground/80 leading-relaxed">{{ selectedVideo.description }}</p>
          </div>

          <!-- 详细标签 -->
          <div class="space-y-3">
            <h4 class="font-semibold text-foreground flex items-center">
              <svg class="w-5 h-5 text-chart-2 mr-2" fill="currentColor" viewBox="0 0 20 20">
                <path
                  fill-rule="evenodd"
                  d="M17.707 9.293a1 1 0 010 1.414l-7 7a1 1 0 01-1.414 0l-7-7A.997.997 0 012 10V5a3 3 0 013-3h5c.256 0 .512.098.707.293l7 7zM5 6a1 1 0 100-2 1 1 0 000 2z"
                  clip-rule="evenodd"
                ></path>
              </svg>
              标签详情
            </h4>
            <div class="flex flex-wrap gap-2">
              <span
                v-if="selectedVideo.isOriginal"
                class="inline-flex items-center px-3 py-2 bg-primary/10 text-primary rounded-lg font-medium border border-primary/20"
              >
                <svg class="w-4 h-4 mr-1.5" fill="currentColor" viewBox="0 0 20 20">
                  <path
                    d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
                  ></path>
                </svg>
                原创作品
              </span>
              <span
                v-if="selectedVideo.isCompleted"
                class="inline-flex items-center px-3 py-2 bg-chart-2/10 text-chart-2 rounded-lg font-medium border border-chart-2/20"
              >
                <svg class="w-4 h-4 mr-1.5" fill="currentColor" viewBox="0 0 20 20">
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd"
                  ></path>
                </svg>
                已完结
              </span>
              <span
                v-if="selectedVideo.background"
                class="inline-flex items-center px-3 py-2 bg-accent/60 text-accent-foreground rounded-lg font-medium"
              >
                {{ backgroundOptions[selectedVideo.background] }} {{ selectedVideo.background }}
              </span>
              <span
                v-for="s in selectedVideo.style"
                :key="s"
                class="inline-flex items-center px-3 py-2 bg-muted/60 text-muted-foreground rounded-lg font-medium border border-border"
              >
                {{ styleOptions[s] }} {{ s }}
              </span>
              <span
                v-if="selectedVideo.hasSystem"
                class="inline-flex items-center px-3 py-2 bg-chart-3/10 text-chart-3 rounded-lg font-medium border border-chart-3/20"
              >
                <svg class="w-4 h-4 mr-1.5" fill="currentColor" viewBox="0 0 20 20">
                  <path
                    fill-rule="evenodd"
                    d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z"
                    clip-rule="evenodd"
                  ></path>
                </svg>
                系统外挂
              </span>
            </div>
          </div>

          <!-- 操作按钮 -->
          <div class="flex items-center space-x-4 pt-4">
            <a
              :href="selectedVideo.url"
              target="_blank"
              rel="noopener noreferrer"
              class="flex-1 px-6 py-3 bg-gradient-to-r from-primary to-chart-2 text-primary-foreground rounded-xl hover:shadow-lg transition-all duration-200 font-semibold text-center flex items-center justify-center space-x-2"
            >
              <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                <path d="M10 12a2 2 0 100-4 2 2 0 000 4z"></path>
                <path
                  fill-rule="evenodd"
                  d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z"
                  clip-rule="evenodd"
                ></path>
              </svg>
              <span>立即观看</span>
            </a>
            <button
              @click="closeVideoModal"
              class="px-6 py-3 bg-muted/40 hover:bg-muted/60 text-muted-foreground hover:text-foreground rounded-xl transition-all duration-200 font-medium"
            >
              关闭
            </button>
          </div>
        </div>
      </div>
    </div>
  </Transition>

  <!-- 删除确认模态框 -->
  <Transition
    enter-active-class="transition-all duration-300 ease-out"
    enter-from-class="opacity-0 scale-95"
    enter-to-class="opacity-100 scale-100"
    leave-active-class="transition-all duration-200 ease-in"
    leave-from-class="opacity-100 scale-100"
    leave-to-class="opacity-0 scale-95"
  >
    <div
      v-if="showDeleteModal"
      class="fixed inset-0 z-50 flex items-center justify-center p-4"
      @click="closeDeleteModal"
    >
      <div class="fixed inset-0 bg-background/80 backdrop-blur-sm"></div>
      <div
        class="relative bg-card border border-border rounded-2xl shadow-2xl p-6 w-full max-w-md"
        @click.stop
      >
        <div class="text-center mb-6">
          <div
            class="w-16 h-16 mx-auto mb-4 rounded-full bg-destructive/10 flex items-center justify-center"
          >
            <svg
              class="w-8 h-8 text-destructive"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-1.964-1.333-2.732 0L3.34 16c-.77 1.333.192 3 1.732 3z"
              ></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-foreground mb-2">确认删除视频</h3>
          <p class="text-sm text-destructive">此操作不可撤销，将永久删除该视频</p>
        </div>

        <div class="flex space-x-3">
          <button
            @click="closeDeleteModal"
            class="flex-1 px-4 py-2 border border-border text-foreground rounded-lg hover:bg-accent transition-colors duration-200"
          >
            取消
          </button>
          <button
            @click="delVideo"
            class="flex-1 px-4 py-2 bg-destructive text-destructive-foreground rounded-lg hover:bg-destructive/90 transition-colors duration-200"
          >
            确认删除
          </button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { videoApi } from '@/api'
import type { VideoData } from '@/models/video'
import { backgroundOptions, styleOptions } from '@/stores/options'
import { formatDate, handleApiError } from '@/utils'
import { toast } from '@yuelioi/toast'
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
const router = useRouter()

import { useVideoStore } from '@/stores/videoStore'
const videoStore = useVideoStore()

const selectedVideo = ref<VideoData | null>(null)
const showVideoModal = ref(false)
const showDeleteModal = ref(false)

const currentPage = ref(1)
const pageSize = ref(12)

const props = defineProps<{
  videos: VideoData[]
  showAuthor?: boolean
}>()

const totalPages = computed(() => Math.ceil(props.videos.length / pageSize.value))

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, start + 4)

  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  return pages
})

const goToAuthor = (id?: number) => {
  if (!id) return
  router.push({ name: 'author', params: { authorId: id } })
}

const openVideo = (video: VideoData) => {
  selectedVideo.value = video
  showVideoModal.value = true
  document.body.style.overflow = 'hidden'
}

const closeVideoModal = () => {
  showVideoModal.value = false
  document.body.style.overflow = ''
}

const delVideoConfirm = (video: VideoData) => {
  console.log(111)
  selectedVideo.value = video
  showDeleteModal.value = true
  document.body.style.overflow = 'hidden'
}

const closeDeleteModal = () => {
  showDeleteModal.value = false
  document.body.style.overflow = ''
}

const delVideo = async () => {
  try {
    const id = selectedVideo.value?.id
    if (!id) throw 'id 不能为空'
    await videoApi.deleteVideo(id)
    toast.success('删除成功')
    closeDeleteModal()
    await videoStore.loadVideos()
  } catch (error: unknown) {
    handleApiError(error, '删除')
  }
}

const filteredVideos = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return props.videos.slice(start, end)
})
</script>

<style scoped>
/* 卡片悬停效果 */
.group:hover {
  transform: translateY(-8px);
}

/* 图片缩放效果 */
.group:hover img {
  transform: scale(1.1);
}

/* 模态框滚动条样式 */
.max-h-\[90vh\]::-webkit-scrollbar {
  width: 8px;
}

.max-h-\[90vh\]::-webkit-scrollbar-track {
  background: var(--muted);
  border-radius: 4px;
}

.max-h-\[90vh\]::-webkit-scrollbar-thumb {
  background: var(--primary);
  border-radius: 4px;
}

.max-h-\[90vh\]::-webkit-scrollbar-thumb:hover {
  background: var(--primary);
  opacity: 0.8;
}
</style>
