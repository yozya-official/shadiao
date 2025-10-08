<template>
  <!-- 页面标题 -->
  <PageActionHeader title="编辑作品信息" subtitle="修改和完善作品信息">
    <template #icon>
      <svg class="w-12 h-12" fill="white" viewBox="0 0 20 20">
        <path
          d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z"></path>
      </svg>
    </template>
  </PageActionHeader>
  <div class="container px-6 py-8 relative">
    <!-- 主内容区域 -->
    <div class="grid grid-cols-1 xl:grid-cols-4 gap-8">
      <!-- 表单区域 -->
      <div class="xl:col-span-3 space-y-8">
        <!-- 基本信息卡片 -->
        <div
          class="bg-card/90 backdrop-blur-sm border rounded-3xl shadow p-8 hover:shadow-lg transition-all duration-300">
          <div class="flex items-center justify-between mb-6">
            <div class="flex items-center">
              <div
                class="w-4 h-4 bg-gradient-to-br from-primary to-chart-2 rounded-full mr-4"></div>
              <h3 class="font-bold">基本信息</h3>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
            <!-- 作品名称 -->
            <div class="space-y-3">
              <label class="flex items-center space-x-2 font-semibold">
                <svg class="w-5 h-5 text-primary" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M4 3a2 2 0 100 4h12a2 2 0 100-4H4z"></path>
                  <path
                    fill-rule="evenodd"
                    d="M3 8h14v7a2 2 0 01-2 2H5a2 2 0 01-2-2V8zm5 3a1 1 0 011-1h2a1 1 0 110 2H9a1 1 0 01-1-1z"
                    clip-rule="evenodd"></path>
                </svg>
                <span>作品名称</span>
                <span class="text-destructive text-sm">*</span>
              </label>
              <input
                v-model="formData.title"
                type="text"
                placeholder="输入作品名称..."
                class="w-full input input-primary" />
            </div>

            <!-- 作品ID (只读) -->
            <div class="space-y-3">
              <label class="flex items-center space-x-2 font-semibold">
                <svg class="w-5 h-5 text-chart-2" fill="currentColor" viewBox="0 0 20 20">
                  <path
                    fill-rule="evenodd"
                    d="M3 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
                    clip-rule="evenodd"></path>
                </svg>
                <span>作品ID</span>
              </label>
              <input
                :value="formData.aid"
                type="text"
                disabled
                class="w-full input input-primary bg-muted/50 cursor-not-allowed" />
            </div>

            <!-- 链接 -->
            <div class="space-y-3 col-span-full">
              <label class="flex items-center space-x-2 font-semibold">
                <svg
                  class="w-5 h-5 text-chart-3"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path>
                </svg>
                <span>作品链接</span>
                <span class="text-destructive text-sm">*</span>
              </label>
              <input
                v-model="formData.url"
                type="url"
                disabled
                placeholder="输入B站视频链接..."
                class="w-full input input-primary" />
            </div>

            <!-- 封面 -->
            <div class="space-y-3 col-span-full">
              <label class="flex items-center space-x-2 font-semibold">
                <svg
                  class="w-5 h-5 text-chart-3"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                </svg>
                <span>作品封面</span>
                <span class="text-destructive text-sm">*</span>
              </label>
              <input
                v-model="formData.cover"
                type="url"
                placeholder="输入B站视频封面链接..."
                class="w-full input input-primary" />
            </div>

            <div
              v-if="formData.cover"
              class="relative group w-full max-w-md mx-auto rounded-2xl overflow-hidden shadow-lg hover:shadow-2xl transition-all duration-300">
              <img
                :src="formData.cover"
                referrerpolicy="no-referrer"
                alt="封面预览"
                class="w-full h-60 object-cover transform group-hover:scale-105 transition-transform duration-300" />

              <!-- 渐变遮罩 -->
              <div
                class="absolute inset-0 bg-gradient-to-t from-background/70 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>

              <!-- 删除按钮 -->
              <button
                @click.stop="formData.cover = ''"
                class="absolute z-20 top-3 right-3 w-9 h-9 flex items-center justify-center rounded-full bg-destructive/90 text-destructive-foreground opacity-0 group-hover:opacity-100 hover:bg-destructive hover:scale-110 transition-all duration-200 shadow-lg backdrop-blur-sm"
                title="删除封面"
                aria-label="删除封面">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>

            <!-- 描述 -->
            <div class="space-y-3 col-span-full">
              <label class="flex items-center space-x-2 font-semibold">
                <svg
                  class="w-5 h-5 text-chart-4"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z"></path>
                </svg>
                <span>推荐理由</span>
              </label>
              <textarea
                v-model="formData.description"
                placeholder="分享你的推荐理由..."
                rows="4"
                class="w-full textarea rounded-xl input-primary resize-none" />
            </div>

            <!-- 时长和播放量 -->
            <div class="space-y-3">
              <label class="flex items-center space-x-2 font-semibold">
                <svg class="w-5 h-5 text-chart-5" fill="currentColor" viewBox="0 0 20 20">
                  <path
                    fill-rule="evenodd"
                    d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z"
                    clip-rule="evenodd"></path>
                </svg>
                <span>时长</span>
              </label>
              <select
                v-model.number="formData.duration"
                class="w-full select select-primary rounded-xl">
                <option :value="k" v-for="(v, k) in durationOptions" :key="k">{{ v }}</option>
              </select>
            </div>

            <div class="space-y-3">
              <label class="flex items-center space-x-2 font-semibold">
                <svg class="w-5 h-5 text-primary" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M10 12a2 2 0 100-4 2 2 0 000 4z"></path>
                  <path
                    fill-rule="evenodd"
                    d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z"
                    clip-rule="evenodd"></path>
                </svg>
                <span>播放量</span>
              </label>
              <select
                v-model.number="formData.views"
                class="w-full select select-primary rounded-xl">
                <option :value="k" v-for="(v, k) in viewsOptions" :key="k">{{ v }}</option>
              </select>
            </div>
          </div>
        </div>

        <!-- 状态设置卡片 -->
        <div
          class="bg-card/90 backdrop-blur-sm border rounded-3xl shadow p-8 hover:shadow-lg transition-all duration-300">
          <div class="flex items-center mb-6">
            <div class="w-4 h-4 bg-gradient-to-br from-chart-2 to-chart-3 rounded-full mr-4"></div>
            <h3 class="font-bold">状态设置</h3>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
            <!-- 是否原创 -->
            <label
              class="group cursor-pointer bg-muted/20 hover:bg-muted/40 rounded-2xl p-3 border transition-all duration-200 flex items-center justify-between">
              <div class="flex items-center space-x-4">
                <div
                  class="size-6 rounded-xl bg-primary/10 flex items-center justify-center group-hover:bg-primary/20 transition-colors duration-200">
                  <svg class="w-5 h-5 text-primary" fill="currentColor" viewBox="0 0 20 20">
                    <path
                      d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
                  </svg>
                </div>
                <div>
                  <span class="font-semibold">原创作品</span>
                  <p class="text-muted-foreground text-sm">非搬运，原创内容</p>
                </div>
              </div>
              <input
                v-model="formData.isOriginal"
                type="checkbox"
                class="w-5 h-5 text-primary border-2 rounded focus:ring-2 focus:ring-primary/50 transition-all duration-200" />
            </label>

            <!-- 是否完结 -->
            <label
              class="group cursor-pointer bg-muted/20 hover:bg-muted/40 rounded-2xl p-3 border transition-all duration-200 flex items-center justify-between">
              <div class="flex items-center space-x-4">
                <div
                  class="w-10 h-10 rounded-xl bg-chart-2/10 flex items-center justify-center group-hover:bg-chart-2/20 transition-colors duration-200">
                  <svg class="w-5 h-5 text-chart-2" fill="currentColor" viewBox="0 0 20 20">
                    <path
                      fill-rule="evenodd"
                      d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                      clip-rule="evenodd"></path>
                  </svg>
                </div>
                <div>
                  <span class="font-semibold">已完结</span>
                  <p class="text-muted-foreground text-sm">系列动画已完成</p>
                </div>
              </div>
              <input
                v-model="formData.isCompleted"
                type="checkbox"
                class="w-5 h-5 text-chart-2 border-2 rounded focus:ring-2 focus:ring-chart-2/50 transition-all duration-200" />
            </label>
          </div>
        </div>

        <!-- 分类设置卡片 -->
        <div
          class="bg-card/90 backdrop-blur-sm border rounded-3xl shadow p-8 hover:shadow-lg transition-all duration-300">
          <div class="flex items-center mb-6">
            <div class="w-4 h-4 bg-gradient-to-br from-chart-3 to-chart-4 rounded-full mr-4"></div>
            <h3 class="font-bold">分类设置</h3>
          </div>

          <div class="space-y-4">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <!-- 背景设定 -->
              <div class="space-y-4">
                <label class="flex items-center space-x-2 font-semibold">
                  <svg class="size-5 text-chart-3" fill="currentColor" viewBox="0 0 20 20">
                    <path
                      fill-rule="evenodd"
                      d="M10 18a8 8 0 100-16 8 8 0 000 16zM4.332 8.027a6.012 6.012 0 011.912-2.706C6.512 5.73 6.974 6 7.5 6A1.5 1.5 0 019 7.5V8a2 2 0 004 0 2 2 0 011.523-1.943A5.977 5.977 0 0116 10c0 .34-.028.675-.083 1H15a2 2 0 00-2 2v2.197A5.973 5.973 0 0110 16v-2a2 2 0 00-2-2 2 2 0 01-2-2 2 2 0 00-1.668-1.973z"
                      clip-rule="evenodd"></path>
                  </svg>
                  <span>背景设定</span>
                </label>
                <select v-model="formData.background" class="w-full select input input-primary">
                  <option value="">选择背景设定...</option>
                  <option v-for="tag in backgroundOptions" :key="tag.id" :value="tag.name">
                    {{ tag.icon }} {{ tag.displayName || tag.name }}
                  </option>
                </select>
              </div>

              <!-- 世界设定 -->
              <div class="space-y-4">
                <label class="flex items-center space-x-2 font-semibold">
                  <svg class="size-5 text-chart-4" fill="currentColor" viewBox="0 0 20 20">
                    <path
                      fill-rule="evenodd"
                      d="M10 18a8 8 0 100-16 8 8 0 000 16zM4.332 8.027a6.012 6.012 0 011.912-2.706C6.512 5.73 6.974 6 7.5 6A1.5 1.5 0 019 7.5V8a2 2 0 004 0 2 2 0 011.523-1.943A5.977 5.977 0 0116 10c0 .34-.028.675-.083 1H15a2 2 0 00-2 2v2.197A5.973 5.973 0 0110 16v-2a2 2 0 00-2-2 2 2 0 01-2-2 2 2 0 00-1.668-1.973z"
                      clip-rule="evenodd"></path>
                  </svg>
                  <span>世界设定</span>
                </label>
                <select v-model="formData.world" class="w-full select input input-primary">
                  <option value="">选择世界设定...</option>
                  <option v-for="tag in worldOptions" :key="tag.id" :value="tag.name">
                    {{ tag.icon }} {{ tag.displayName || tag.name }}
                  </option>
                </select>
              </div>
            </div>

            <!-- 风格选择 -->
            <label class="flex items-center space-x-2 font-semibold">
              <svg class="size-5 text-chart-5" fill="currentColor" viewBox="0 0 20 20">
                <path
                  fill-rule="evenodd"
                  d="M7 4a3 3 0 016 0v4a3 3 0 11-6 0V4zm4 10.93A7.001 7.001 0 0017 8a1 1 0 10-2 0A5 5 0 015 8a1 1 0 00-2 0 7.001 7.001 0 006 6.93V17H6a1 1 0 100 2h8a1 1 0 100-2h-3v-2.07z"
                  clip-rule="evenodd"></path>
              </svg>
              <span>动画风格（可多选）</span>
            </label>

            <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-3">
              <label
                v-for="tag in styleOptions"
                :key="tag.id"
                class="group cursor-pointer bg-muted/20 hover:bg-muted/40 rounded-xl p-2 border transition-all duration-200 flex items-center justify-between"
                :class="formData.style.includes(tag.name) ? 'border-primary bg-primary/10' : ''">
                <div class="flex items-center space-x-3">
                  <span class="">{{ tag.icon }}</span>
                  <span class="font-medium">{{ tag.displayName || tag.name }}</span>
                </div>
                <input
                  v-model="formData.style"
                  type="checkbox"
                  :value="tag.name"
                  class="w-4 h-4 text-primary border-2 rounded focus:ring-2 focus:ring-primary/50 transition-all duration-200" />
              </label>
            </div>

            <!-- 系统外挂 -->
            <label
              class="group cursor-pointer bg-muted/20 hover:bg-muted/40 rounded-2xl p-3 border transition-all duration-200 flex items-center justify-between"
              :class="formData.hasSystem ? 'border-primary bg-primary/10' : ''">
              <div class="flex items-center space-x-4">
                <div
                  class="w-10 h-10 rounded-xl bg-chart-3/10 flex items-center justify-center group-hover:bg-chart-3/20 transition-colors duration-200">
                  <svg class="w-5 h-5 text-chart-3" fill="currentColor" viewBox="0 0 20 20">
                    <path
                      fill-rule="evenodd"
                      d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z"
                      clip-rule="evenodd"></path>
                  </svg>
                </div>
                <div>
                  <span class="font-semibold">系统外挂</span>
                  <p class="text-muted-foreground text-sm">主角拥有系统、金手指等外挂能力</p>
                </div>
              </div>
              <input
                v-model="formData.hasSystem"
                type="checkbox"
                class="w-5 h-5 text-chart-3 border-2 rounded focus:ring-2 focus:ring-chart-3/50 transition-all duration-200" />
            </label>
          </div>
        </div>
      </div>

      <!-- 操作区域 -->
      <div class="xl:col-span-1">
        <div class="sticky top-8 space-y-6">
          <!-- 操作按钮 -->
          <div class="bg-card/90 backdrop-blur-sm border rounded-2xl shadow p-3">
            <h3 class="font-bold mb-6 flex items-center">
              <svg class="w-5 h-5 text-primary mr-2" fill="currentColor" viewBox="0 0 20 20">
                <path
                  fill-rule="evenodd"
                  d="M3 4a1 1 0 011-1h4a1 1 0 010 2H6.414l2.293 2.293a1 1 0 11-1.414 1.414L5 6.414V8a1 1 0 01-2 0V4zm9 1a1 1 0 010-2h4a1 1 0 011 1v4a1 1 0 01-2 0V6.414l-2.293 2.293a1 1 0 11-1.414-1.414L13.586 5H12zm-9 7a1 1 0 012 0v1.586l2.293-2.293a1 1 0 111.414 1.414L6.414 15H8a1 1 0 010 2H4a1 1 0 01-1-1v-4zm13-1a1 1 0 011 1v4a1 1 0 01-1 1h-4a1 1 0 010-2h1.586l-2.293-2.293a1 1 0 111.414-1.414L15 13.586V12a1 1 0 011-1z"
                  clip-rule="evenodd"></path>
              </svg>
              操作
            </h3>
            <div class="space-y-4">
              <button
                @click="handleUpdate"
                :disabled="!formData.title || !formData.url"
                class="w-full btn btn-lg btn-primary border text-primary-foreground shadow hover:shadow-lg disabled:cursor-not-allowed transition-all duration-200 font-semibold flex items-center justify-center space-x-2">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M5 13l4 4L19 7"></path>
                </svg>
                <span>保存修改</span>
              </button>
              <button
                @click="handleCancel"
                class="w-full btn btn-lg border bg-muted/40 hover:bg-muted/60 text-muted-foreground transition-all duration-200 font-medium flex items-center justify-center space-x-2">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M6 18L18 6M6 6l12 12"></path>
                </svg>
                <span>取消编辑</span>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { PageActionHeader } from '@yuelioi/ui'

import { toast } from '@yuelioi/toast'

const videoStore = useVideoStore()
const { backgroundOptions, styleOptions, worldOptions } = storeToRefs(videoStore)
const { loadTags } = videoStore

import { durationOptions, viewsOptions } from '@/stores/options'

const route = useRoute()
const router = useRouter()

const formData = reactive<VideoData>({
  title: '',
  aid: 0,
  url: '',
  cover: '',
  duration: 0,
  description: '',
  views: 0,
  isOriginal: false,
  isCompleted: false,
  world: '',
  background: '',
  style: [],
  hasSystem: false,
  ctime: '',
})

// 加载视频数据
onMounted(async () => {
  const videoId = route.params.id as string
  if (!videoId) {
    toast.error('视频ID不存在')
    router.push('/')
    return
  }

  try {
    const response = await videoApi.getVideoById(videoId)
    const v = new Video(response.data.video)
    Object.assign(formData, response.data.video)

    formData.isCompleted = v.isCompleted()
    formData.isOriginal = v.isOriginal()

    formData.background = v.background()?.name
    formData.world = v.world()?.name

    formData.style = v.style()?.map((s) => s.name)
    formData.hasSystem = v.hasSystem()
  } catch (error: unknown) {
    handleApiError(error, '加载')
    // router.push('/')
  }
})

// 更新视频
const handleUpdate = async () => {
  if (!formData.title || !formData.url) {
    toast.error('请填写作品名称和链接')
    return
  }

  const videoId = route.params.id as string

  try {
    await videoApi.updateVideo(videoId, formData)
    toast.success('视频更新成功！')
    router.push('/')
  } catch (error: unknown) {
    handleApiError(error, '更新')
  }
}

// 取消编辑
const handleCancel = () => {
  router.back()
}

onMounted(async () => {
  await loadTags()
})
</script>
