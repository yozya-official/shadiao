<template>
  <!-- 页面标题 -->
  <PageHeader title="内容审核中心" subtitle="审核待发布的优质内容">
    <template #icon>
      <svg class="w-8 h-8" fill="white" viewBox="0 0 20 20">
        <path
          fill-rule="evenodd"
          d="M3 4a1 1 0 011-1h4a1 1 0 010 2H6.414l2.293 2.293a1 1 0 11-1.414 1.414L5 6.414V8a1 1 0 01-2 0V4zm9 1a1 1 0 010-2h4a1 1 0 011 1v4a1 1 0 01-2 0V6.414l-2.293 2.293a1 1 0 11-1.414-1.414L13.586 5H12zm-9 7a1 1 0 012 0v1.586l2.293-2.293a1 1 0 111.414 1.414L6.414 15H8a1 1 0 010 2H4a1 1 0 01-1-1v-4zm13-1a1 1 0 011 1v4a1 1 0 01-1 1h-4a1 1 0 010-2h1.586l-2.293-2.293a1 1 0 111.414-1.414L15 13.586V12a1 1 0 011-1z"
          clip-rule="evenodd"
        ></path>
      </svg>
    </template>
  </PageHeader>

  <div class="container px-6 py-8 relative">
    <!-- 空状态 -->
    <div v-if="!videos.length" class="text-center py-20">
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
            d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
          ></path>
        </svg>
      </div>
      <h3 class="text-3xl font-bold mb-4">审核列表空空如也</h3>
      <p class="text-muted-foreground text-lg max-w-md mx-auto">暂时没有待审核的视频内容</p>
    </div>

    <!-- 审核列表 -->
    <div
      v-else
      class="bg-card/90 backdrop-blur-sm border border-border rounded-3xl shadow-xl overflow-hidden"
    >
      <!-- 列表头部 -->
      <div class="bg-primary/5 border-b border-border px-8 py-4">
        <div class="flex items-center space-x-3">
          <div class="w-3 h-3 bg-gradient-to-br from-primary to-chart-2 rounded-full"></div>
          <h2 class="text-xl font-bold">待审核内容</h2>
          <span class="px-3 py-1 bg-primary/10 text-primary text-sm rounded-full font-medium">
            {{ videos.length }} 项
          </span>
        </div>
      </div>

      <!-- 审核项目列表 -->
      <div class="divide-y divide-border">
        <div
          v-for="(video, index) in videos"
          :key="video.aid"
          class="group p-4 md:p-8 hover:bg-accent/20 transition-all duration-200"
        >
          <div class="flex flex-col md:flex-row items-start gap-4 md:gap-6">
            <!-- 移动端顶部行：序号 + 头像 + 操作按钮 -->
            <div class="flex items-center justify-between w-full md:w-auto md:contents">
              <!-- 序号 -->
              <div
                class="flex-shrink-0 w-10 h-10 md:w-12 md:h-12 rounded-xl bg-muted/30 flex items-center justify-center text-base md:text-lg font-bold text-muted-foreground group-hover:bg-primary/10 group-hover:text-primary transition-all duration-200"
              >
                {{ index + 1 }}
              </div>

              <!-- 作者头像 -->
              <div class="flex-shrink-0 md:order-none mr-auto">
                <div
                  class="w-12 h-12 md:w-16 md:h-16 rounded-2xl bg-gradient-to-br from-primary/20 to-chart-2/20 p-1 group-hover:from-primary/40 group-hover:to-chart-2/40 transition-all duration-300"
                >
                  <img
                    :src="video.author?.avatar || 'https://via.placeholder.com/64'"
                    :alt="video.author?.name"
                    referrerpolicy="no-referrer"
                    class="w-full h-full rounded-xl object-cover"
                  />
                </div>
              </div>

              <!-- 移动端操作按钮 -->
              <div class="flex md:hidden items-center space-x-2">
                <button
                  @click="approveVideo(video)"
                  class="p-2 bg-chart-2/10 hover:bg-chart-2/20 border border-chart-2/30 text-chart-2 rounded-lg transition-all duration-200"
                  title="通过审核"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M5 13l4 4L19 7"
                    ></path>
                  </svg>
                </button>

                <button
                  @click="rejectVideo(video)"
                  class="p-2 bg-destructive/10 hover:bg-destructive/20 border border-destructive/30 text-destructive rounded-lg transition-all duration-200"
                  title="驳回"
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

                <button
                  @click="goToVideo(video)"
                  class="p-2 bg-primary/10 hover:bg-primary/20 border border-primary/30 text-primary rounded-lg transition-all duration-200"
                  title="查看原视频"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
                    ></path>
                  </svg>
                </button>
              </div>
            </div>

            <!-- 视频信息 -->
            <div class="flex-1 space-y-3 md:space-y-4 w-full">
              <!-- 标题 -->
              <h3
                class="text-base md:text-xl font-bold group-hover:text-primary transition-colors duration-200 line-clamp-2"
              >
                {{ video.title }}
              </h3>

              <!-- 基本信息 -->
              <div class="grid grid-cols-1 gap-2 md:gap-4 text-xs md:text-sm">
                <div class="space-y-1.5 md:space-y-2">
                  <div class="flex items-center space-x-2">
                    <svg
                      class="w-3.5 h-3.5 md:w-4 md:h-4 text-chart-2 flex-shrink-0"
                      fill="currentColor"
                      viewBox="0 0 20 20"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z"
                        clip-rule="evenodd"
                      ></path>
                    </svg>
                    <span class="text-muted-foreground">作者:</span>
                    <span class="font-medium truncate">{{ video.author?.name }}</span>
                  </div>

                  <div class="flex items-center space-x-2">
                    <svg
                      class="w-3.5 h-3.5 md:w-4 md:h-4 text-chart-5 flex-shrink-0"
                      fill="currentColor"
                      viewBox="0 0 20 20"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z"
                        clip-rule="evenodd"
                      ></path>
                    </svg>
                    <span class="text-muted-foreground">时长:</span>
                    <span class="font-medium">{{ video.duration }} 分钟</span>
                    <span class="text-muted-foreground ml-auto md:ml-4">播放:</span>
                    <span class="font-medium">{{ video.views.toLocaleString() }}</span>
                  </div>

                  <!-- 移动端简化显示背景和世界 -->
                  <div class="flex items-center space-x-2 md:hidden">
                    <svg
                      class="w-3.5 h-3.5 text-chart-3 flex-shrink-0"
                      fill="currentColor"
                      viewBox="0 0 20 20"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zM4.332 8.027a6.012 6.012 0 011.912-2.706C6.512 5.73 6.974 6 7.5 6A1.5 1.5 0 019 7.5V8a2 2 0 004 0 2 2 0 011.523-1.943A5.977 5.977 0 0116 10c0 .34-.028.675-.083 1H15a2 2 0 00-2 2v2.197A5.973 5.973 0 0110 16v-2a2 2 0 00-2-2 2 2 0 01-2-2 2 2 0 00-1.668-1.973z"
                        clip-rule="evenodd"
                      ></path>
                    </svg>
                    <span class="font-medium truncate">
                      {{
                        video.background()?.displayName || '' || video.background()?.name || '未知'
                      }}
                      /
                      {{ video.world()?.displayName || video.world()?.name || '未知' }}
                    </span>
                  </div>
                </div>

                <!-- 桌面端详细信息 -->
                <div class="hidden md:grid md:grid-cols-2 gap-4">
                  <div class="space-y-2">
                    <div class="flex items-center space-x-2">
                      <svg class="w-4 h-4 text-chart-3" fill="currentColor" viewBox="0 0 20 20">
                        <path
                          fill-rule="evenodd"
                          d="M10 18a8 8 0 100-16 8 8 0 000 16zM4.332 8.027a6.012 6.012 0 011.912-2.706C6.512 5.73 6.974 6 7.5 6A1.5 1.5 0 019 7.5V8a2 2 0 004 0 2 2 0 011.523-1.943A5.977 5.977 0 0116 10c0 .34-.028.675-.083 1H15a2 2 0 00-2 2v2.197A5.973 5.973 0 0110 16v-2a2 2 0 00-2-2 2 2 0 01-2-2 2 2 0 00-1.668-1.973z"
                          clip-rule="evenodd"
                        ></path>
                      </svg>
                      <span class="text-muted-foreground">背景:</span>
                      <span class="font-medium">
                        {{ video.background()?.displayName || video.background()?.name || '' }}
                      </span>
                    </div>
                    <div class="flex items-center space-x-2">
                      <svg class="w-4 h-4 text-chart-4" fill="currentColor" viewBox="0 0 20 20">
                        <path
                          fill-rule="evenodd"
                          d="M10 18a8 8 0 100-16 8 8 0 000 16zM4.332 8.027a6.012 6.012 0 011.912-2.706C6.512 5.73 6.974 6 7.5 6A1.5 1.5 0 019 7.5V8a2 2 0 004 0 2 2 0 011.523-1.943A5.977 5.977 0 0116 10c0 .34-.028.675-.083 1H15a2 2 0 00-2 2v2.197A5.973 5.973 0 0110 16v-2a2 2 0 00-2-2 2 2 0 01-2-2 2 2 0 00-1.668-1.973z"
                          clip-rule="evenodd"
                        ></path>
                      </svg>
                      <span class="text-muted-foreground">世界:</span>
                      <span class="font-medium">
                        {{ video.world()?.displayName || video.world()?.name || '' }}
                      </span>
                    </div>
                  </div>

                  <div class="space-y-2">
                    <div class="flex items-center space-x-2">
                      <svg class="w-4 h-4 text-primary" fill="currentColor" viewBox="0 0 20 20">
                        <path d="M10 12a2 2 0 100-4 2 2 0 000 4z"></path>
                        <path
                          fill-rule="evenodd"
                          d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z"
                          clip-rule="evenodd"
                        ></path>
                      </svg>
                      <span class="text-muted-foreground">播放量:</span>
                      <span class="font-medium">{{ video.views.toLocaleString() }}</span>
                    </div>
                  </div>
                </div>
              </div>

              <!-- 风格标签 -->
              <div v-if="video.style.length > 0" class="flex flex-wrap gap-1.5 md:gap-2">
                <span class="text-muted-foreground text-xs md:text-sm">风格:</span>
                <span
                  v-for="(s, i) in video.style()"
                  :key="i"
                  class="inline-flex items-center px-1.5 py-0.5 md:px-2 md:py-1 bg-accent/20 text-accent-foreground text-xs rounded-md font-medium"
                >
                  {{ s.displayName || s.name }}
                </span>
              </div>

              <!-- 状态标识 -->
              <div class="flex flex-wrap items-center gap-3 md:gap-6 text-xs md:text-sm">
                <div class="flex items-center space-x-1.5 md:space-x-2">
                  <div
                    :class="[
                      'w-1.5 h-1.5 md:w-2 md:h-2 rounded-full',
                      video.isOriginal() ? 'bg-primary' : 'bg-muted-foreground',
                    ]"
                  ></div>
                  <span :class="video.isOriginal() ? 'text-primary' : 'text-muted-foreground'">
                    {{ video.isOriginal() ? '原创' : '非原创' }}
                  </span>
                </div>

                <div class="flex items-center space-x-1.5 md:space-x-2">
                  <div
                    :class="[
                      'w-1.5 h-1.5 md:w-2 md:h-2 rounded-full',
                      video.isCompleted() ? 'bg-chart-2' : 'bg-muted-foreground',
                    ]"
                  ></div>
                  <span :class="video.isCompleted() ? 'text-chart-2' : 'text-muted-foreground'">
                    {{ video.isCompleted() ? '已完结' : '连载中' }}
                  </span>
                </div>

                <div class="flex items-center space-x-1.5 md:space-x-2">
                  <div
                    :class="[
                      'w-1.5 h-1.5 md:w-2 md:h-2 rounded-full',
                      video.hasSystem() ? 'bg-chart-3' : 'bg-muted-foreground',
                    ]"
                  ></div>
                  <span :class="video.hasSystem() ? 'text-chart-3' : 'text-muted-foreground'">
                    {{ video.hasSystem() ? '有系统' : '无系统' }}
                  </span>
                </div>
              </div>
            </div>

            <!-- 桌面端操作按钮 -->
            <div class="hidden md:flex flex-shrink-0 items-center space-x-3">
              <button
                @click="approveVideo(video)"
                class="group/btn relative p-3 bg-chart-2/10 hover:bg-chart-2/20 border border-chart-2/30 text-chart-2 rounded-xl transition-all duration-200 hover:shadow-lg hover:-translate-y-1"
                title="通过审核"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M5 13l4 4L19 7"
                  ></path>
                </svg>
                <div
                  class="absolute -bottom-8 left-1/2 transform -translate-x-1/2 px-2 py-1 bg-chart-2 text-white text-xs rounded opacity-0 group-hover/btn:opacity-100 transition-opacity duration-200 pointer-events-none whitespace-nowrap"
                >
                  通过审核
                </div>
              </button>

              <button
                @click="rejectVideo(video)"
                class="group/btn relative p-3 bg-destructive/10 hover:bg-destructive/20 border border-destructive/30 text-destructive rounded-xl transition-all duration-200 hover:shadow-lg hover:-translate-y-1"
                title="驳回"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M6 18L18 6M6 6l12 12"
                  ></path>
                </svg>
                <div
                  class="absolute -bottom-8 left-1/2 transform -translate-x-1/2 px-2 py-1 bg-destructive text-white text-xs rounded opacity-0 group-hover/btn:opacity-100 transition-opacity duration-200 pointer-events-none whitespace-nowrap"
                >
                  驳回内容
                </div>
              </button>

              <button
                @click="goToVideo(video)"
                class="group/btn relative p-3 bg-primary/10 hover:bg-primary/20 border border-primary/30 text-primary rounded-xl transition-all duration-200 hover:shadow-lg hover:-translate-y-1"
                title="查看原视频"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
                  ></path>
                </svg>
                <div
                  class="absolute -bottom-8 left-1/2 transform -translate-x-1/2 px-2 py-1 bg-primary text-white text-xs rounded opacity-0 group-hover/btn:opacity-100 transition-opacity duration-200 pointer-events-none whitespace-nowrap"
                >
                  查看视频
                </div>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { toast } from '@yuelioi/toast'

const route = useRoute()

const videos = ref<Video[]>([])

// 加载未审核视频
const loadUnReviewedVideos = async () => {
  try {
    const res = await videoApi.getUnreviewedVideos()
    videos.value = res.data.videos.map((v: Partial<Video>) => new Video(v))
  } catch (err) {
    handleApiError(err, '加载')
  }
}

// 通过视频
const approveVideo = async (video: Video) => {
  try {
    await videoApi.updateVideoReviewStatus(video.id!, true)
    videos.value = videos.value.filter((v) => v.id !== video.id)
    toast.success(`《${video.title}》 已通过审核`)
  } catch (err) {
    handleApiError(err, '审核')
  }
}

// 驳回视频
const rejectVideo = async (video: Video) => {
  try {
    await videoApi.deleteVideo(video.id!)
    videos.value = videos.value.filter((v) => v.id !== video.id)
    toast.success(`《${video.title}》 已被驳回`)
  } catch (err) {
    handleApiError(err, '驳回')
  }
}

watch(
  () => route.fullPath,
  async () => {
    await loadUnReviewedVideos()
  },
)

// 跳转视频
const goToVideo = (video: Video) => {
  window.open(video.url, '_blank')
}

onMounted(() => {
  loadUnReviewedVideos()
})
</script>

<style scoped>
button:hover {
  transform: translateY(-2px);
}
</style>
